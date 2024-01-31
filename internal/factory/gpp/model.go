package gpp

import (
	"GoCFX/internal/logging"
	"regexp"
)

type panel struct {
	accession       string
	isValid         bool
	tests           map[string]bool
	internalControl string
}

type Document struct {
	originalFilename string
	isValid          bool
	panels           []*panel
}

func New(logger *logging.Logger, filename string) (Document, error) {

	matrix := getData(filename)

	headers := getHeaders(matrix[0])

	d := Document{
		originalFilename: filename,
	}

	for _, line := range matrix[1:] {
		p := newPanel(logger, line, headers)
		if p.isValid {
			d.panels = append(d.panels, p)
		}
	}

	return d, nil
}

func newPanel(logger *logging.Logger, line []string, headers map[int]string) *panel {

	p := &panel{
		accession:       line[2],
		internalControl: line[4],
		tests:           make(map[string]bool),
	}

	rx := regexp.MustCompile(`\d{10}`)

	if !rx.Match([]byte(p.accession)) {
		return p
	}

	if p.internalControl != "PRES" {
		return p
	}

	var counterPos int

	for index, testCode := range headers {
		value := line[index]
		switch value {
		case "NEG":
			p.tests[testCode] = false
		case "POS":
			p.tests[testCode] = true
			counterPos++
		default:
			return p
		}
	}

	if counterPos > 1 {
		return p
	}

	p.isValid = true

	return p
}
