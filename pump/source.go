package pump

import (
	"github.com/cachepump/cachepump/cache"
	"github.com/cachepump/cachepump/provider"

	logger "github.com/AntonYurchenko/log-go"
)

// Source define available fields for description a source.
type Source struct {
	Tipe       string              `yaml:"type"`
	Value      string              `yaml:"value"`
	Rule       string              `yaml:"rule"`
	HttpSource provider.HttpSource `yaml:"http"`
}

// asFunc is a function which generates a function for scheduler.
// A returned function should be update data from source in in memory cache.
func (src *Source) asFunc(name string) func() {
	switch src.Tipe {
	case "http":
		return provider.HttpPump(&src.HttpSource, name)
	default:
		return func() {
			if err := cache.Set(name, []byte(src.Value)); err != nil {
				logger.ErrorF("Key %+[1]q has not been added to chacke, error: %[2]v", name, err)
			}
		}
	}
}
