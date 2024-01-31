package uti

import (
	"GoCFX/internal/logging"
	"fmt"
	"log"
)

type Document struct {
	accessions       []accession
	originalFilename string
	sourceType       string
	isValid          bool
}

type accession struct {
	Accession string
	Panel1    panel
	Panel2    panel
	Panel3    panel
	UINT      string
	isValid   bool
}

type panel struct {
	Tests   []test
	IC      test
	isValid bool
}

type test struct {
	Name   string
	Value  int
	Result string
}

func New(logger *logging.Logger, filename string) (Document, error) {

	d := Document{sourceType: "UTI"}

	d.originalFilename = filename

	data, err := ReadCSV(filename)
	if err != nil {
		return d, err
	}

	for i := 1; i < len(data); i = i + 6 {
		a, err2 := newAccession(data[i : i+6])
		if err2 != nil {
			log.Println("creation UTI doc warning ", err2)
			continue
		}
		d.accessions = append(d.accessions, a)
	}

	d.isValid = true

	return d, nil
}

func newAccession(data [][]string) (accession, error) {
	a := accession{}

	if len(data) != 6 {
		return a, fmt.Errorf("accession with invalid length found %d", len(data))
	}

	if !(data[0][3] == data[2][3] && data[0][3] == data[4][3] && isValidAccession(data[0][3])) {
		return a, fmt.Errorf("invalid accession found %s", data[0][3])
	}

	a.Accession = data[0][3]

	p1, err := newPanel(data[0:2])
	if err != nil {
		return a, fmt.Errorf("accession %s crateing panel1 error", a.Accession)
	}
	a.Panel1 = p1

	p2, err := newPanel(data[2:4])
	if err != nil {
		return a, fmt.Errorf("accession %s crateing panel1 error", a.Accession)
	}
	a.Panel2 = p2

	p3, err := newPanel(data[4:6])
	if err != nil {
		return a, fmt.Errorf("accession %s crateing panel1 error", a.Accession)
	}
	a.Panel3 = p3

	a.isValid = true

	a.UINT = interpretUINT(&a)

	return a, nil
}

func newPanel(data [][]string) (panel, error) {

	const (
		icIndexFromEnd         = 4
		lastHeaderIndexFromEnd = 5
		firstHeaderIndex       = 5
		icHeader               = "ic"
	)

	p := panel{}

	if data[0][len(data[0])-icIndexFromEnd] != icHeader {
		return p, fmt.Errorf("ic not found")
	}

	ic, err := newTest("ic", data[1][len(data[0])-icIndexFromEnd+1])
	if err != nil {
		return p, fmt.Errorf("invalid ic found in panel with accession %s: %w", data[0][3], err)
	}

	p.IC = ic

	for i := firstHeaderIndex; i < len(data[0])-lastHeaderIndexFromEnd; i = i + 2 {
		t, err := newTest(data[0][i], data[1][i+1])
		if err != nil {
			log.Println("Invalid test found:", err)
		}

		p.Tests = append(p.Tests, t)
	}

	p.isValid = true

	return p, nil
}

func newTest(name string, value string) (test, error) {

	t := test{}

	if !isValidTestName(name) {
		return t, fmt.Errorf("invalid test found %s", name)
	}

	t.Name = name

	if value == NA {
		t.Result = "ND"
		return t, nil
	}

	intValue, err := parseTestValue(value)
	if err != nil {
		return t, err
	}

	t.Value = intValue

	switch t.Name {
	case "ic", "San", "Pag", "AS":
	default:
		t.Result = interpret(t.Name, intValue)
	}

	return t, nil
}
