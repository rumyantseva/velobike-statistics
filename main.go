package main

import (
	"database/sql"
	"github.com/Sirupsen/logrus"
	_ "github.com/lib/pq"
	"github.com/rumyantseva/go-velobike/velobike"
	. "github.com/rumyantseva/velobike-statistics/models"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
)

// Count station statistics, store it in the database
func main() {
	log := logrus.New()

	conn, err := sql.Open(
		"postgres",
		"postgres://postgres:postgres@localhost:5432/velobike?sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}

	DB := reform.NewDB(
		conn,
		postgresql.Dialect,
		reform.NewPrintfLogger(log.Printf),
	)

	client := velobike.NewClient(nil)

	parkings, _, err := client.Parkings.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range parkings.Items {
		station := &Station{
			Station:     *item.Id,
			Address:     *item.Address,
			Lon:         *item.Position.Lon,
			Lat:         *item.Position.Lat,
			TotalPlaces: int32(*item.TotalPlaces),
			FreePlaces:  int32(*item.FreePlaces),
			IsLocked:    *item.IsLocked,
		}
		DB.Save(station)
		if err := DB.Save(station); err != nil {
			log.Fatal(err)
		}
	}
}
