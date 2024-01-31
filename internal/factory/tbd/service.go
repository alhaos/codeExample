package tbd

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var testDefinitionMap = map[string]testDefinition{
	"BMayo": {37, "BMayo"},
	"BBss":  {34, "BBss"},
	"BD":    {34, "BD"},
	"BMic":  {37, "BMic"},
	"BMiya": {35, "BMiya"},
	"Ricke": {34, "Ricke"},
	"AP":    {35, "AP"},
	"EE":    {36, "EE"},
	"EC":    {35, "EC"},
	"ic":    {40, "ic"},
}

func (d Document) OriginalFilename() string {
	return d.originalFilename
}

func (d Document) TestType() string {
	return "tbd"
}

func (d Document) String() string {

	sb := strings.Builder{}
	sb.WriteString("accession,TestName,TestResult")

	for _, p := range d.panels {
		for _, t := range p.tests {
			sb.WriteString(fmt.Sprintf("\r\n%s,%s,%d", p.accession, t.name, t.value))
		}
	}

	return sb.String()

}

func (d Document) IsValid() bool {
	return d.isValid
}

func checkAccession(accession string) bool {
	rx := regexp.MustCompile(`\d{10,11}`)
	return rx.Match([]byte(accession))
}

func parseIntFomFloatString(s string) (int, error) {
	float, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}

	float = math.Round(float)

	return int(float), nil
}

func readFields(line []string) map[int]string {
	fieldMap := make(map[int]string)
	for i, s := range line {
		for field := range testDefinitionMap {
			if field == s {
				fieldMap[i] = s
			}
		}
	}
	return fieldMap
}
