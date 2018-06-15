package main

import (
	"log"
	"flag"
	"playground/repository"
	"playground/converter"
)

var (
	timelineFilepath string
	hiveFilepath     string
)

func init() {
	flag.StringVar(&timelineFilepath, "timeline", "Location History.json", "")
	flag.StringVar(&hiveFilepath, "hive", "location.csv", "Output CSV for Hive")
	flag.Parse()
}

func main() {
	log.Println("Converting Google Timeline to Apache Hive data file")

	//
	// Load history
	//
	log.Print("loading history from file...")

	timeline, err := repository.LoadTimeline(timelineFilepath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("done, entries:", timeline.Size())

	//
	// Convert history
	//
	log.Print("converting data...")

	hiveHistory := converter.TimelineToHive(timeline)

	log.Println("done!")

	//
	// Save output
	//
	log.Println("saving data...")

	if err := repository.SaveHiveOutput(hiveFilepath, hiveHistory); err != nil {
		log.Fatal(err)
	}

	log.Println("done")
}
