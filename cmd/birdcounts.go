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

// birdcountsCmd represents the birdcounts command
var birdcountsCmd = &cobra.Command{
	Use:   "birdcounts",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		stationIds := viper.GetIntSlice("stations")
		fmt.Printf("stationIds=%#v\nconfig=%#v\nfile=%#v\n", stationIds, viper.AllSettings(), viper.ConfigFileUsed())
		_, counts := birdweather.BirdsForStation(fmt.Sprint(stationIds[0]))
		for i, count := range counts {
			fmt.Printf("%d: %#v\n", i, count)
		}
	},
}

func init() {
	rootCmd.AddCommand(birdcountsCmd)
}
