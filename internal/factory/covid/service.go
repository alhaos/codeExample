package covid

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
		icCutoff = 40
	)

	sb := strings.Builder{}
	sb.WriteString("Accession,TestName,TestResult")

	for _, p := range d.panels {

		var (
			autoInterpretResult string
		)

		d.logger.Debugf("accession:  %v", p.accession)
		d.logger.Debugf("isSet    :  %v", p.ic.isSet)
		d.logger.Debugf("value    :  %v", p.ic.value)

		if !p.ic.isSet || p.ic.value > icCutoff {
			d.logger.Debugf("ivalid ic found %s", p.accession)
			continue
		}

		switch p.autoInterpretation {
		case "2019-nCoV Detected":
			autoInterpretResult = "D"
		case "Negative":
			autoInterpretResult = "ND"
		case "Invalid":
			autoInterpretResult = "ON"
		case "Presumptive positive":
			autoInterpretResult = "ON"
		default:
			d.logger.Errorf("invalid auto interpretation filed value %s", p.accession)
			continue
		}

		sb.WriteString(fmt.Sprintf("\r\n%s,C19,%s", p.accession, autoInterpretResult))
	}

	return sb.String()

}
