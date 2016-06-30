package models

// generated with gopkg.in/reform.v1

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type stationTable struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *stationTable) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("station_statistics").
func (v *stationTable) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *stationTable) Columns() []string {
	return []string{"stat_id", "station", "address", "lon", "lat", "total_places", "free_places", "is_locked"}
}

// NewStruct makes a new struct for that view or table.
func (v *stationTable) NewStruct() reform.Struct {
	return new(Station)
}

// NewRecord makes a new record for that table.
func (v *stationTable) NewRecord() reform.Record {
	return new(Station)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *stationTable) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// StationTable represents station_statistics view or table in SQL database.
var StationTable = &stationTable{
	s: parse.StructInfo{Type: "Station", SQLSchema: "", SQLName: "station_statistics", Fields: []parse.FieldInfo{{Name: "ID", Type: "int32", Column: "stat_id"}, {Name: "Station", Type: "string", Column: "station"}, {Name: "Address", Type: "string", Column: "address"}, {Name: "Lon", Type: "float64", Column: "lon"}, {Name: "Lat", Type: "float64", Column: "lat"}, {Name: "TotalPlaces", Type: "int32", Column: "total_places"}, {Name: "FreePlaces", Type: "int32", Column: "free_places"}, {Name: "IsLocked", Type: "bool", Column: "is_locked"}}, PKFieldIndex: 0},
	z: new(Station).Values(),
}

// String returns a string representation of this struct or record.
func (s Station) String() string {
	res := make([]string, 8)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "Station: " + reform.Inspect(s.Station, true)
	res[2] = "Address: " + reform.Inspect(s.Address, true)
	res[3] = "Lon: " + reform.Inspect(s.Lon, true)
	res[4] = "Lat: " + reform.Inspect(s.Lat, true)
	res[5] = "TotalPlaces: " + reform.Inspect(s.TotalPlaces, true)
	res[6] = "FreePlaces: " + reform.Inspect(s.FreePlaces, true)
	res[7] = "IsLocked: " + reform.Inspect(s.IsLocked, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Station) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.Station,
		s.Address,
		s.Lon,
		s.Lat,
		s.TotalPlaces,
		s.FreePlaces,
		s.IsLocked,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Station) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.Station,
		&s.Address,
		&s.Lon,
		&s.Lat,
		&s.TotalPlaces,
		&s.FreePlaces,
		&s.IsLocked,
	}
}

// View returns View object for that struct.
func (s *Station) View() reform.View {
	return StationTable
}

// Table returns Table object for that record.
func (s *Station) Table() reform.Table {
	return StationTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Station) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Station) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Station) HasPK() bool {
	return s.ID != StationTable.z[StationTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Station) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = StationTable
	_ reform.Struct = new(Station)
	_ reform.Table  = StationTable
	_ reform.Record = new(Station)
	_ fmt.Stringer  = new(Station)
)

func init() {
	parse.AssertUpToDate(&StationTable.s, new(Station))
}
