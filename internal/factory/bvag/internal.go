package bvag

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
	"Mega1": {operationLessOrEqual, 30},
	"Lacto": {operationGreaterOrEqual, 3},
	"BF":    {operationLessOrEqual, 32},
	"GV":    {operationGreaterOrEqual, 3},
	"BVAB2": {operationLessOrEqual, 38},
	"AV":    {operationGreaterOrEqual, 3},
	"Mob":   {operationLessOrEqual, 32},
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
