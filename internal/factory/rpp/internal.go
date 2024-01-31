package rpp

import (
	"GoCFX/internal/logging"
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var cutOffMap = map[string]float64{
	"H1":    39.0,
	"H3":    40.0,
	"pdm09": 39.0,
	"Flu A": 40.0,
	"Flu B": 39.0,
	"RSV A": 38.0,
	"RSV B": 40.0,
	"IC":    40.0,
	"AdV":   40.0,
	"HEV":   40.0,
	"MPV":   38.0,
	"PIV1":  40.0,
	"PIV2":  39.0,
	"PIV3":  39.0,
	"PIV4":  39.0,
	"HBoV":  37.0,
	"HRV":   38.0,
	"229E":  38.0,
	"NL63":  39.0,
	"OC43":  37.0,
	"CP":    37.0,
	"MP":    35.0,
	"LP":    33.0,
	"BP":    34.0,
	"BPP":   33.0,
	"SP":    33.0,
	"HI":    39.0,
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
