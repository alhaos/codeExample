package fsController

import (
	"fmt"
	"io"
	"os"
)

func dirAvailable(name string) error {

	stat, err := os.Stat(name)
	if err != nil {
		return err
	}

	if !stat.IsDir() {
		return fmt.Errorf("%s is not direcotry", name)
	}

	return nil

}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	return destFile.Sync()
}
