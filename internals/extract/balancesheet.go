package extract

var (
	SheetCodeGeneralIndustry  = "1210000"
	SheetCodeGeneralIndustry2 = "1220000"
)

func (r statementReader) ReadBalanceSheet() (res interface{}, err error) {
	for _, sheet := range r.file.GetSheetList() {
		if sheet == SheetCodeGeneralIndustry {
			res, err = r.readBalanceSheetTypeOfGeneral()
		}
	}

	return
}
