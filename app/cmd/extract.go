package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"weisurya/stockie/extract"

	"github.com/spf13/cobra"
)

var (
	extractCmd = &cobra.Command{
		Use: "extract",
		Run: run,
	}
)

func init() {
	rootCmd.AddCommand(extractCmd)
}

func run(cmd *cobra.Command, args []string) {

	tickerDirs, err := ioutil.ReadDir("./data/financial_statements")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, tickerDir := range tickerDirs {
		annualDirs, err := ioutil.ReadDir(fmt.Sprintf("./data/financial_statements/%v", tickerDir.Name()))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, annualDir := range annualDirs {

			quarterDirs, err := ioutil.ReadDir(fmt.Sprintf("./data/financial_statements/%v/%v", tickerDir.Name(), annualDir.Name()))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			for _, quarterDir := range quarterDirs {

				reader, err := extract.NewStatementReader(fmt.Sprintf("./data/financial_statements/%v/%v/%v/statement.xlsx", tickerDir.Name(), annualDir.Name(), quarterDir.Name()))
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				fmt.Printf("Reading: %v %v-%v...\n", tickerDir.Name(), annualDir.Name(), quarterDir.Name())

				bs, err := reader.ReadBalanceSheetTypeOfGeneral()
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				b, _ := json.MarshalIndent(bs, "", "  ")
				fmt.Println(string(b))

				os.Exit(0)
			}
		}
	}

}
