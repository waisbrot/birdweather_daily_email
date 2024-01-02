/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/waisbrot/birdweather_daily_email/birdweather"
)

// hourlyMetricsCmd represents the hourlyMetrics command
var hourlyMetricsCmd = &cobra.Command{
	Use:   "hourlyMetrics",
	Short: "Fetch and push bird metrics",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		stationIds := viper.GetIntSlice("stations")
		for _, stationId := range stationIds {
			fmt.Printf("Recording counts for station %d\n", stationId)
			birdweather.RecordCountsForStationPastHours(fmt.Sprint(stationId), 1)
		}
	},
}

func init() {
	rootCmd.AddCommand(hourlyMetricsCmd)
}
