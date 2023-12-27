package birdweather

import (
	"context"
	"net/http"

	"github.com/Khan/genqlient/graphql"
	"github.com/waisbrot/birdweather_daily_email/metrics"
)

type BirdCount struct {
	Name        string
	SciName     string
	ImageURL    string
	ImageCredit string
	Count       int
}

func BirdsForStation(stationid string) (string, []BirdCount) {
	client := graphql.NewClient("https://app.birdweather.com/graphql", http.DefaultClient)
	counts, err := dailyCounts(context.Background(), client, stationid)
	if err != nil {
		panic(err)
	}

	var result = []BirdCount{}
	for _, count := range counts.Station.TopSpecies {
		var bc BirdCount
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
