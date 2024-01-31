package opRpp

import (
	"fmt"
	"regexp"
)

var cutOffMap = map[string]cutOff{
	"Flu A-H1":      {"H1", 39.0},
	"Flu A-H3":      {"H3", 40.0},
	"Flu A-H1pdm09": {"pdm09", 39.0},
	"Flu A":         {"FluA", 40.0},
	"Flu B":         {"FluB", 39.0},
	"RSV A":         {"RSVA", 38.0},
	"RSV B":         {"RSVB", 40.0},
	"IC":            {"IC", 40.0},
	"AdV":           {"AdV", 40.0},
	"HEV":           {"HEV", 40.0},
	"MPV":           {"MPV", 38.0},
	"PIV1":          {"PIV1", 40.0},
	"PIV2":          {"PIV2", 39.0},
	"PIV3":          {"PIV3", 39.0},
	"PIV4":          {"PIV4", 39.0},
	"HBoV":          {"HBoV", 37.0},
	"HRV":           {"HRV", 38.0},
	"229E":          {"229E", 38.0},
	"NL63":          {"NL63", 39.0},
	"OC43":          {"OC43", 37.0},
	"CP":            {"CP", 37.0},
	"MP":            {"MP", 35.0},
	"LP":            {"LP", 33.0},
	"BP":            {"BP", 34.0},
	"BPP":           {"BPP", 33.0},
	"SP":            {"SP", 33.0},
	"HI":            {"HI", 39.0},
}

type cutOff struct {
	name  string
	value float64
}

func parseHeaders(line []string) (map[int]string, error) {

	m := make(map[int]string)

	if len(line) < 9 {
		return m, fmt.Errorf("invalid headers count")
	}

	const fieldStartIndex = 5

	for i := fieldStartIndex; i <= len(line)-3; i = i + 2 {
		m[i] = line[i]
	}

	return m, nil
}

func isDataLine(line []string) bool {
	rx := regexp.MustCompile(`^\d{10}$`)
	return rx.MatchString(line[3])
}
