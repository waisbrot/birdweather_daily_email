package birdweather

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Khan/genqlient/graphql"
	"github.com/waisbrot/birdweather_daily_email/metrics"
	"github.com/waisbrot/birdweather_daily_email/structs"
)

func BirdsForStation(stationid string) (string, []structs.BirdCount) {
	client := graphql.NewClient("https://app.birdweather.com/graphql", http.DefaultClient)
	counts, err := dailyCounts(context.Background(), client, stationid)
	if err != nil {
		panic(err)
	}

	var result = []structs.BirdCount{}
	for _, count := range counts.Station.TopSpecies {
		var bc structs.BirdCount
		bc.Name = count.Species.CommonName
		bc.SciName = count.Species.ScientificName
		bc.ImageURL = count.Species.ImageUrl
		bc.ImageCredit = count.Species.ImageCredit
		bc.Count = count.Breakdown.AlmostCertain

		result = append(result, bc)
	}
	metrics.RecordFetch(counts.Station.Name, len(result))
	return counts.Station.Name, result
}

func RecordCountsForStationPastHours(stationId string, hours int) {
	client := graphql.NewClient("https://app.birdweather.com/graphql", http.DefaultClient)
	duration := InputDuration{
		Count: hours,
		Unit:  "hour",
	}
	counts, err := hourlyCounts(context.Background(), client, stationId, duration)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Got %d species for station %s\n", len(counts.Station.TopSpecies), stationId)
	metrics.RecordSpecies(counts.Station.Name, len(counts.Station.TopSpecies))
	for _, count := range counts.Station.TopSpecies {
		metrics.RecordBird(counts.Station.Name, count.Species.CommonName, count.Breakdown.AlmostCertain)
		fmt.Printf("Recorded %d of %s\n", count.Breakdown.AlmostCertain, count.Species.CommonName)
	}
}
