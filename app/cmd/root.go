package cmd

import (
	"context"
	"weisurya/stockie"
	"weisurya/stockie/app/cmd/helper"
	"weisurya/stockie/internals/crawler"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Short: "stock-data-crawler",
		Long:  "stock-data-crawler is Go-based system for people who wants to retrieve data from IDX and pre-process it into clean data",
	}
)

var (
	crawlerService stockie.ICrawlerService
)

func Execute() {
	cobra.OnInitialize(initServices)

	helper.Log = helper.NewLog()

	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}

func initServices() {
	ctx := context.Background()

	crawlerService = crawler.New(ctx)
}
