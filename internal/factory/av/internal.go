package av

import (
	"encoding/csv"
	"os"
	"regexp"
)

const (
	operationLessOrEqual = iota
	operationGreaterOrEqual
)

var cutOffMap = map[string]cutOff{
	"EC":    {operationLessOrEqual, 28},
	"GBS":   {operationLessOrEqual, 29},
	"SA":    {operationLessOrEqual, 28},
	"GAS":   {operationLessOrEqual, 30},
	"EF":    {operationLessOrEqual, 30},
	"LR":    {operationLessOrEqual, 29},
	"Lacto": {operationGreaterOrEqual, 3},
}

func isDataLine(line []string) bool {
	rx := regexp.MustCompile(`^\d{10}$`)
	return rx.MatchString(line[3])
}

func readCSV(filename string) ([][]string, error) {

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	csvR := csv.NewReader(f)

	csvR.Comma = ','
	csvR.TrimLeadingSpace = true
	csvR.Comment = '#'

	all, err := csvR.ReadAll()
	if err != nil {
		return nil, err
	}

	return all, nil
}
