package mup

import (
	"fmt"
	"strconv"
	"strings"
)

func (d Document) OriginalFilename() string {
	return d.originalFilename
}

func (d Document) IsValid() bool {
	return d.isValid
}

func (d Document) TestType() string {
	return "mup"
}

func (d Document) String() string {
	sb := strings.Builder{}
	sb.WriteString("accession,TestName,TestResult")

	for _, p := range d.panels {
		for name, value := range p.tests {
			if value.pos {
				sb.WriteString(fmt.Sprintf(
					"\r\n%s,%s,%s",
					p.accession,
					name,
					strconv.FormatFloat(value.value, 'f', 2, 64),
				))
				continue
			}
			sb.WriteString(fmt.Sprintf("\r\n%s,%s,ND", p.accession, name))
		}
	}

	return sb.String()
}
