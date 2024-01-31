package opRpp

import (
	"GoCFX/internal/logging"
	"encoding/csv"
	"fmt"
	"os"
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

func New(logger *logging.Logger, filename string) (Document, error) {

	var d Document

	d.originalFilename = filename

	d.isValid = true

	headers := make(map[int]string)

	f, err := os.Open(filename)
	if err != nil {
		return d, err
	}

	defer f.Close()

	cr := csv.NewReader(f)
	matrix, err := cr.ReadAll()
	if err != nil {
		return d, err
	}

	for i, line := range matrix {
		if i == 1 {
			headers, err = parseHeaders(line)
			if err != nil {
				return d, err
			}
			continue
		}

		if isDataLine(line) {
			p, err := newPanel(logger, headers, line)
			if err != nil {
				logger.Errorf("unable to create opRpp panel %s:", err.Error())
			}

			if p.icValid {
				d.panels = append(d.panels, p)
			} else {
				logger.Errorf("invalid ic found for file %s at line %d", filename, i)
			}
		}
	}
	return d, nil
}

func newPanel(logger *logging.Logger, headers map[int]string, line []string) (panel, error) {

	const (
		detected    = "detected"
		notDetected = "not detected"
	)

	if len(line) < 11 {
		return panel{}, fmt.Errorf("invalie fields count")
	}

	p := panel{
		accession: line[3],
	}

	var (
		co  cutOff
		ok  bool
		v   float64
		err error
	)

	for index, name := range headers {

		co, ok = cutOffMap[name]

		if name == "IC" {
			v, err = strconv.ParseFloat(line[index+1], 64)
			if err != nil {
				continue
			}

			if v < co.value {
				p.icValid = true
				continue
			}
		}

		if !ok {
			return panel{}, fmt.Errorf("no cutOff found for test %s", name)
		}

		if line[index+1] == "N/A" {
			p.tests = append(p.tests, test{
				name:   name,
				result: notDetected,
			})
			continue
		}

		v, err = strconv.ParseFloat(line[index+1], 64)
		if err != nil {
			return panel{}, err
		}

		if v > co.value {
			p.tests = append(p.tests, test{
				name:   name,
				result: detected,
			})
		}

		p.tests = append(p.tests, test{
			name:   name,
			result: notDetected,
		})
	}

	return p, nil
}
