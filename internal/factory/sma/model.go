package sma

import (
	"GoCFX/internal/logging"
	"encoding/csv"
	"os"
)

type Document struct {
	logger           *logging.Logger
	originalFilename string
	isValid          bool
	panels           []*panel
}

type panel struct {
	sampleName      string
	smn1            string
	smn2            string
	snps            string
	classification  string
	conversionEvent string
}

func New(logger *logging.Logger, filename string) (*Document, error) {

	const (
		sampleNameIndex      = 0
		smn1index            = 1
		smn2index            = 2
		snpsIndex            = 3
		classificationIndex  = 4
		conversionEventIndex = 7
	)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	r := csv.NewReader(file)

	data, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	d := &Document{
		logger:           logger,
		originalFilename: filename,
		isValid:          true,
	}

	for _, row := range data[5:] {
		p := &panel{
			sampleName:      row[sampleNameIndex],
			smn1:            row[smn1index],
			smn2:            row[smn2index],
			snps:            row[snpsIndex],
			classification:  row[classificationIndex],
			conversionEvent: row[conversionEventIndex],
		}
		d.panels = append(d.panels, p)
	}

	return d, nil
}
