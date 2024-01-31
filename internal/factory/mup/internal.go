package mup

import (
	"encoding/csv"
	"fmt"
	"os"
)

var cutOffMap = map[string]float64{
	"UU": 34.0,
	"MH": 33.0,
	"MG": 31.0,
	"UP": 34.0,
	"IC": 40.0,
}

func getData(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	r := csv.NewReader(f)

	data, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func parseHeaders(line []string) (map[int]string, error) {

	if len(line) < 9 {
		return nil, fmt.Errorf("invalid headers count")
	}

	const fieldStartIndex = 5

	m := make(map[int]string)

	for i := fieldStartIndex; i <= len(line)-3; i = i + 2 {
		m[i] = line[i]
	}

	return m, nil
}
