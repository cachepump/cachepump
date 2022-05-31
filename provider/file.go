package provider

import (
	"io/ioutil"

	logger "github.com/AntonYurchenko/log-go"
	"github.com/cachepump/cachepump/cache"
)

// File define available fields for file source.
type File struct {
	Path string `yaml:"path"`
}

// Pump generate an job function for updating data from file sources.
func (f *File) Pump(name string) func() {
	return func() {

		bytes, err := ioutil.ReadFile(f.Path)
		if err != nil {
			logger.ErrorF("I cannot read file, source name: %[1]q error: %[2]v", name, err)
			return
		}

		err = cache.Set(name, bytes)
		if err != nil {
			logger.ErrorF("I cannot save file body to cache, source name: %[1]q, error: %[2]v", name, err)
			return
		}
		logger.InfoF("Data from source %q has been updated.", name)
	}
}
