package main

import (
	"database/sql"
	"os"
	"os/signal"
	"syscall"
	"time"

	. "github.com/rumyantseva/velobike-statistics/models"

	_ "github.com/lib/pq"
	"github.com/rumyantseva/go-velobike/v4/velobike"
	"github.com/sirupsen/logrus"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
)

// Count station statistics, store it in the database
func main() {
	log := logrus.New()
	log.SetOutput(os.Stdout)
	log.Info("Starting the app...")

	// Create connection
	conn, err := sql.Open(
		"postgres",
		"postgres://db-user:db-password@velostat-db:5432/velostat?sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Ping the DB...")
	for i := 0; i < 10; i++ {
		err = conn.Ping()
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatal("Can't ping DB: %v", err)
	}

	// Create a DB instance
	db := reform.NewDB(
		conn,
		postgresql.Dialect,
		reform.NewPrintfLogger(log.Printf),
	)

	log.Info("Query the DB...")
	for i := 0; i < 10; i++ {
		_, err = db.Query("SELECT 1")
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatal("Can't connect to the DB: %v", err)
	}

	// Get the list of parkings
	client := velobike.NewClient()

	c := time.Tick(5 * time.Minute)
	ret := time.Tick(1 * time.Hour)
	start := make(chan struct{}, 1)
	stop := make(chan struct{}, 1)

	log.Info("Starting the ticker...")
	start <- struct{}{}

	go func() {
		for {
			select {
			case <-c:
			case <-start:
				log.Info("Save statistics...")
				err := saveStats(client, db)
				if err != nil {
					log.Errorf("Got an error: %v", err)
					stop <- struct{}{}
				}

			case <-ret:
				stop <- struct{}{}
			}
		}
	}()

	// Shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case x := <-interrupt:
		log.Infof("Received a signal: %s, the app will be stopped.", x.String())

	case <-stop:
		log.Info("The app is ready to be stopped")
	}
}

func saveStats(client *velobike.Client, db *reform.DB) error {
	parkings, _, err := client.Parkings.List()
	if err != nil {
		return err
	}

	for _, item := range parkings.Items {
		station := &Station{
			Station:     *item.ID,
			Address:     *item.Address,
			Lon:         *item.Position.Lon,
			Lat:         *item.Position.Lat,
			TotalPlaces: int32(*item.TotalPlaces),
			FreePlaces:  int32(*item.FreePlaces),
			IsLocked:    *item.IsLocked,
		}

		return db.Save(station)
	}

	return nil
}
