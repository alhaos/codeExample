package sma

import (
	"fmt"
)

/*
func isValidAccession(accession string) bool {
	rx := regexp.MustCompile(`\d{10}`)
	result := rx.MatchString(accession)
	return result
}
*/

func normalResultValue(smn1 string, smn2 string) (string, error) {

	switch smn2 {
	case "0", "1", "2", ">=3":
	default:
		return "", fmt.Errorf("invalid smn2 reslut")
	}

	switch smn1 {
	case ">=3":
		return "N3", nil
	case "2":
		return "N2", nil
	default:
		return "", fmt.Errorf("invalid smn1 result")
	}

}
