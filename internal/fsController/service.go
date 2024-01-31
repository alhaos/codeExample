package fsController

import (
	"GoCFX/internal/factory"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func (f *fsController) SourceFiles() map[string][]string {

	m := make(map[string][]string)

	for client, dirs := range f.config.SourceDirectories {
		for _, dir := range dirs {

			if err := dirAvailable(dir); err != nil {
				f.logger.Errorf("%s client dir %s has error: %s", client, dir, err)
				continue
			}

			entries, err := os.ReadDir(dir)
			if err != nil {
				f.logger.Errorf("%s client dir %s reading error: %s", client, dir, err)
			}

			for _, entry := range entries {
				if !entry.IsDir() {
					filename := filepath.Join(dir, entry.Name())
					m[client] = append(m[client], filename)
					f.logger.Infof("found file %s is %s", filename, client)
				}
			}
		}
	}

	return m
}

func (f *fsController) ExportDoc(doc factory.Document, counter int64) error {

	if !doc.IsValid() {
		f.logger.Errorf("invalid doc skipperd, %s", doc.OriginalFilename())
	}

	filename := fmt.Sprintf("r_%d-%s-%s.csv", counter, doc.TestType(), time.Now().Format("20060102-150405"))

	for _, directory := range f.config.ExportDirectories {

		p := filepath.Join(directory, filename)

		err := os.WriteFile(p, []byte(doc.String()), 0666)
		if err != nil {
			return err
		}

		f.logger.Infof("file %s exported to %s", doc.OriginalFilename(), p)
	}

	return nil
}

func (f *fsController) ArchiveDoc(doc factory.Document, counter int64) error {

	base := filepath.Base(doc.OriginalFilename())

	filename := fmt.Sprintf("%d-%s", counter, base)

	var wasCopied bool

	for _, directory := range f.config.ArchiveDirectories {

		p := filepath.Join(directory, filename)

		err := copyFile(doc.OriginalFilename(), p)
		if err != nil {
			continue
		}

		wasCopied = true

		f.logger.Infof("file %s archived to %s", doc.OriginalFilename(), p)
	}

	if wasCopied {
		err := os.Remove(doc.OriginalFilename())
		if err != nil {
			return err
		}
		f.logger.Infof("file %s removed", doc.OriginalFilename())
	}

	return nil
}

func (f *fsController) ReadAndUpdateCounter() (int64, error) {

	file, err := os.OpenFile(f.config.CounterFilename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var count int64

	err = binary.Read(file, binary.LittleEndian, &count)
	if err != nil && err != io.EOF {
		return 0, err
	}

	count++

	_, err = file.Seek(0, 0)
	if err != nil {
		return 0, err
	}

	err = binary.Write(file, binary.LittleEndian, count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
