package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	crawlTickerCmd = &cobra.Command{
		Use:   "extract_ticker",
		Short: "Crawl IDX ticker",
		Long:  "Start crawling latest ticker. Currently only support IDX. This system will update information in ./data/tickers/idx.csv",
		Run:   startCrawlingTicker,
	}
)

func init() {
	rootCmd.AddCommand(crawlTickerCmd)
}

type ListedCompany struct {
	Ticker string `json:"KodeEmiten"`
	Name   string `json:"NamaEmiten"`
}

func startCrawlingTicker(cmd *cobra.Command, args []string) {
	fmt.Println("IDX is protected by Cloudflare and for time being, we couldn't collect the information programatically. Please manually update from http://www.idx.co.id/umbraco/Surface/Helper/GetEmiten?emitenType=s and update ./data/tickers/idx.json")
}
