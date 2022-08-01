package crawler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type DailyTrade struct {
	ID                   int     `json:"Id"`
	Name                 string  `json:"Name"`
	Code                 string  `json:"Code"`
	StockSubSectorId     int     `json:"StockSubSectorId"`
	SubSectorName        string  `json:"SubSectorName"`
	StockSectorId        int     `json:"StockSectorId"`
	SectorName           string  `json:"SectorName"`
	NewSubIndustryId     int     `json:"NewSubIndustryId"`
	NewSubIndustryName   string  `json:"NewSubIndustryName"`
	NewIndustryId        int     `json:"NewIndustryId"`
	NewIndustryName      string  `json:"NewIndustryName"`
	NewSubSectorId       int     `json:"NewSubSectorId"`
	NewSubSectorName     string  `json:"NewSubSectorName"`
	NewSectorId          int     `json:"NewSectorId"`
	NewSectorName        string  `json:"NewSectorName"`
	Last                 float64 `json:"Last"`
	PrevClosingPrice     float64 `json:"PrevClosingPrice"`
	AdjustedClosingPrice float64 `json:"AdjustedClosingPrice"`
	AdjustedOpenPrice    float64 `json:"AdjustedOpenPrice"`
	AdjustedHighPrice    float64 `json:"AdjustedHighPrice"`
	AdjustedLowPrice     float64 `json:"AdjustedLowPrice"`
	Volume               float64 `json:"Volume"`
	Frequency            float64 `json:"Frequency"`
	Value                float64 `json:"Value"`
	OneDay               float64 `json:"OneDay"`
	OneWeek              float64 `json:"OneWeek"`
	OneMonth             float64 `json:"OneMonth"`
	ThreeMonth           float64 `json:"ThreeMonth"`
	SixMonth             float64 `json:"SixMonth"`
	OneYear              float64 `json:"OneYear"`
	ThreeYear            float64 `json:"ThreeYear"`
	FiveYear             float64 `json:"FiveYear"`
	TenYear              float64 `json:"TenYear"`
	Mtd                  float64 `json:"Mtd"`
	Yth                  float64 `json:"Yth"`
	Per                  float64 `json:"Per"`
	Pbr                  float64 `json:"Pbr"`
	Capitalization       float64 `json:"Capitalization"`
	BetaOneYear          float64 `json:"BetaOneYear"`
	StdevOneYear         float64 `json:"StdevOneYear"`
	LastDate             string  `json:"LastDate"`
	LastUpdate           string  `json:"LastUpdate"`
	Roe                  float64 `json:"Roe"`
}

var (
	baseFolderPathDailyTrade = "./data/daily_trade"
	fullDailyTradeFileName   = "_full.json"
)

func (s *Service) isTodayDataDownloaded(date time.Time) bool {
	filePath := fmt.Sprintf("%v/%v/%v/%v", baseFolderPathDailyTrade, date.Format("2006"), date.Format("01"), date.Format("02"))

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

func (s *Service) downloadDailyTrade(date string) (res []DailyTrade, err error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://pasardana.id/api/StockSearchResult/GetAll?date="+date, nil)
	if err != nil {
		return
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Referer", "https://pasardana.id/stock/search")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36")

	logrus.Infof("[Downloading...] Daily trade information for %v", date)

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		resBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}
		logrus.Error(string(resBody))

		return res, errors.New(http.StatusText(resp.StatusCode))
	}

	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return
	}

	logrus.Infof("[Downloaded] Daily trade information for %v", date)

	return
}

func (s *Service) isStockMarketOpen(date time.Time, data []DailyTrade) (bool, error) {
	dt, err := time.Parse("2006-01-02T15:04:05", data[0].LastDate)
	if err != nil {
		return false, err
	}

	d1 := dt.Format("2006-01-02")
	d2 := date.Format("2006-01-02")

	return d1 == d2, nil
}

func (s *Service) storeDailyTradeIntoCSV(date time.Time, data []DailyTrade) error {
	filePath := fmt.Sprintf("%v/%v/%v/%v", baseFolderPathDailyTrade, date.Format("2006"), date.Format("01"), date.Format("02"))

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if err = os.MkdirAll(filePath, os.ModePerm); err != nil {
			return err
		}
	}

	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(fmt.Sprintf("%v/%v", filePath, fullDailyTradeFileName), file, 0777); err != nil {
		return err
	}

	for _, dailyTrade := range data {
		filePath := fmt.Sprintf("%v/%v.json", filePath, dailyTrade.Code)

		file, err = json.MarshalIndent(dailyTrade, "", " ")
		if err != nil {
			return err
		}

		if err = ioutil.WriteFile(filePath, file, 0777); err != nil {
			return err
		}
	}

	return nil
}
