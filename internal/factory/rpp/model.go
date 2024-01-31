package rpp

import (
	"GoCFX/internal/logging"
	"fmt"
)

type Document struct {
	originalFilename string
	isValid          bool
	accessions       []accession
}

type accession struct {
	accessionID string
	panels      []panel
}

type panel struct {
	accessionID string
	tests       []test
}

type test struct {
	name   string
	result string
}

type headerInfo struct {
	headersMap map[int]string
	icIndex    int
}

func New(logger *logging.Logger, filename string) (Document, error) {

	const minLineCount = 3

	d := Document{
		originalFilename: filename,
	}

	matrix, err := readCSV(filename)
	if err != nil {
		return d, err
	}

	if len(matrix) < minLineCount {
		return d, fmt.Errorf("line count less then %d", minLineCount)
	}

	if len(matrix)%2 == 0 {
		return d, fmt.Errorf("invalid rpp file even line count")
	}

	accessionPanelMap := make(map[string][]panel)

	for i, _ := range matrix[1:] {
		if i%2 == 0 {

			p, err := newPanel(logger, matrix[i+1], matrix[i+2])
			if err != nil {
				logger.Errorf("invalid panel detected: %s", err.Error())
			}

			accessionPanelMap[p.accessionID] = append(accessionPanelMap[p.accessionID], p)
		}
	}

	if len(accessionPanelMap) > 0 {
		for acc, panels := range accessionPanelMap {
			a := accession{
				accessionID: acc,
			}
			for _, p := range panels {
				a.panels = append(a.panels, p)
			}
			d.accessions = append(d.accessions, a)
		}
	}

	d.isValid = true

	return d, nil
}
