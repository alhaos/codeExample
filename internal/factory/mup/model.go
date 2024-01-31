package mup

import (
	"GoCFX/internal/logging"
	"fmt"
	"strconv"
)

type Document struct {
	originalFilename string
	isValid          bool
	panels           []*panel
}

type panel struct {
	accession string
	isValid   bool
	ic        *test
	tests     map[string]*test
}

type test struct {
	value float64
	pos   bool
}

func New(logger *logging.Logger, filename string) (Document, error) {

	d := Document{
		originalFilename: filename,
	}

	data, err := getData(filename)
	if err != nil {
		return d, err
	}

	headers, err := parseHeaders(data[1])
	if err != nil {
		return d, err
	}

	for _, line := range data[2:] {
		p, err2 := newPanel(headers, line)
		if err2 != nil {
			logger.Errorf("unable to create new panel for file %s", filename)
			continue
		}

		if p.isValid {
			d.panels = append(d.panels, p)
		}
	}
	return d, nil
}

func newPanel(headers map[int]string, line []string) (*panel, error) {

	p := &panel{
		accession: line[4],
		tests:     make(map[string]*test),
	}

	for i, name := range headers {

		value := line[i+1]
		t, err := newTest(name, value)
		if err != nil {
			return nil, err
		}

		if name == "ic" {
			p.ic = t
			continue
		}

		p.tests[name] = t
	}

	if p.ic.pos {
		p.isValid = true
	}

	return p, nil
}

func newTest(name string, value string) (*test, error) {

	cutoff, ok := cutOffMap[name]
	if !ok {
		return nil, fmt.Errorf("no cutOff found for test %s", name)
	}

	if value == "N/A" {
		return &test{value: 0, pos: false}, nil
	}

	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil, err
	}

	if floatValue > cutoff {
		return &test{value: 0, pos: false}, nil
	}

	return &test{value: floatValue, pos: true}, nil
}
