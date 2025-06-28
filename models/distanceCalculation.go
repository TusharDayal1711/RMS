package models

type CoordinatesReq struct {
	Longitude float64 `db:"longitude"`
	Latitude  float64 `db:"latitude"`
}
