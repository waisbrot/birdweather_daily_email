/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"time"

	"github.com/spf13/cobra"
	"github.com/waisbrot/birdweather_daily_email/email"
	"github.com/waisbrot/birdweather_daily_email/structs"
)

// templatetestCmd represents the templatetest command
var templatetestCmd = &cobra.Command{
	Use:   "templatetest",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("templatetest called")
		templ := structs.EmailTemplate{
			Day: time.Sunday,
			Stations: []structs.StationTemplate{
				{
					Name: "Example station",
					Id:   42,
					Counts: []structs.BirdCount{
						{
							Name:        "Chickadee",
							SciName:     "Dee-dee-dee",
							ImageURL:    "https://media.birdweather.com/species/120/Black-cappedChickadee-standard-8ff37bd75ba774bdb7ae61cc453602dd.jpg",
							ImageCredit: "Testing",
							Count:       200,
						},
						{
							Name:        "Golden-crowned Kinglet",
							SciName:     "Regulus satrapa",
							ImageURL:    "https://media.birdweather.com/species/41/Golden-crownedKinglet-standard-5a054aa6979fac7d1b338d4e7cb77887.jpg",
							ImageCredit: "Testing",
							Count:       24,
						},
					},
				},
			},
		}
		templReader := email.RenderTemplate(templ)
		bytes, err := io.ReadAll(templReader)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(bytes))
	},
}

func init() {
	rootCmd.AddCommand(templatetestCmd)
}
