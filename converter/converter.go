package converter

import "playground/model"

func TimelineToHive(timeline *model.TimelineHistory) *model.HiveHistory {
	hiveHistory := &model.HiveHistory{
		Locations: make([]*model.HiveLocation, len(timeline.Locations)),
	}

	for i, record := range timeline.Locations {
		hiveHistory.Locations[i] = &model.HiveLocation{
			Latitude:        float64(record.LatitudeE7) / 1e7,
			Longitude:       float64(record.LongitudeE7) / 1e7,
			CreateTimestamp: int64(record.TimestampMs),
		}
	}

	return hiveHistory
}
