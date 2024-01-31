package candida

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
	"CK":  {"CCK", operationLessOrEqual, 32},
	"CG":  {"CG", operationLessOrEqual, 36},
	"CD":  {"CD", operationLessOrEqual, 34},
	"CP":  {"CCP", operationLessOrEqual, 33},
	"CTp": {"CTp", operationLessOrEqual, 34},
	"CA":  {"CCA", operationLessOrEqual, 37},
	"CL":  {"CL", operationLessOrEqual, 34},
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
