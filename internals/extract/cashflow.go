package extract

// 1210000
type CFGeneralIndustry struct {
}

type FromOperatingActivities struct {
	ReceiptsFromCustomers                 int `loc:"B7"`  // Receipts from customers
	ReceiptsOfSubsidy                     int `loc:"B8"`  // Receipts of subsidy
	ReceiptsFromRoyaltiesAndOtherRevenues int `loc:"B9"`  // Receipts from royalties, fees, commissions and other revenues
	ReceiptsFromContractsHeld             int `loc:"B10"` // Receipts from contracts held-for-dealing or trading purposes
}
