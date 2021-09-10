package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	crawlFinancialStatementCmd = &cobra.Command{
		Use:   "crawl_statement",
		Short: "Crawl IDX financial statement",
		Long:  "Start crawling financial statements. Currently only support IDX. This system refers to the list of tickers in ./data/tickers/idx.csv",
		Run:   startCrawlingFinancialStatement,
	}
)

var (
	startYear int
	endYear   int
	quarter   int
)

func init() {
	rootCmd.AddCommand(crawlFinancialStatementCmd)

	crawlFinancialStatementCmd.Flags().IntVarP(
		&startYear,
		"startyear",
		"s",
		2016,
		"[Required] Start year",
	)
	crawlFinancialStatementCmd.MarkFlagRequired("startyear")

	crawlFinancialStatementCmd.Flags().IntVarP(
		&endYear,
		"endyear",
		"e",
		2021,
		"[Required] End year",
	)
	crawlFinancialStatementCmd.MarkFlagRequired("endyear")

	crawlFinancialStatementCmd.Flags().IntVarP(
		&quarter,
		"endquarter",
		"q",
		2,
		"[Required] End quarter",
	)
	crawlFinancialStatementCmd.MarkFlagRequired("endquarter")
}

func startCrawlingFinancialStatement(cmd *cobra.Command, args []string) {
	if err := crawlerService.StartCrawlFinancialStatement(
		startYear, endYear, quarter,
	); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	logrus.Infof("Done pulling IDX financial statements from %v to %v-%v", startYear, endYear, quarter)
}
