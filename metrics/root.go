package metrics

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var produceMetrics bool = false

func Init(args []string) {
	if viper.IsSet("influx.url") {
		initInflux(args)
	} else {
		fmt.Fprintf(os.Stderr, "No influx.url defined. Skipping metrics production.")
	}
}

func Finish() {
	if produceMetrics {
		finishInflux()
	}
}

func RecordFetch(station string, speciesCount int) {
	if produceMetrics {
		recordInfluxFetch(station, speciesCount)
	}
}

func RecordInvoked() {
	if produceMetrics {
		recordInfluxInvoked()
	}
}

func RecordEmail(recipientCount int, bodyLength int) {
	if produceMetrics {
		recordInfluxEmail(recipientCount, bodyLength)
	}
}

func RecordBird(stationName string, birdName string, count int) {
	if produceMetrics {
		recordInfluxBird(stationName, birdName, count)
	}
}
