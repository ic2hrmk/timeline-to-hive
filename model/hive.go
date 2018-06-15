package model

type HiveHistory struct {
	Locations []*HiveLocation
}

type HiveLocation struct {
	Latitude        float64
	Longitude       float64
	CreateTimestamp int64
}
