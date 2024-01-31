package sma

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
	return "sma"
}

func (d *Document) String() string {

	sb := strings.Builder{}
	sb.WriteString("accession,TestName,TestResult")

	for _, p := range d.panels {

		accession := p.sampleName

		smn1 := p.smn1

		smn2 := p.smn2

		smn0 := p.snps
		if smn0 == "" {
			smn0 = "ND"
		}

		classification := p.classification

		/*
			if !isValidAccession(accession) {
				d.logger.Debugf("accession validation failed [%s]", accession)
				continue
			}
		*/

		var smni string

		switch classification {
		case "Homozygote affected":
			smni = "HP"
		case "Carrier":
			smni = "C"
		case "Risk Factor":
			smni = "AR"
		case "No Call":
			smni = `\1190`
		case "Normal":
			r, err := normalResultValue(smn1, smn2)
			if err != nil {
				d.logger.Error(err)
				continue
			}
			smni = r
		default:
			d.logger.Errorf("invalid classification value [ %s ]", classification)
			continue
		}

		var conversionEvent string
		switch p.conversionEvent {
		case "":
			conversionEvent = "No conversion"
		case "Type I":
			conversionEvent = "Type I"
		case "Type II":
			conversionEvent = "Type II"
		default:
			conversionEvent = p.conversionEvent
		}

		sb.WriteString(fmt.Sprintf("\r\n%s,SMN1,%s", accession, smn1))
		sb.WriteString(fmt.Sprintf("\r\n%s,SMN2,%s", accession, smn2))
		sb.WriteString(fmt.Sprintf("\r\n%s,SMN0,%s", accession, smn0))
		sb.WriteString(fmt.Sprintf("\r\n%s,SMNI,%s", accession, smni))
		sb.WriteString(fmt.Sprintf("\r\n%s,SMCE,%s", accession, conversionEvent))

	}

	return sb.String()

}
