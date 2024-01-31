package uti

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const NA = "N/A"

const validAccessionRegExp = `^\d{10}$`

var ValidTestNames = []string{
	"KA", "PA", "KO", "KP", "ECC", "EC",
	"SM", "PV", "PM", "Efm", "Efs", "SS",
	"SA", "CO", "CA", "GBS", "SE", "AS",
	"San", "CF", "AU", "CK", "CU", "MM",
	"AB", "Pag", "PS", "ic",
}

var InterpretationMap = map[string][]int{
	"KA":  {22, 28, 33},
	"PA":  {22, 31, 35},
	"KO":  {22, 31, 35},
	"KP":  {18, 25, 31},
	"EC":  {21, 28, 30},
	"SM":  {21, 25, 29},
	"PV":  {21, 25, 30},
	"PM":  {20, 28, 29},
	"SS":  {20, 26, 29},
	"SA":  {20, 28, 30},
	"CO":  {24, 29, 35},
	"CA":  {23, 28, 34},
	"SE":  {21, 25, 30},
	"CF":  {18, 24, 27},
	"AU":  {19, 26, 30},
	"CK":  {22, 27, 33},
	"CU":  {23, 29, 35},
	"MM":  {17, 24, 32},
	"AB":  {24, 30, 33},
	"PS":  {18, 24, 27},
	"Efm": {22, 28, 35},
	"Efs": {20, 26, 28},
	"ECC": {20, 26, 28},
	"GBS": {21, 27, 31},
}

func isValidAccession(accession string) bool {
	match, err := regexp.Match(validAccessionRegExp, []byte(accession))
	if err != nil {
		log.Fatalln(fmt.Errorf("invalid accession regexp %s: %w", validAccessionRegExp, err))
	}
	return match
}

func interpretUINT(accession *accession) string {

	var isCoCaFound bool
	var dhDmCounter int
	var isNotCoCaFound bool

	for _, test := range accession.Panel1.Tests {
		if isDhDmResult(test.Result) {
			dhDmCounter++
			if isCoCaTest(test.Name) {
				isCoCaFound = true
			} else {
				isNotCoCaFound = true
			}
		}
	}

	for _, test := range accession.Panel2.Tests {
		if isDhDmResult(test.Result) {
			dhDmCounter++
			if isCoCaTest(test.Name) {
				isCoCaFound = true
			} else {
				isNotCoCaFound = true
			}
		}
	}

	for _, test := range accession.Panel3.Tests {
		if isDhDmResult(test.Result) {
			dhDmCounter++
			if isCoCaTest(test.Name) {
				isCoCaFound = true
			} else {
				isNotCoCaFound = true
			}
		}
	}

	if isCoCaFound && !isNotCoCaFound {
		return "UINT,PC"
	}

	if isCoCaFound && isNotCoCaFound {
		return "UINT,P"
	}

	if dhDmCounter >= 3 {

		for i := range accession.Panel1.Tests {
			accession.Panel1.Tests[i].Result = "."
		}

		for i := range accession.Panel2.Tests {
			accession.Panel2.Tests[i].Result = "."
		}

		for i := range accession.Panel3.Tests {
			accession.Panel3.Tests[i].Result = "."
		}

		return `UINT,\1167`
	}

	if dhDmCounter == 1 || dhDmCounter == 2 {
		return "UINT,P"
	}

	return "UINT,N"
}

func isCoCaTest(name string) bool {
	switch name {
	case "CO", "CA", "AU", "CU":
		return true
	default:
		return false
	}
}

func isDhDmResult(result string) bool {
	switch result {
	case "DH", "DM":
		return true
	default:
		return false
	}
}

func ReadCSV(filename string) ([][]string, error) {

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

func isValidTestName(testName string) bool {
	for _, name := range ValidTestNames {
		if testName == name {
			return true
		}
	}
	return false
}

func parseTestValue(value string) (int, error) {

	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}

	return int(floatValue), nil
}

func interpret(name string, value int) string {
	cutOff := InterpretationMap[name]
	if value < cutOff[0] {
		return "DH"
	}
	if value < cutOff[1] {
		return "DM"
	}
	if value < cutOff[2] {
		return "DNS"
	}
	return "ND"
}
