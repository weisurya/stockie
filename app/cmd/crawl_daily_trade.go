package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	crawlDailyTradeCmd = &cobra.Command{
		Use:   "crawl_daily_trade",
		Short: "Crawl IDX daily trade",
		Long:  "Start crawling latest daily trade. Currently only support IDX. This system will update information in ./data/daily_trade folder. If startdate and/or enddate not defined, it will be filled to today's date by default.",
		Run:   startCrawlingDailyTrade,
	}
)

var (
	startDate string
	endDate   string
)

func init() {
	rootCmd.AddCommand(crawlDailyTradeCmd)

	crawlDailyTradeCmd.Flags().StringVarP(
		&startDate,
		"startdate",
		"s",
		"",
		"[Optional] Start date. Format: YYYY-MM-DD",
	)

	crawlDailyTradeCmd.Flags().StringVarP(
		&endDate,
		"enddate",
		"e",
		"",
		"[Optional] End date. Format: YYYY-MM-DD",
	)
}

func startCrawlingDailyTrade(cmd *cobra.Command, args []string) {
	if err := crawlerService.StartCrawlDailyTrade(
		startDate,
		endDate,
	); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	logrus.Infof("Done pulling IDX daily trade from %v to %v", startDate, endDate)
}
