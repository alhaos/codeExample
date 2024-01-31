package delta

import (
	"fmt"
	"strings"
)

func (d Document) OriginalFilename() string {

	return d.originalFilename

}

func (d Document) IsValid() bool {

	return true

}

func (d Document) TestType() string {

	return "delta"

}

func (d Document) String() string {

	const (
		l452rCutoff = 35
		p681rCutoff = 33
		icCutoff    = 40
	)

	sb := strings.Builder{}
	sb.WriteString("Accession,TestName,TestResult")

	for _, p := range d.panels {

		var (
			l452rResult, p681rResult, k417nResult string
		)

		if !p.ic.isSet || p.ic.value > icCutoff {
			d.logger.Debugf("ivalid ic found %s", p.accession)
			continue
		}

		if p.l452r.isSet {
			if p.l452r.value < l452rCutoff {
				l452rResult = "Detected"
			} else {
				l452rResult = "Not Detected"
			}
		} else {
			l452rResult = "Not Detected"
		}

		if p.p681r.isSet {
			if p.p681r.value < p681rCutoff {
				p681rResult = "Detected"
			} else {
				p681rResult = "Not Detected"
			}
		} else {
			p681rResult = "Not Detected"
		}

		if p.k417n.isSet {
			k417nResult = "Detected"
		} else {
			k417nResult = "Not Detected"
		}

		sb.WriteString(fmt.Sprintf("\r\n%s,L452R,%s", p.accession, l452rResult))
		sb.WriteString(fmt.Sprintf("\r\n%s,P681R,%s", p.accession, p681rResult))
		sb.WriteString(fmt.Sprintf("\r\n%s,K417N,%s", p.accession, k417nResult))
	}

	return sb.String()

}
