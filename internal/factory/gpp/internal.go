package gpp

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"strings"
)

func getData(filename string) [][]string {

	var (
		headerFound bool
		m           [][]string
	)

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	r := csv.NewReader(f)

	for {
		line, err := r.Read()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil && !strings.Contains(err.Error(), "wrong number of fields") {
			break
		}

		if len(line) > 0 && line[0] == "Batch name" {
			headerFound = true
			r.FieldsPerRecord = len(line)
		}

		if len(line) > 0 && headerFound {
			m = append(m, line)
		}
	}
	return m
}

func getHeaders(line []string) map[int]string {

	const testStartIndex = 5

	codeMap := map[string]string{
		"Shigella":               "X014",
		"Entamoeba histolytica":  "X357",
		"Campylobacter":          "X024",
		"Salmonella":             "X010",
		"Cryptosporidium":        "X359",
		"Adenovirus 40/41":       "X520",
		"C. difficile toxin A/B": "X140",
		"Vibrio cholerae":        "X198",
		"STEC stx1/stx2":         "X032",
		"Rotavirus A":            "X385",
		"E. coli O157":           "X026",
		"Giardia":                "X355",
		"ETEC LT/ST":             "X040",
		"Norovirus GI/GII":       "X371",
	}

	m := make(map[int]string)

	for i, s := range line[testStartIndex : len(line)-1] {
		m[i+testStartIndex] = codeMap[s]
	}

	return m
}
