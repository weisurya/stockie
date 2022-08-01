package crawler

import (
	"time"
)

func setDateRange(startDate, endDate string) (sd, ed time.Time, err error) {
	sd = time.Now()
	ed = time.Now()

	if startDate != "" {
		sd, err = time.Parse("2006-01-02", startDate)
		if err != nil {
			return
		}
	}

	if endDate != "" {
		ed, err = time.Parse("2006-01-02", endDate)
		if err != nil {
			return
		}
	}

	return
}
