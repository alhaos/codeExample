package fx

import (
	"fmt"
	"strings"
)

func (d Document) OriginalFilename() string {
	return d.originalFilename
}

func (d Document) IsValid() bool {
	return d.isValid
}

func (d Document) String() string {
	sb := strings.Builder{}
	sb.WriteString("accession,TestName,TestResult")

	for _, p := range d.panels {
		sb.WriteString(fmt.Sprintf("\r\n%s,FXF1,%s", p.accession, p.fxf1))
		sb.WriteString(fmt.Sprintf("\r\n%s,FXF2,%s", p.accession, p.fxf2))
		sb.WriteString(fmt.Sprintf("\r\n%s,Z56I,%s", p.accession, p.z56i))
	}

	return sb.String()
}

func (d Document) TestType() string {
	return "fx"
}

func interpretFXF(value string) string {
	if strings.TrimSpace(value) == "" {
		return "."
	}
	return value
}

func interpretZ56I(value string) string {
	v := strings.TrimSpace(value)
	switch v {
	case "Normal":
		return "N"
	case "Intermediate":
		return "I"
	case "Premutation":
		return "P"
	case "Full mutation":
		return "F"
	default:
		return v
	}
}
