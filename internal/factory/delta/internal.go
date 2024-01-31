package delta

import "regexp"

func isValidAccession(accession string) bool {
	rx := regexp.MustCompile(`\d{10}`)
	r := rx.MatchString(accession)
	return r
}
