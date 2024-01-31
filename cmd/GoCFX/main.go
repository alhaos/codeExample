package main

import (
	"GoCFX/internal/config"
	"GoCFX/internal/factory"
	"GoCFX/internal/fsController"
	"GoCFX/internal/logging"
	"flag"
)

func main() {

	// region init config

	filenamePointer := flag.String("config", "config.yml", "GoCFX config file")
	flag.Parse()
	conf := config.New(*filenamePointer)

	// endregion

	// region init logging

	logger := logging.New(conf.Logging)

	logger.Debugf("start")
	defer logger.Debug("finish")

	// endregion

	// region init file system controller

	fsc := fsController.New(conf.FsController, logger)

	// endregion

	// region init cfx factory

	factory := factory.New(logger)

	// endregion

	testFilesMap := fsc.SourceFiles()

	for testType, filenames := range testFilesMap {
		for _, filename := range filenames {

			doc := factory.ParseCfxFile(filename, testType)

			counter, err := fsc.ReadAndUpdateCounter()
			if err != nil {
				logger.Fatalf("unable read counter file")
			}

			err = fsc.ExportDoc(doc, counter)
			if err != nil {
				logger.Errorf("unable export file %s: %s", doc.OriginalFilename(), err.Error())
			}

			err = fsc.ArchiveDoc(doc, counter)
			if err != nil {
				logger.Errorf("unable archive file %s: %s", doc.OriginalFilename(), err.Error())
			}
		}
	}
}
