# Stockie

Stockie is a CLI-based system to automate data crawling, data clean, data pre-processing, and data analysis for IDX stock market.

## Getting Started

* Make sure you have Go 1.12+ CLI installed in your system
* Clone this repository
* Run `go run app/main.go -h` to read further documentation

## User journey

* [ ] As a user, I want to refresh list IDX tickers data via CLI, so I could use it for financial statement crawling process
* [X] As a user, I want to crawl financial statements of all IDX tickers via CLI, so I don't need to download it manually
* [ ] As a user, I want to clean up XLSX-based financial statement into SQL-based information, so I don't need to copy n' paste the number 1-by-1
* [ ] As a user, I want to have a capability to pre-process quarter-based financial statement into TTM financial statement, so I could use it for stock analysis
