package th

import (
	"GoCFX/internal/logging"
	"encoding/csv"
	"fmt"
	"os"
)

type Document struct {
	originalFilename string
	isValid          bool
	panels           []panel
}

type panel struct {
	isValid   bool
	accession string
	r506q     r506q
	g20210a   g20210a
	c677t     c677t
	a1298c    a1298c
}

type r506q string

type g20210a string

type c677t string

type a1298c string

func New(logger *logging.Logger, filename string) (Document, error) {

	d := Document{}

	d.originalFilename = filename

	fileContent, err := os.Open(filename)
	if err != nil {
		return d, err
	}

	defer func(fileContent *os.File) {
		_ = fileContent.Close()
	}(fileContent)

	csvReader := csv.NewReader(fileContent)

	matrix, err := csvReader.ReadAll()
	if err != nil {
		return d, fmt.Errorf("th csv reading error: %w", err)
	}

	if len(matrix) < 2 {
		return d, fmt.Errorf("th csv line count less then 3")
	}

	for _, line := range matrix[2:] {

		p := panel{isValid: true}

		p.accession, err = newAccession(line[3])
		if err != nil {
			p.isValid = false
			continue
		}

		p.r506q, err = newR506q(line[8])
		if err != nil {
			p.isValid = false
			continue
		}

		p.g20210a, err = newQ20210a(line[11])
		if err != nil {
			p.isValid = false
			continue
		}

		p.c677t, err = newC677t(line[7])
		if err != nil {
			p.isValid = false
			continue
		}

		p.a1298c, err = newA1298c(line[10])
		if err != nil {
			p.isValid = false
			continue
		}

		if p.isValid {
			d.panels = append(d.panels, p)
		}
	}

	if len(d.panels) > 0 {
		d.isValid = true
	}

	return d, err

}
