package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	crawlCorporateActionCalendar = &cobra.Command{
		Use:   "crawl_corporate_action",
		Short: "Crawl IDX Corporate Action",
		Long:  "Start crawling latest corporate action. Currently only support IDX. This system will update information in ./data/daily_trade folder. If startmonth and/or endmonth not defined, it will be filled to today's date by default.",
		Run:   startCrawlingCorporateActionCalendar,
	}
)

var (
	isDownloadAll bool
)

func init() {
	rootCmd.AddCommand(crawlCorporateActionCalendar)

	crawlCorporateActionCalendar.Flags().BoolVarP(
		&isDownloadAll,
		"is_download_all",
		"a",
		false,
		"[Optional] Download all corporate action. By default false. If it is false, the system will only retrieve information based on this month's date",
	)

	crawlCorporateActionCalendar.Flags().StringVarP(
		&startDate,
		"startdate",
		"s",
		"",
		"[Optional] Start date. Format: YYYY-MM-DD",
	)

	crawlCorporateActionCalendar.Flags().StringVarP(
		&endDate,
		"enddate",
		"e",
		"",
		"[Optional] End date. Format: YYYY-MM-DD",
	)
}

func startCrawlingCorporateActionCalendar(cmd *cobra.Command, args []string) {
	if err := crawlerService.StartCrawlCorporateActions(
		isDownloadAll,
		startDate,
		endDate,
	); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	logrus.Infof("Done pulling corporate action data | Download all: %v | from %v to %v", isDownloadAll, startDate, endDate)
}
