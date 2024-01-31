package factory

import (
	"GoCFX/internal/logging"
)

type (
	Factory interface {
		ParseCfxFile(filename string, testType string) Document
	}

	factory struct {
		logger *logging.Logger
	}
)

func New(logger *logging.Logger) Factory {
	return &factory{
		logger: logger,
	}
}

type Document interface {
	OriginalFilename() string
	IsValid() bool
	TestType() string
	String() string
}
