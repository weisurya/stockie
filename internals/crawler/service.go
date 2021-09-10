package crawler

import (
	"context"

	"github.com/sirupsen/logrus"
)

type Service struct {
	ctx context.Context
}

func New(ctx context.Context) *Service {
	return &Service{ctx: ctx}
}

func (s *Service) StartCrawlCorporateActions(
	isDownloadAll bool,
	startMonth, endMonth string,
) error {
	sd, ed, err := setDateRange(startMonth, endMonth)
	if err != nil {
		return err
	}

	for d := sd; !d.After(ed); d = d.AddDate(0, 1, 0) {
		res, err := s.downloadCorporateAction(isDownloadAll, d.Format("2006-01-02"))
		if err != nil {
			return err
		}

		ca, err := s.createCorporateActionMap(res)
		if err != nil {
			return err
		}

		if err := s.storeCorporateActionIntoJSON(ca); err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) StartCrawlDailyTrade(
	startDate, endDate string,
) error {
	sd, ed, err := setDateRange(startDate, endDate)
	if err != nil {
		return err
	}

	for d := sd; !d.After(ed); d = d.AddDate(0, 1, 0) {
		if s.isTodayDataDownloaded(d) {
			logrus.Infof("Today daily trade %v has been downloaded. Please check in %v", d.Format("2006-01-02"), baseFolderPathDailyTrade)

			continue
		}

		res, err := s.downloadDailyTrade(d.Format("2006-01-02"))
		if err != nil {
			return err
		}

		isOpen, err := s.isStockMarketOpen(d, res)
		if err != nil {
			return err
		}

		if !isOpen {
			logrus.Infof("Market is not open at %v. Skipping the process to the next day", d.Format("2006-01-02"))

			continue
		}

		if err := s.storeDailyTradeIntoCSV(d, res); err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) StartCrawlFinancialStatement(
	startYear, endYear int,
	quarter int,
) error {
	companies, err := s.getAllIDXTickers(idxFilePath)
	if err != nil {
		return err
	}

	if err = s.downloadAllStatements(companies, startYear, endYear, quarter); err != nil {
		return err
	}

	return nil
}
