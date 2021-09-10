package stockie

type ICrawlerService interface {
	StartCrawlCorporateActions(isDownloadAll bool, startMonth, endMonth string) error
	StartCrawlDailyTrade(startDate, endDate string) error
	StartCrawlFinancialStatement(startYear, endYear, quarter int) error
}
