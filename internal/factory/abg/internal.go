package abg

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"regexp"
	"strings"
)

var cutOffMap = map[string]testData{
	"Ba07319991_s1": {"Z957", 28, 1, "aac(6')-Ib-cr"},
	"Ba04646145_s1": {"Z959", 28, 1, "qnrS"},
	"Ba04646160_s1": {"Z961", 27, 1, "qnrA"},
	"Ba07319988_s1": {"Z967", 25, 1, "sul1"},
	"Ba07320003_s1": {"Z969", 27, 1, "sul2"},
	"Ba04646152_s1": {"Z971", 30, 1, "BlaKPC"},
	"Ba04646149_s1": {"Z977", 29, 1, "CTXM-G1"},
	"Ba04646142_s1": {"Z979", 29, 1, "CTXM-G2-bla"},
	"Ba04646134_s1": {"Z981", 25, 1, "BlaOKP-C"},
	"Ba04931076_s1": {"Z991", 25, 1, "BlaNDM-1"},
	"Ba04646155_s1": {"Z993", 27, 1, "BlaVIM"},
	"Ba04646131_s1": {"Z995", 27, 1, "IMP-1-CarbB"},
	"Ba04646158_s1": {"C018", 26, 1, "BlaMP"},
	"Ba04646135_s1": {"N111", 26, 1, "BlaCMY"},
	"Ba04646120_s1": {"T154", 28, 1, "DHA beta-lactamase"},
	"Ba04646126_s1": {"Z112", 27, 1, "FOX-AmpC"},
	"Ba04646133_s1": {"Z114", 28, 1, "blaOXA1"},
	"Ba04646139_s1": {"Z116", 29, 1, "OXA-23"},
	"Ba04646138_s1": {"101Y", 28, 1, "blaOXA"},
	"Ba04646137_s1": {"102Y", 26, 1, "ErmA"},
	"Ba04230913_s1": {"103Y", 27, 1, "ermB"},
	"Ba07319994_s1": {"104Y", 30, 1, "ermC"},
	"Ba04230908_s1": {"105Y", 26, 1, "Methicillin 1"},
	"TetB_APYMKP3":  {"107Y", 25, 1, "TetB"},
	"Ba07921939_s1": {"107Y", 25, 1, "TetB"},
	"Ba04230915_s1": {"108Y", 25, 1, "tetM"},
	"Ba04646147_s1": {"109Y", 28, 1, "vanA"},
	"Ba04646150_s1": {"110Y", 28, 1, "vanB"},
	"TEM":           {"TEM", 29, 1, "TEM"},
}

type testData struct {
	code string
	ct   float64
	sd   float64
	name string
}

// parseTestName extract test name from data
func parseTestName(testName string) string {

	const (
		temNameIndex = 0
	)

	splits := strings.Split(testName, "_")
	if splits[temNameIndex] == "TEM" {
		return "TEM"
	}

	if len(splits) < 3 {
		return testName
	}

	name := strings.Join(splits[len(splits)-2:], "_")

	return name
}

// isDataRow check is data row or not
func isDataRow(row []string) bool {

	const (
		wellNumberFieldIndex = 0
		accessionFieldIndex  = 3
		testNameFieldIndex   = 4
		minFieldCount        = 11
	)

	// checking for a minimum number of fields
	fieldCount := len(row)
	if fieldCount < minFieldCount {
		return false
	}

	// checking if filed index 0 contains an integer
	rowNumberRx := regexp.MustCompile(`^\d+$`)
	wellNumber := row[wellNumberFieldIndex]
	if !rowNumberRx.Match([]byte(wellNumber)) {
		return false
	}

	// validate accession
	accessionRx := regexp.MustCompile(`^\d{10}$`)
	accession := row[accessionFieldIndex]
	if !accessionRx.MatchString(accession) {
		return false
	}

	if row[testNameFieldIndex] == "" {
		return false
	}

	var found bool
	for name := range cutOffMap {
		if strings.Index(row[testNameFieldIndex], string(name)) != -1 {
			found = true
			break
		}
	}

	return found
}

// dataRowsFromExcelFile extract data rows from excel file
func dataRowsFromExcelFile(filename string) ([]*dataRow, error) {

	xlDoc, err := excelize.OpenFile(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to open abg excel file: %w", err)
	}

	rows, err := xlDoc.GetRows("Results")
	if err != nil {
		return nil, fmt.Errorf("unable to open sheet Results in abg excel document: %w", err)
	}

	var dataRows []*dataRow

	const (
		accessionIndex = 3
		testNameIndex  = 4
		ctIndex        = 9
		sdIndex        = 10
	)

	for _, row := range rows {
		if isDataRow(row) {

			tn := parseTestName(row[testNameIndex])

			dr, err := newDataRow(
				row[accessionIndex],
				tn,
				row[ctIndex],
				row[sdIndex],
			)
			if err != nil {
				return nil, err
			}

			dataRows = append(dataRows, dr)
		}
	}

	return dataRows, nil
}
