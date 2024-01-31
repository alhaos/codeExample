package uti

import (
	"fmt"
	"strings"
)

func (d Document) String() string {

	sb := strings.Builder{}
	sb.WriteString("accession,TestName,TestResult")

	for _, accession := range d.accessions {
		for _, test := range accession.Panel1.Tests {
			sb.WriteString(fmt.Sprintf("\r\n%s,%s,%s", accession.Accession, test.Name, test.Result))
		}
		for _, test := range accession.Panel2.Tests {
			if test.Name == "AS" {
				continue
			}
			sb.WriteString(fmt.Sprintf("\r\n%s,%s,%s", accession.Accession, test.Name, test.Result))
		}
		for _, test := range accession.Panel3.Tests {
			if test.Name == "San" || test.Name == "Pag" {
				continue
			}
			sb.WriteString(fmt.Sprintf("\r\n%s,%s,%s", accession.Accession, test.Name, test.Result))
		}
		sb.WriteString(fmt.Sprintf("\r\n%s,%s", accession.Accession, accession.UINT))
	}

	return sb.String()
}

func (d Document) OriginalFilename() string {
	return d.originalFilename
}

func (d Document) TestType() string {
	return "UTI"
}

func (d Document) IsValid() bool {
	return d.isValid
}
