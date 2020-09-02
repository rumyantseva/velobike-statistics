package main

import (
	"database/sql"
	"log"
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

	// Create connection
	conn, err := sql.Open(
		"postgres",
		"postgres://db-user:db-password@velostat-db:5432/velostat?sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}

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
	stop := make(chan struct{}, 1)
	go func() {
		for {
			select {
			case <-c:
				log.Info("Save statistics...")
				saveStats(client, db)

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

	case <-ret:
		log.Info("The app is ready to be stopped")
	}
}

func saveStats(client *velobike.Client, db *reform.DB) {
	parkings, _, err := client.Parkings.List()
	if err != nil {
		log.Fatal(err)
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
		db.Save(station)
		if err := db.Save(station); err != nil {
			log.Fatal(err)
		}
	}
}
