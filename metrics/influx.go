package metrics

import (
	"context"
	"crypto/tls"
	"time"

	influxdb "github.com/influxdata/influxdb-client-go/v2"
	influxapi "github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/spf13/viper"
)

var influxClient influxdb.Client
var writeAPI influxapi.WriteAPIBlocking
var invokeTime time.Time

func initInflux() {
	influxClient = influxdb.NewClient(viper.GetString("influx.url"), viper.GetString("influx.token"))
	influxdb.DefaultOptions().AddDefaultTag("application", "birdweather_digest").SetTLSConfig(&tls.Config{InsecureSkipVerify: true})
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
