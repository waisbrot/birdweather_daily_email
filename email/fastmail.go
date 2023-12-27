package email

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/smtp"
	"net/textproto"

	"github.com/jordan-wright/email"
	"github.com/spf13/viper"
	"github.com/waisbrot/birdweather_daily_email/metrics"
)

func SendMail(recipients []string, subject string, body io.Reader) {
	html, err := io.ReadAll(body)
	if err != nil {
		panic(err)
	}
	mail := &email.Email{
		To:      recipients,
		From:    viper.GetString("email.sender"),
		Subject: subject,
		Text:    []byte("Only HTML provided"),
		HTML:    html,
		Headers: textproto.MIMEHeader{},
	}
	auth := smtp.PlainAuth("", viper.GetString("email.smtp.user"), viper.GetString("email.smtp.pass"), viper.GetString("email.smtp.host"))
	tlsConfig := new(tls.Config)
	tlsConfig.ServerName = viper.GetString("email.smtp.host")
	err = mail.SendWithTLS(fmt.Sprintf("%s:%d", viper.GetString("email.smtp.host"), viper.GetInt("email.smtp.port")), auth, tlsConfig)
	if err != nil {
		panic(err)
	}
	metrics.RecordEmail(len(recipients), len(html))
}
