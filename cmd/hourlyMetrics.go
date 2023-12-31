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
			birdweather.RecordCountsForStationPastMinutes(fmt.Sprint(stationId), 30)
		}
	},
}

func init() {
	rootCmd.AddCommand(hourlyMetricsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hourlyMetricsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hourlyMetricsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
