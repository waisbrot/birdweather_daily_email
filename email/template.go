package email

import (
	"bytes"
	"html/template"
	"io"
	"path"
	"time"

	"github.com/spf13/viper"
	"github.com/waisbrot/birdweather_daily_email/birdweather"
)

type EmailTemplate struct {
	Day      time.Weekday
	Stations []StationTemplate
}

type StationTemplate struct {
	Name   string
	Id     int
	Counts []birdweather.BirdCount
}

func readTemplate() *template.Template {
	templateFile := viper.GetString("email.template")
	base := path.Base(templateFile)
	template, err := template.New(base).ParseFiles(templateFile)
	if err != nil {
		panic(err)
	}
	return template
}

func RenderTemplate(variables EmailTemplate) io.Reader {
	template := readTemplate()
	buffer := new(bytes.Buffer)
	err := template.Execute(buffer, variables)
	if err != nil {
		panic(err)
	}
	return buffer
}
