package delta

import (
	"GoCFX/internal/logging"
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
)

type Document struct {
	logger           *logging.Logger
	originalFilename string
	isValid          bool
	panels           []*panel
}

func New(logger *logging.Logger, filename string) (*Document, error) {

	xlDoc, err := excelize.OpenFile(filename)
	if err != nil {
		return nil, err
	}

	defer xlDoc.Close()

	sheetList := xlDoc.GetSheetList()

	if len(sheetList) == 0 {
		return nil, fmt.Errorf("no sheets found")
	}

	rows, err := xlDoc.GetRows(sheetList[0])

	d := &Document{
		originalFilename: filename,
		logger:           logger,
	}

	for _, row := range rows[2:] {
		p, err := newPanel(row)
		if err != nil {
			logger.Error("invalid panel found: %s", err.Error())
		}
		d.panels = append(d.panels, p)
	}

	return d, nil
}

type panel struct {
	accession string
	l452r     *test
	p681r     *test
	k417n     *test
	ic        *test
}

func newPanel(row []string) (*panel, error) {

	var err error

	// minimum fields count check
	fieldsCount := len(row)
	if fieldsCount < 14 {
		return nil, fmt.Errorf("minimum row field count violeted [%d]", fieldsCount)
	}

	accession := row[3]
	if !isValidAccession(accession) {
		return nil, fmt.Errorf("invalid accession found [ %s ]", accession)
	}

	p := &panel{accession: accession}

	p.l452r, err = newTest(row[6])
	if err != nil {
		return nil, err
	}

	p.p681r, err = newTest(row[8])
	if err != nil {
		return nil, err
	}

	p.k417n, err = newTest(row[10])
	if err != nil {
		return nil, err
	}

	p.ic, err = newTest(row[12])
	if err != nil {
		return nil, err
	}

	return p, nil
}

type test struct {
	value float64
	isSet bool
}

func newTest(value string) (*test, error) {

	t := &test{}

	if value == "N/A" {
		t.isSet = false
		return t, nil
	}

	valueFloat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil, err
	}

	t.value = valueFloat
	t.isSet = true

	return t, nil
}
