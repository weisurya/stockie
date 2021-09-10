package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"weisurya/stockie/internals/extract"

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
	temp := map[string][]string{}

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

				for _, tabName := range reader.GetAllTabs() {
					temp[tabName] = append(temp[tabName], tickerDir.Name())
					// switch tabName {
					// case extract.SheetCodeGeneralIndustry:

					// }
				}
				break

				// bs, err := reader.ReadBalanceSheet()
				// if err != nil {
				// 	fmt.Println(err)
				// 	os.Exit(1)
				// }

				// data, _ := utils.CustomMarshaller{
				// 	Value: bs,
				// }.MarshalWithNameTag()
				// fmt.Println(string(data))

				// os.Exit(0)
			}

			break
		}
	}

	for key, value := range temp {
		if len(value) > 5 {
			value = value[:5]
		}
		fmt.Printf("%v: %v\n", key, value)
	}

}

var keyMatchRegex = regexp.MustCompile(`\"(\w+)\":`)
var wordBarrierRegex = regexp.MustCompile(`(\w)([A-Z])`)

type marshaller struct {
	Value interface{}
}

func (c marshaller) MarshalJSON() ([]byte, error) {
	marshalled, err := json.Marshal(c.Value)

	converted := keyMatchRegex.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			return bytes.ToLower(wordBarrierRegex.ReplaceAll(
				match,
				[]byte(`${1}_${2}`),
			))
		},
	)

	return converted, err
}
