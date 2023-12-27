/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"github.com/waisbrot/birdweather_daily_email/cmd"
	"github.com/waisbrot/birdweather_daily_email/metrics"
)

func main() {
	configFile := os.Getenv("BIRDWEATHER_CONFIG_FILE")
	if configFile == "" {
		configFile = "/etc/birdweather/config.yaml"
	}
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading config file `%s`: %v\n", configFile, err)
		os.Exit(1)
	}
	metrics.Init()
	metrics.RecordInvoked()
	cmd.Execute()
	metrics.Finish()
}
