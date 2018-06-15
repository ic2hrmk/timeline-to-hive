package repository

import (
	"playground/model"

	"encoding/csv"
	"os"
	"fmt"
)

func SaveHiveOutput(filePath string, history *model.HiveHistory) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range history.Locations {
		//	0. Latitude
		//	1. Longitude
		//	2. Timestamp
		//	3. Username

		row := make([]string, 4)
		row[0] = fmt.Sprintf("%.7f", record.Latitude)
		row[1] = fmt.Sprintf("%.7f", record.Longitude)
		row[2] = fmt.Sprintf("%d", record.CreateTimestamp)
		row[3] = "Roman Kaporin"

		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}
