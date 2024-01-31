package th

import (
	"fmt"
	"regexp"
)

func newAccession(s string) (string, error) {

	rx := regexp.MustCompile(`\d{10,11}`)

	if !rx.Match([]byte(s)) {
		return "", fmt.Errorf("invalid accession found: %s", s)
	}

	return s, nil
}

func newR506q(value string) (r506q, error) {
	switch value {
	case "Wt":
		return `\1035`, nil
	case "Het":
		return `\1036`, nil
	case "Homo":
		return `\1037`, nil
	default:
		return "error", fmt.Errorf("invalid R506Q test value found: %s", value)
	}
}

func newQ20210a(value string) (g20210a, error) {
	switch value {
	case "Wt":
		return `\1032`, nil
	case "Het":
		return `\1033`, nil
	case "Homo":
		return `\1034`, nil
	default:
		return "error", fmt.Errorf("invalid g20210a test value found: %s", value)
	}
}

func newC677t(value string) (c677t, error) {
	switch value {
	case "Wt":
		return `Wt`, nil
	case "Het":
		return `Het`, nil
	case "Homo":
		return `Homo`, nil
	default:
		return "error", fmt.Errorf("invalid c677t test value found: %s", value)
	}
}

func newA1298c(value string) (a1298c, error) {
	switch value {
	case "Wt":
		return `Wt`, nil
	case "Het":
		return `Het`, nil
	case "Homo":
		return `Homo`, nil
	default:
		return "error", fmt.Errorf("invalid a1298c test value found: %s", value)
	}
}
