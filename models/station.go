package models

//go:generate reform

type (
	//reform:station_statistics
	Station struct {
		ID      int32  `reform:"stat_id,pk"`
		Station string `reform:"station"`
		Address string `reform:"address"`
		Lon         float64 `reform:"lon"`
		Lat         float64 `reform:"lat"`
		TotalPlaces int32   `reform:"total_places"`
		FreePlaces  int32   `reform:"free_places"`
		IsLocked    bool    `reform:"is_locked"`
	}
)
