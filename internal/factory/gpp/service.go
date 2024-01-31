package gpp

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

func (d Document) TestType() string {
	return "gpp"
}

func (d Document) String() string {

	sb := strings.Builder{}
	sb.WriteString("accession,TestName,TestResult")

	for _, p := range d.panels {
		for name, value := range p.tests {
			if value {
				sb.WriteString(fmt.Sprintf("\r\n%s,%s,POSITIVE", p.accession, name))
				continue
			}
			sb.WriteString(fmt.Sprintf("\r\n%s,%s,NEGATIVE", p.accession, name))
		}
	}

	return sb.String()
}
