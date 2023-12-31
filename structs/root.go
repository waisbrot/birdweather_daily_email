package structs

import "time"

type BirdCount struct {
	Name        string
	SciName     string
	ImageURL    string
	ImageCredit string
	Count       int
}

type EmailTemplate struct {
	Day      time.Weekday
	Stations []StationTemplate
}

type StationTemplate struct {
	Name   string
	Id     int
	Counts []BirdCount
}
