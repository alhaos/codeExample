package utiw

import (
	"GoCFX/internal/logging"
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var cutOffMap = map[string]float64{
	"KA":  34.0,
	"PA":  30.0,
	"KO":  33.0,
	"KP":  33.0,
	"ECC": 31.0,
	"EC":  32.0,
	"SM":  32.0,
	"PV":  30.0,
	"PM":  35.0,
	"Efm": 32.0,
	"Efs": 31.0,
	"SS":  35.0,
	"SA":  31.0,
	"CO":  35.0,
	"CA":  33.0,
	"GBS": 30.0,
	"SE":  23.0,
	"CF":  31.0,
	"CK":  35.0,
	"MM":  33.0,
	"AB":  36.0,
	"PS":  35.0,
}

func readCSV(filename string) ([][]string, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	csvReader := csv.NewReader(file)

	m, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return m, nil
}

func newHeaderInfo(line []string) (headerInfo, error) {

	const (
		firstHeaderIndex        = 5
		icZeroBasedIndexFromEnd = 3
		minFieldsCount          = 11
	)

	if len(line) < minFieldsCount {
		return headerInfo{}, fmt.Errorf("the number of fields is less than the required minimum")
	}

	hm := headerInfo{
		headersMap: make(map[int]string),
		icIndex:    0,
	}

	for i, s := range line[firstHeaderIndex : len(line)-(icZeroBasedIndexFromEnd+2)] {

		if i%2 == 1 {
			continue
		}

		if s == "" {
			continue
		}

		hm.headersMap[i+firstHeaderIndex] = s

	}

	hm.icIndex = len(line) - icZeroBasedIndexFromEnd - 1

	return hm, nil
}

func newPanel(logger *logging.Logger, headerLine []string, dataLine []string) (panel, error) {

	const icCutOff = 40.0

	hi, err := newHeaderInfo(headerLine)
	if err != nil {
		return panel{}, fmt.Errorf("invalid header line detected")
	}

	acc := headerLine[3]

	if !isValidAccession(acc) {
		return panel{}, fmt.Errorf("invalid accession detected %s", acc)
	}

	p := panel{
		accessionID: headerLine[3],
	}

	for index, name := range hi.headersMap {
		value := dataLine[index+1]
		t, err := newTest(name, value)
		if err != nil {
			logger.Errorf("invalid test detected: %s", err.Error())
			continue
		}
		p.tests = append(p.tests, t)
	}

	// parse float from ic raw value
	icValue := dataLine[hi.icIndex+1]
	icFloatValue, err := strconv.ParseFloat(icValue, 64)
	if err != nil {
		return panel{}, fmt.Errorf("invalid ic detectd [%s]", icValue)
	}

	// validate IC cutOff
	if icFloatValue > icCutOff {
		return p, fmt.Errorf("ic cutOff failed %2f", icFloatValue)
	}

	// return error if tests is empty
	if len(p.tests) == 0 {
		return p, fmt.Errorf("empty panel detected")
	}

	return p, nil
}

func newTest(name string, value string) (test, error) {

	const (
		detected    = "Detected"
		notDetected = "Not Detected"
	)

	t := test{
		name: name,
	}

	cutOff, ok := cutOffMap[name]
	if !ok {
		return test{}, fmt.Errorf("no cutOff found for test %s", name)
	}

	if value == "N/A" {
		t.result = notDetected
		return t, nil
	}

	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return test{}, err
	}

	if floatValue <= cutOff {
		t.result = detected
	}

	t.result = notDetected

	return t, nil

}

func isValidAccession(acc string) bool {
	rx := regexp.MustCompile(`^\d{10}$`)
	return rx.MatchString(acc)
}
