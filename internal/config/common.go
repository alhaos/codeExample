package config

import (
	"GoCFX/internal/fsController"
	"GoCFX/internal/logging"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"time"
)

type Config struct {
	Logging      *logging.Config      `yaml:"logging"`
	FsController *fsController.Config `yaml:"fsController"`
}

func New(filename string) *Config {

	data, err := os.ReadFile(filename)
	if err != nil {
		crashLog(err.Error())
	}

	c := &Config{}

	err = yaml.Unmarshal(data, c)
	if err != nil {
		crashLog(err.Error())
	}

	return c
}

func crashLog(message string) {
	p, _ := os.Executable()
	d := filepath.Dir(p)
	crashFile := filepath.Join(d, "crash_"+time.Now().Format("20060102-150405"))
	_ = os.WriteFile(crashFile, []byte(message), 0666)
	_, _ = fmt.Fprintln(os.Stderr, message)
	os.Exit(1)
}
