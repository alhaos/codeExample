package tbd

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
	accession string
	icValid   bool
	tests     []test
}

type test struct {
	name  string
	value int
}

type testDefinition struct {
	cutOff int
	alias  string
}

// New is constructor for fx.Doc struct
func New(logger *logging.Logger, filename string) (Document, error) {

	d := Document{
		originalFilename: filename,
	}

	f, err := os.Open(filename)
	if err != nil {
		return d, fmt.Errorf("unable open source file: %w", err)
	}

	defer f.Close()

	r := csv.NewReader(f)

	matrix, err := r.ReadAll()
	if err != nil {
		return d, fmt.Errorf("unable to read tbd file %s, %w", filename, err)
	}

	if len(matrix) < 3 {
		return d, fmt.Errorf("invalid fields in tbd file %s", filename)
	}

	fieldsMap := readFields(matrix[1])

	for _, line := range matrix[2:] {

		accession := line[3]

		if !checkAccession(accession) {
			continue
		}

		p := panel{
			accession: accession,
		}

		for index, testName := range fieldsMap {
			if line[index] == "+" {
				value, err := parseIntFomFloatString(line[index+1])
				if err != nil {
					continue
				}
				if testName == "ic" {
					if value <= testDefinitionMap["ic"].cutOff {
						p.icValid = true
					}
				} else {
					if value <= testDefinitionMap[testName].cutOff {
						p.tests = append(p.tests, test{
							name:  testDefinitionMap[testName].alias,
							value: value,
						})
					}
				}
			}
		}

		if len(p.tests) > 0 && p.icValid {
			d.panels = append(d.panels, p)
		}
	}

	if len(d.panels) > 0 {
		d.isValid = true
	}
	return d, err
}
