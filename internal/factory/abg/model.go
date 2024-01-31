package abg

import (
	"GoCFX/internal/logging"
	"strconv"
)

type Document struct {
	originalFilename string
	isValid          bool
	panels           []*panel
	panelsMap        map[string]*panel
}

type panel struct {
	accession string
	isValid   bool
	tem       value
	tests     []*test
	testsMap  map[string]*test
}

type test struct {
	name      string
	values    []*value
	valuesMap map[string]*value
}

func (t test) hasTreeValue() bool {
	return len(t.values) == 3
}

func (t test) ct() float64 {
	return t.values[0].ct
}

func (t test) hasNullValue() bool {
	for _, v := range t.values {
		if v.isNull {
			return true
		}
	}
	return false
}

func (t test) sd() float64 {
	return t.values[0].sd
}

type value struct {
	ct     float64
	sd     float64
	isNull bool
}

type dataRow struct {
	accession, testName, ct, sd string
}

func New(logger *logging.Logger, filename string) (*Document, error) {

	d := &Document{
		originalFilename: filename,
		isValid:          true,
		panelsMap:        make(map[string]*panel),
	}

	// get data rows from file
	rows, err := dataRowsFromExcelFile(filename)
	if err != nil {
		logger.Error(err)
		return d, err
	}

	for _, row := range rows {
		err = addDataRowToDoc(d, row)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
	}

	return d, nil
}

func addDataRowToDoc(d *Document, row *dataRow) error {

	var (
		p     *panel
		t     *test
		exist bool
	)

	p, exist = d.panelsMap[row.accession]
	if !exist {
		p = newPanel(row.accession)
		d.panels = append(d.panels, p)
		d.panelsMap[row.accession] = p
	}

	t, exist = p.testsMap[row.testName]
	if !exist {
		t = newTest(row.testName)
		p.tests = append(p.tests, t)
		p.testsMap[row.testName] = t
	}

	v, err := newValue(row.sd, row.sd)
	if err != nil {
		return err
	}

	t.values = append(t.values, v)

	return nil
}

func newPanel(accession string) *panel {
	p := &panel{
		accession: accession,
		isValid:   false,
		tem:       value{},
		tests:     []*test{},
		testsMap:  make(map[string]*test),
	}

	return p
}

func newTest(name string) *test {

	t := &test{
		name:      name,
		values:    []*value{},
		valuesMap: make(map[string]*value),
	}

	return t
}

func newValue(ct, sd string) (*value, error) {
	if ct == "" || sd == "" {
		return &value{isNull: true}, nil
	}

	ctFloat, err := strconv.ParseFloat(ct, 64)
	if err != nil {
		return nil, err
	}

	sdFloat, err := strconv.ParseFloat(sd, 64)
	if err != nil {
		return nil, err
	}

	v := &value{
		ct: ctFloat,
		sd: sdFloat,
	}

	return v, nil
}

func newDataRow(accession string, testName string, ct string, sd string) (*dataRow, error) {
	return &dataRow{accession: accession, testName: testName, ct: ct, sd: sd}, nil
}
