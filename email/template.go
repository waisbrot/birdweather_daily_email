package email

import (
	"bytes"
	"html/template"
	"io"
	"path"

	"github.com/spf13/viper"
	"github.com/waisbrot/birdweather_daily_email/structs"
)

func readTemplate() *template.Template {
	templateFile := viper.GetString("email.template")
	base := path.Base(templateFile)
	template, err := template.New(base).ParseFiles(templateFile)
	if err != nil {
		panic(err)
	}
	return template
}

func RenderTemplate(variables structs.EmailTemplate) io.Reader {
	template := readTemplate()
	buffer := new(bytes.Buffer)
	err := template.Execute(buffer, variables)
	if err != nil {
		panic(err)
	}
	return buffer
}
