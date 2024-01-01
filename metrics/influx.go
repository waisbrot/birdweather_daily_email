package metrics

import (
	"context"
	"fmt"
	"time"

	influxdb "github.com/influxdata/influxdb-client-go/v2"
	influxapi "github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/spf13/viper"
)

var influxClient influxdb.Client
var writeAPI influxapi.WriteAPIBlocking
var invokeTime time.Time

func initInflux(args []string) {
	var command string
	if len(args) > 0 {
		command = args[0]
	} else {
		command = ""
	}
	options := &influxdb.Options{}
	options.SetApplicationName("birdweather")
	options.AddDefaultTag("command", command)
	influxClient = influxdb.NewClientWithOptions(
		viper.GetString("influx.url"),
		viper.GetString("influx.token"),
		options)
	writeAPI = influxClient.WriteAPIBlocking(viper.GetString("influx.org"), viper.GetString("influx.bucket"))

	produceMetrics = true
}

func finishInflux() {
	duration := time.Since(invokeTime)
	p := influxdb.NewPointWithMeasurement("execution").
		AddField("latency", duration.Milliseconds()).
		SetTime(time.Now())
	writePoint(p)
	influxClient.Close()
}

func writePoint(p *write.Point) {
	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wrote data-point %#v\n", p)
}

func recordInfluxFetch(station string, speciesCount int) {
	p := influxdb.NewPointWithMeasurement("fetch").
		AddTag("station", station).
		AddField("species", speciesCount).
		SetTime(time.Now())
	writePoint(p)
}

func recordInfluxInvoked() {
	invokeTime = time.Now()
}

func recordInfluxEmail(recipientCount int, bodyLength int) {
	p := influxdb.NewPointWithMeasurement("email").
		AddField("recipients", recipientCount).
		AddField("body_length", bodyLength).
		SetTime(time.Now())
	writePoint(p)
}

func recordInfluxBird(stationName string, birdName string, count int) {
	p := influxdb.NewPointWithMeasurement("bird").
		AddTag("station", stationName).
		AddTag("name", birdName).
		AddField("count", count).
		SetTime(time.Now())
	writePoint(p)
}
