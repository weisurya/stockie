package crawler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"time"

	"github.com/sirupsen/logrus"
)

type CorporateAction struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Title       string `json:"title"`
	Start       string `json:"start"`
	Agenda      string `json:"agenda"`
	Type        string `json:"jenis"`
	Location    string `json:"location"`
	FKStockID   int    `json:"fkStockId"`
}

var (
	baseFolderPathCorpActions = "./data/corporate_actions"
	fullCorpActionFileName    = "corporate_action.json"
)

func (s *Service) downloadCorporateAction(isDownloadAll bool, date string) (res []CorporateAction, err error) {
	url := "https://pasardana.id/api/IdxCalendarCorporateAction/GetFullCalendar"

	if !isDownloadAll {
		url += fmt.Sprintf("?monthDate=%v", date)
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return res, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Referer", "https://pasardana.id/stock/search")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36")

	logrus.Infof("[Downloading...] corporate action information | All: %v | Month: %v", isDownloadAll, date)

	resp, err := client.Do(req)
	if err != nil {
		return res, err
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

	logrus.Infof("[Downloaded] corporate action information | All: %v | Month: %v", isDownloadAll, date)

	return
}

func (s *Service) createCorporateActionMap(data []CorporateAction) (res map[string]map[string]map[string][]CorporateAction, err error) {
	res = make(map[string]map[string]map[string][]CorporateAction)

	for _, ca := range data {
		var date time.Time
		date, err = time.Parse("2006-01-02T15:04:05", ca.Start)
		if err != nil {
			return res, err
		}

		caMonth := map[string][]CorporateAction{}
		caMonth[date.Format("01")] = append(caMonth[date.Format("01")], ca)

		caYear := map[string]map[string][]CorporateAction{}
		caYear[date.Format("2006")] = caMonth

		res[ca.Title] = caYear
	}

	for _, ca := range res {
		for _, ca2 := range ca {
			for _, ca3 := range ca2 {
				sort.Slice(ca3, func(i, j int) bool {
					return ca3[i].Start < ca3[j].Start
				})
			}
		}
	}

	return res, nil
}

func (s *Service) storeCorporateActionIntoJSON(
	data map[string]map[string]map[string][]CorporateAction,
) error {
	for ticker, data2 := range data {
		for year, data3 := range data2 {
			for month, corporateActions := range data3 {
				filePath := fmt.Sprintf("%v/%v/%v/%v", baseFolderPathCorpActions, ticker, year, month)

				if _, err := os.Stat(filePath); os.IsNotExist(err) {
					if err = os.MkdirAll(filePath, os.ModePerm); err != nil {
						return err
					}
				}

				file, err := json.MarshalIndent(corporateActions, "", " ")
				if err != nil {
					return err
				}

				if err = ioutil.WriteFile(fmt.Sprintf("%v/%v", filePath, fullCorpActionFileName), file, 0777); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
