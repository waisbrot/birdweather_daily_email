package email

import (
	"bytes"
	"html/template"
	"io"
	"os"
	"path"
	"time"

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
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	execDir := path.Dir(exePath)
	templateFile := path.Join(execDir, "countEmail.tmpl")
	template, err := template.New("countEmail.tmpl").ParseFiles(templateFile)
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
