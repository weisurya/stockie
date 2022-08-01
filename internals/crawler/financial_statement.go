package crawler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"weisurya/stockie/app/cmd/helper"

	"github.com/Jeffail/tunny"
	"github.com/sirupsen/logrus"
)

type Company struct {
	Ticker       string `json:"KodeEmiten"`
	Name         string `json:"NamaEmiten"`
	ListingDate  string
	Share        int
	ListingBoard string
}

const (
	idxFilePath                   = "./data/tickers/idx.json"
	financialStatementCSVFilename = "statement.xlsx"
)

func (s *Service) getAllIDXTickers(filePath string) (res []Company, err error) {
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}

	if err = json.Unmarshal([]byte(f), &res); err != nil {
		return
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

	baseUrl, _ := url.Parse("https://www.idx.co.id/Portals/0/StaticData/ListedCompanies/Corporate_Actions/New_Info_JSX/Jenis_Informasi/01_Laporan_Keuangan/02_Soft_Copy_Laporan_Keuangan//Laporan Keuangan Tahun ")
	baseUrl.Path += fmt.Sprintf("%v/%v/%v/FinancialStatement-%v-%v-%v.xlsx", year, tw, ticker, year, q, ticker)

	baseUrl.RawQuery = baseUrl.Query().Encode()

	logrus.Info(baseUrl.String())

	return baseUrl.String()

}

func downloadFinancialStatement(url, ticker string, year, quarter int) error {
	filePath := fmt.Sprintf("./data/financial_statements/%v/%v/%v/", ticker, strconv.Itoa(year), strconv.Itoa(quarter))
	fullPath := filePath + financialStatementCSVFilename

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if err = os.MkdirAll(filePath, os.ModePerm); err != nil {
			return err
		}
	}

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		logrus.Infof("Downloading %v %v-%v financial statement...", ticker, year, quarter)

		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}

		req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.80 Safari/537.36")

		resp, err := client.Do(req)
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
		} else if resp.StatusCode == http.StatusForbidden {
			helper.Log.Error("%v %v-%v financial statement was restricted to be downloaded", ticker, year, quarter)
			helper.Log.Error(url)
		} else {
			// logrus.Infof("%v %v-%v financial statement has not existed yet", ticker, year, quarter)
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

func (s *Service) downloadAllStatements(companies []Company, startYear, endYear, endQuarter int) error {
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
					panic(err)
				}
			}
		}
	}

	return nil
}
