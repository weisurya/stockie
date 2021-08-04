package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
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

// 121000
type BalanceSheet struct {
	Assets Assets
}

type Assets struct {
	CurrentAssets    CurrentAssets
	NonCurrentAssets NonCurrentAssets
	Total            int // Total assets - B123
}

type CurrentAssets struct {
	CashAndCashEquivalent int // Cash and cash equivalents - B7
	TradeReceivables      TradeReceivables
	OtherReceivables      OtherReceivables
	Inventories           int // Current inventories - B41
	Others                OtherCurrentAssets
	Total                 int // Total current assets - B56
}

type TradeReceivables struct {
	ThirdParties   int // Trade receivables third parties - B18
	RelatedParties int // Trade receivables related parties - B19
	Total          int
}

type OtherReceivables struct {
	ThirdParties   int // Other receivables third parties - B36
	RelatedParties int // Other receivables related parties - B37
	Total          int
}

type OtherCurrentAssets struct {
	Advances           int // Other current advances - B47
	PrepaidTaxes       int // Current prepaid taxes - 48
	NonFinancialAssets int // Other current non-financial assets - B53
	Total              int
}

type NonCurrentAssets struct {
	InvestmentsInJointVentures int // Investments in joint ventures - B72
	DeferredTaxAssets          int // Deferred tax assets - B86
	Plantations                Plantations
	PropertyPlanEquipment      int // Property, plant and equipment - B101
	Goodwill                   int // Goodwill - B118
	Others                     OtherNonCurrentAssets
	Total                      int // Total non-current assets - B122
}

type Plantations struct {
	AssetsMature   int // Plantation assets mature - B96
	AssetsImmature int // Plantation assets immature -- B97
	Plasma         int // Plasma plantations -- B98
}

type OtherNonCurrentAssets struct {
	ReceivablesRelatedParties int // Other non-current receivables related parties -- B68
	ClaimsForTaxRefund        int // Non-current claims for tax refund -- B116
	NonFinancialAssets        int // Other non-current non-financial assets - B121
}

type LiabilitiesAndEquity struct {
	TradePayables TradePayables
}

type TradePayables struct {
	ThirdParties      int // Trade payables third parties - B130
	RelatedParties    int // Trade payables related parties - B131
	OtherThirdParties int // Other payables third parties - 133
	Total             int
}

func run(cmd *cobra.Command, args []string) {
	f, err := excelize.OpenFile("./data/financial_statements/AALI/2016/1/statement.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	balanceSheet := BalanceSheet{
		Assets: Assets{
			CurrentAssets: CurrentAssets{
				CashAndCashEquivalent: extractNumber(f, "1210000", "B7"),
				TradeReceivables: TradeReceivables{
					ThirdParties:   extractNumber(f, "1210000", "B18"),
					RelatedParties: extractNumber(f, "1210000", "B19"),
					Total:          extractNumber(f, "1210000", "B18") + extractNumber(f, "1210000", "B19"),
				},
			},
		},
	}

	fmt.Println(balanceSheet)

}

func extractNumber(f *excelize.File, sheet, axis string) int {
	res, err := f.GetCellValue(sheet, axis)
	if err != nil {
		return 0
	}

	res = strings.ReplaceAll(res, ".0", "")

	n, err := strconv.Atoi(res)
	if err != nil {
		return 0
	}

	return n
}
