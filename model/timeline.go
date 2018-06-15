package model

type TimelineHistory struct {
	Locations []*TimelineLocation `json:"locations"`
}

func (history TimelineHistory) Size() int {
	return len(history.Locations)
}

type TimelineLocation struct {
	TimestampMs int64 `json:"timestampMs,string"`
	LatitudeE7  int64 `json:"latitudeE7"`
	LongitudeE7 int64 `json:"longitudeE7"`
}
