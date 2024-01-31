package factory

import (
	"GoCFX/internal/factory/abg"
	"GoCFX/internal/factory/av"
	"GoCFX/internal/factory/bvag"
	"GoCFX/internal/factory/candida"
	"GoCFX/internal/factory/covid"
	"GoCFX/internal/factory/delta"
	"GoCFX/internal/factory/fx"
	"GoCFX/internal/factory/gpp"
	"GoCFX/internal/factory/mup"
	"GoCFX/internal/factory/opRpp"
	"GoCFX/internal/factory/rpp"
	"GoCFX/internal/factory/sma"
	"GoCFX/internal/factory/tbd"
	"GoCFX/internal/factory/th"
	"GoCFX/internal/factory/uti"
	"GoCFX/internal/factory/utiw"
)

func (c factory) ParseCfxFile(filename string, testType string) Document {
	switch testType {
	case "fx":
		doc, err := fx.New(c.logger, filename)
		if err != nil {
			c.logger.Errorf("invalid fx file found %s", filename)
		}
		return doc
	case "tbd":
		doc, err := tbd.New(c.logger, filename)
		if err != nil {
			c.logger.Errorf("invalid tbd file found %s", filename)
		}
		return doc
	case "gpp":
		doc, err := gpp.New(c.logger, filename)
		if err != nil {
			c.logger.Errorf("invalid tbd file found %s", filename)
		}
		return doc
	case "mup":
		doc, err := mup.New(c.logger, filename)
		if err != nil {
			c.logger.Errorf("invalid mup file found %s", filename)
		}
		return doc
	case "th":
		doc, err := th.New(c.logger, filename)
		if err != nil {
			c.logger.Errorf("invalid th file found %s", filename)
		}
		return doc
	case "uti":
		doc, err := uti.New(c.logger, filename)
		if err != nil {
			c.logger.Errorf("invalid th file found %s", filename)
		}
		return doc
	case "abg":
		doc, err := abg.New(c.logger, filename)
		if err != nil {
			c.logger.Errorf("invalid abg file found %s", filename)
		}
		return doc
	case "opRpp":
		doc, err := opRpp.New(c.logger, filename)
		if err != nil {
			c.logger.Errorf("invalid opRpp file found %s", filename)
		}
		return doc
	case "rpp":
		doc, err := rpp.New(c.logger, filename)
		if err != nil {
			c.logger.Errorf("invalid rpp file found %s", filename)
		}
		return doc
	case "av":
		doc, err := av.New(c.logger, filename)
		if err != nil {
			c.logger.Errorf("invalid av file found %s", filename)
		}
		return doc
	case "bvag":
		doc, err := bvag.New(c.logger, filename)
		if err != nil {
			c.logger.Errorf("invalid bvag file found %s", filename)
		}
		return doc
	case "candida":
		doc, err := candida.New(c.logger, filename)
		if err != nil {
			c.logger.Errorf("invalid candida file found %s", filename)
		}
		return doc
	case "utiw":
		doc, err := utiw.New(c.logger, filename)
		if err != nil {
			c.logger.Errorf("invalid utiw file found %s", filename)
		}
		return doc
	case "sma":
		doc, err := sma.New(c.logger, filename)
		if err != nil {
			c.logger.Errorf("invalid sma file found %s", filename)
		}
		return doc
	case "delta":
		doc, err := delta.New(c.logger, filename)
		if err != nil {
			c.logger.Errorf("invalid delta file found %s", filename)
		}
		return doc
	case "covid":
		doc, err := covid.New(c.logger, filename)
		if err != nil {
			c.logger.Errorf("invalid covid file found %s", filename)
		}
		return doc
	default:
		return nil
	}
}
