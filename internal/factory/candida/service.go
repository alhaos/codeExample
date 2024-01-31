package candida

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
	return "candida"
}

func (d Document) String() string {
	sb := strings.Builder{}
	sb.WriteString("accession,TestName,TestResult")

	for _, p := range d.panels {
		for _, t := range p.tests {
			sb.WriteString(fmt.Sprintf("\r\n%s,%s,%s", p.accession, t.name, t.result))
		}
	}

	return sb.String()
}
