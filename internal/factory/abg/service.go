package abg

import (
	"fmt"
	"strings"
)

func (d *Document) OriginalFilename() string {
	return d.originalFilename
}

func (d *Document) IsValid() bool {
	return d.isValid
}

func (d *Document) TestType() string {
	return "abg"
}

func (d *Document) String() string {

	sb := strings.Builder{}
	sb.WriteString("accession,TestName,TestResult")

	for _, p := range d.panels {
		for _, t := range p.tests {

			if t.name == "TEM" {
				continue
			}

			co, exist := cutOffMap[t.name]
			if !exist {
				continue
			}

			if !t.hasTreeValue() {
				sb.WriteString(fmt.Sprintf("\r\n%s,%s,%s", p.accession, t.name, "ND"))
				continue
			}

			if t.hasNullValue() {
				sb.WriteString(fmt.Sprintf("\r\n%s,%s,%s", p.accession, t.name, "ND"))
				continue
			}

			ct := t.ct()
			if ct <= co.ct {
				sb.WriteString(fmt.Sprintf("\r\n%s,%s,%s", p.accession, t.name, "D"))
				continue

			}

			sd := t.sd()
			if sd <= co.sd {
				sb.WriteString(fmt.Sprintf("\r\n%s,%s,%s", p.accession, t.name, "D"))
				continue
			}
		}

		if len(p.tests) != 1 {
			continue
		}

		tem, exist := p.testsMap["TEM"]
		if !exist {
			continue
		}

		if len(tem.values) != 3 {
			continue
		}

		if tem.ct() > 29 {
			continue
		}

		sb.WriteString(fmt.Sprintf("\r\n%s,TEM,%s", p.accession, "Fail"))
	}

	return sb.String()
}
