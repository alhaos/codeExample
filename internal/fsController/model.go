package fsController

import (
	"GoCFX/internal/factory"
	"GoCFX/internal/logging"
)

type (
	FsController interface {
		SourceFiles() map[string][]string
		ExportDoc(doc factory.Document, counter int64) error
		ArchiveDoc(doc factory.Document, counter int64) error
		ReadAndUpdateCounter() (int64, error)
	}

	fsController struct {
		config *Config
		logger *logging.Logger
	}
)

type Config struct {
	SourceDirectories  map[string][]string `yaml:"sourceDirectories"`
	ExportDirectories  []string            `yaml:"exportDirectories"`
	CounterFilename    string              `yaml:"counterFilename"`
	ArchiveDirectories []string            `yaml:"archiveDirectories"`
}

func New(config *Config, logger *logging.Logger) FsController {
	fsc := &fsController{
		config: config,
		logger: logger,
	}
	return fsc
}
