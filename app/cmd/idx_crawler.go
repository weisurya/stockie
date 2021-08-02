package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/Jeffail/tunny"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	idxCrawlerCmd = &cobra.Command{
		Use:   "crawl",
		Short: "Crawl IDX financial statement",
		Long:  "Start crawling financial statements. Currently only support IDX. This system refers to the list of tickers in ./data/tickers/idx.csv",
		Run:   startCrawling,
	}
)

var (
	startYear int
	endYear   int
	quarter   int
)

func init() {
	rootCmd.AddCommand(idxCrawlerCmd)

	idxCrawlerCmd.Flags().IntVarP(
		&startYear,
		"startyear",
		"s",
		2016,
		"[Required] Start year",
	)
	idxCrawlerCmd.MarkFlagRequired("startyear")

	idxCrawlerCmd.Flags().IntVarP(
		&endYear,
		"endyear",
		"e",
		2021,
		"[Required] End year",
	)
	idxCrawlerCmd.MarkFlagRequired("endyear")

	idxCrawlerCmd.Flags().IntVarP(
		&quarter,
		"endquarter",
		"q",
		2,
		"[Required] End quarter",
	)
	idxCrawlerCmd.MarkFlagRequired("endquarter")
}

type Company struct {
	Ticker       string
	Name         string
	ListingDate  string
	Share        int
	ListingBoard string
}

const (
	idxFilePath                   = "./data/tickers/idx.csv"
	financialStatementCSVFilename = "statement.xlsx"
)

func startCrawling(cmd *cobra.Command, args []string) {
	companies, err := getAllIDXTickers(idxFilePath)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	if err = downloadAllStatements(companies, startYear, endYear, quarter); err != nil {
		fmt.Println(err)
		panic(err)
	}

	logrus.Infof("Done pulling IDX financial statements from %v to %v-%v", startYear, endYear, quarter)
}

func getAllIDXTickers(filePath string) (res []Company, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}

	defer f.Close()

	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return
	}

	for i, v := range records {
		if i == 0 {
			continue
		}

		shares, _ := strconv.Atoi(strings.ReplaceAll(v[4], ",", ""))

		res = append(res, Company{
			Ticker:       v[1],
			Name:         v[2],
			ListingDate:  v[3],
			Share:        shares,
			ListingBoard: v[5],
		})
	}

	return res, nil
}

func createFinancialStatementURL(ticker string, year, quarter int) string {
	q := ""

	for i := 1; i <= quarter; i++ {
		q += "I"
		if quarter == 4 {
			q = "Tahunan"
		}
	}

	tw := func() (res string) {
		switch quarter {
		case 1:
			res = "TW1"
		case 2:
			res = "TW2"
		case 3:
			res = "TW3"
		case 4:
			res = "Audit"
		}

		return res
	}()

	return "https://www.idx.co.id/Portals/0/StaticData/ListedCompanies/Corporate_Actions/New_Info_JSX/Jenis_Informasi/01_Laporan_Keuangan/02_Soft_Copy_Laporan_Keuangan//Laporan%20Keuangan%20Tahun%20" + fmt.Sprintf("%v/%v/%v/FinancialStatement-%v-%v-%v.xlsx", year, tw, ticker, year, q, ticker)

}

func downloadFinancialStatement(url, ticker string, year, quarter int) error {
	filePath := fmt.Sprintf("./data/financial_statements/%v/%v/%v/", ticker, strconv.Itoa(year), strconv.Itoa(quarter))

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if err = os.MkdirAll(filePath, os.ModePerm); err != nil {
			return err
		}
	}

	fullPath := filePath + financialStatementCSVFilename
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		logrus.Infof("Downloading %v %v-%v financial statement...", ticker, year, quarter)

		resp, err := http.Get(url)
		if err != nil {
			return err
		}

		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			logrus.Infof("Downloaded! Storing information in ./data/financial_statements/%v/%v/%v", ticker, year, quarter)

			out, err := os.Create(filePath + financialStatementCSVFilename)
			if err != nil {
				return err
			}

			defer out.Close()

			if _, err := io.Copy(out, resp.Body); err != nil {
				return err
			}

			logrus.Infof("File stored in ./data/financial_statements/%v/%v/%v", ticker, year, quarter)
		} else {
			logrus.Infof("%v %v-%v financial statement has not existed yet", ticker, year, quarter)

			if err = os.Remove(filePath); err != nil {
				return err
			}
		}

	}

	return nil
}

type ConnDownloader struct {
	Ticker  string
	Year    int
	Quarter int
}

func downloadAllStatements(companies []Company, startYear, endYear, endQuarter int) error {
	pool := tunny.NewFunc(runtime.NumCPU(), func(payload interface{}) interface{} {
		data := payload.(ConnDownloader)

		err := downloadFinancialStatement(
			createFinancialStatementURL(data.Ticker, data.Year, data.Quarter),
			data.Ticker,
			data.Year,
			data.Quarter,
		)

		return err
	})

	defer pool.Close()

	for _, v := range companies {
		for year := startYear; year <= endYear; year++ {
			maxQuarter := 4
			if year == endYear {
				maxQuarter = endQuarter
			}

			for quarter := 1; quarter <= maxQuarter; quarter++ {
				res := pool.Process(ConnDownloader{
					Ticker:  v.Ticker,
					Year:    year,
					Quarter: quarter,
				})

				var err error
				if res != nil {
					err = res.(error)
				}

				if err != nil {
					fmt.Println(err)
					panic(err)
				}
			}
		}
	}

	return nil
}
