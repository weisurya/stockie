package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Short: "stock-data-crawler",
		Long:  "stock-data-crawler is Go-based system for people who wants to retrieve data from IDX and pre-process it into clean data",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
