package av

import (
	"GoCFX/internal/logging"
	"fmt"
	"strconv"
)

type Document struct {
	originalFilename string
	isValid          bool
	panels           []panel
}

type panel struct {
	accession string
	isValid   bool
	icValid   bool
	tests     []test
}

type test struct {
	name   string
	result string
}

type headersInfo struct {
	headersMap map[int]string
	icIndex    int
}

type cutOff struct {
	operation int
	value     float64
}

func New(logger *logging.Logger, filename string) (Document, error) {

	d := Document{
		originalFilename: filename,
		isValid:          true,
	}

	matrix, err := readCSV(filename)
	if err != nil {
		return d, err
	}

	hInfo, err := newHeadersInfo(matrix[1])
	if err != nil {
		return d, err
	}

	for _, line := range matrix[2:] {
		if isDataLine(line) {
			p, err := newPanel(logger, hInfo, line)
			if err != nil {
				logger.Errorf("unable to create opRpp panel %s:", err.Error())
				continue
			}
			d.panels = append(d.panels, p)
		}
	}
	return d, nil
}

func newPanel(logger *logging.Logger, hInfo headersInfo, line []string) (panel, error) {

	const (
		minFieldCount = 11
	)

	if len(line) < minFieldCount {
		return panel{}, fmt.Errorf("invalie fields count")
	}

	p := panel{
		accession: line[3],
	}

	for index, name := range hInfo.headersMap {

		value := line[index+1]
		t, err := newTest(name, value)
		if err != nil {
			logger.Errorf("create test error: %s", err.Error())
		}

		p.tests = append(p.tests, t)
	}

	icValue := line[hInfo.icIndex+1]
	if !isValidIc(icValue) {
		return panel{}, fmt.Errorf("IC cutoff validation failed")
	}

	return p, nil
}

func isValidIc(value string) bool {

	const icCutOff = 40.0

	if value == "N/A" {
		return false
	}

	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return false
	}

	if floatValue > icCutOff {
		return false
	}

	return true
}

func newTest(name string, value string) (test, error) {

	const (
		detected    = "Detected"
		notDetected = "Not Detected"
	)

	t := test{
		name: name,
	}

	co, ok := cutOffMap[name]
	if !ok {
		return t, fmt.Errorf("no cutOff found for test %s", name)
	}

	if value == "N/A" {
		t.result = notDetected
		return t, nil
	}

	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return t, err
	}

	switch co.operation {
	case operationGreaterOrEqual:
		if floatValue >= co.value {
			t.result = detected
		} else {
			t.result = notDetected
		}
	case operationLessOrEqual:
		if floatValue <= co.value {
			t.result = detected
		} else {
			t.result = notDetected
		}
	default:
		return t, fmt.Errorf("invalid operation detected in cutOff [%d]", co.operation)
	}

	return t, nil
}

func newHeadersInfo(line []string) (headersInfo, error) {

	const (
		dataFieldStartIndex = 5
		minFieldCount       = 11
	)

	hi := headersInfo{
		headersMap: make(map[int]string),
	}

	if len(line) < minFieldCount {
		return hi, fmt.Errorf("minimum field count constraint violeited")
	}

	for i, headerName := range line[dataFieldStartIndex:] {

		if i%2 == 1 {
			continue
		}

		if headerName == "IC" {
			hi.icIndex = i + dataFieldStartIndex
			break
		}

		hi.headersMap[i+dataFieldStartIndex] = headerName
	}

	return hi, nil
}
