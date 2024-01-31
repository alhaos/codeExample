package th

import (
	"fmt"
	"strings"
)

func (d Document) String() string {
	sb := strings.Builder{}
	sb.WriteString("accession,TestName,TestResult")

	for _, p := range d.panels {
		sb.WriteString(fmt.Sprintf("\r\n%s,FVLT,%s", p.accession, p.r506q))
		sb.WriteString(fmt.Sprintf("\r\n%s,PGTR,%s", p.accession, p.g20210a))
		sb.WriteString(fmt.Sprintf("\r\n%s,MTH1,%s", p.accession, p.c677t))
		sb.WriteString(fmt.Sprintf("\r\n%s,MTH2,%s", p.accession, p.a1298c))
	}

	return sb.String()
}

func (d Document) OriginalFilename() string {
	return d.originalFilename
}

func (d Document) TestType() string {
	return "th"
}

func (d Document) IsValid() bool {
	return d.isValid
}
