package fx

import (
	"GoCFX/internal/logging"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
)

type Document struct {
	originalFilename string
	isValid          bool
	panels           []*panel
}

type panel struct {
	accession string
	fxf1      string
	fxf2      string
	z56i      string
}

// New is constructor for document struct
func New(logger *logging.Logger, filename string) (Document, error) {

	d := Document{}

	d.originalFilename = filename

	f, err := os.Open(filename)
	if err != nil {
		return d, fmt.Errorf("unable open source file: %w", err)
	}

	defer f.Close()

	r := csv.NewReader(f)

	r.FieldsPerRecord = 7

	for i := 0; ; i++ {
		fields, err := r.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if i > 4 {
			d.panels = append(d.panels, &panel{
				accession: fields[0],
				fxf1:      interpretFXF(fields[1]),
				fxf2:      interpretFXF(fields[2]),
				z56i:      interpretZ56I(fields[6]),
			})
		}
	}

	if len(d.panels) > 0 {
		d.isValid = true
	}

	return d, err
}
