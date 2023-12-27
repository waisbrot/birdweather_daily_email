/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/waisbrot/birdweather_daily_email/birdweather"
	"github.com/waisbrot/birdweather_daily_email/email"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		stations := []email.StationTemplate{}
		stationIds := viper.GetIntSlice("stations")
		for _, stationId := range stationIds {
			stationName, counts := birdweather.BirdsForStation(fmt.Sprint(stationId))
			templ := email.StationTemplate{}
			templ.Name = stationName
			templ.Id = stationId
			templ.Counts = counts
			stations = append(stations, templ)
		}
		emails := viper.GetStringSlice("email.recipients")
		if len(emails) == 0 {
			fmt.Printf("No emails listed so no emails sent!\n")
			return
		}
		yesterday := time.Now()
		yesterday = yesterday.Add(time.Hour * -24)
		templ := email.EmailTemplate{
			Day:      yesterday.Weekday(),
			Stations: stations,
		}
		emailBody := email.RenderTemplate(templ)
		subject := fmt.Sprintf("Yesterday's Birds (%s)", templ.Day.String())
		email.SendMail(emails, subject, emailBody)
		fmt.Printf("Sent %d emails\n", len(emails))
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)
}
