package provider

import (
	logger "github.com/AntonYurchenko/log-go"
	"github.com/cachepump/cachepump/cache"
)

// Static define available fields for static source.
type Static struct {
	Value string `yaml:"value"`
}

// Pump generate an job function for updating data from static sources.
func (s *Static) Pump(name string) func() {
	return func() {
		if err := cache.Set(name, []byte(s.Value)); err != nil {
			logger.ErrorF("Key %+[1]q has not been added to chacke, error: %[2]v", name, err)
		}
	}
}
