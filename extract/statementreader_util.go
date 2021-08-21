package extract

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

const (
	tagSheetLoc = "loc"
	tagIsTitle  = "is_title"

	fieldTotal = "Total"
)

type statementReader struct {
	file *excelize.File
}

func NewStatementReader(filePath string) (r statementReader, err error) {
	r.file, err = excelize.OpenFile(filePath)
	return
}

func (r statementReader) getEntityName() string {
	ticker, _ := r.file.GetCellValue("1000000", "B5")

	return ticker
}

func (r statementReader) getTicker() string {
	ticker, _ := r.file.GetCellValue("1000000", "B7")

	return ticker
}

func (r statementReader) getMainIndustry() string {
	ticker, _ := r.file.GetCellValue("1000000", "B9")

	return ticker
}

func (r statementReader) getSector() string {
	ticker, _ := r.file.GetCellValue("1000000", "B10")

	return ticker
}

func (r statementReader) getSubSector() string {
	ticker, _ := r.file.GetCellValue("1000000", "B11")

	return ticker
}

func (r statementReader) getPeriod() (startDate, endDate string) {
	startDate, _ = r.file.GetCellValue("1000000", "B18")
	endDate, _ = r.file.GetCellValue("1000000", "B19")

	return startDate, endDate
}

func (r statementReader) extractNumber(sheet, axis string) int {
	res, err := r.file.GetCellValue(sheet, axis)
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

func (r statementReader) setValueToSubjectFromSheet(subject interface{}, sheet string) error {
	v := reflect.ValueOf(subject).Elem()
	if !v.CanAddr() {
		return fmt.Errorf("cannot assign to the item passed, item must be a pointer in order to assign")
	}

	findTagName := func(t reflect.StructTag) string {
		if res, ok := t.Lookup(tagSheetLoc); ok {
			return res
		}

		return ""
	}

	isTitle := func(t reflect.StructTag) bool {
		if _, ok := t.Lookup(tagIsTitle); ok {
			return true
		}

		return false
	}

	total := 0

	for i := 0; i < v.NumField(); i++ {
		typeField := v.Type().Field(i)

		if isTitle(typeField.Tag) {
			continue
		}

		tagName := findTagName(typeField.Tag)

		value := r.extractNumber(sheet, tagName)

		fieldVal := v.Field(i)
		fieldVal.Set(reflect.ValueOf(value))

		total += value

		if typeField.Name == fieldTotal {
			fieldVal.Set(reflect.ValueOf(total))
		}
	}

	return nil
}
