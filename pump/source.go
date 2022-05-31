package pump

import (
	"github.com/cachepump/cachepump/provider"

	logger "github.com/AntonYurchenko/log-go"
)

// Source define available fields for description a source.
type Source struct {
	Rule      string           `yaml:"rule"`
	StaticSrc *provider.Static `yaml:"static"`
	HttpSrc   *provider.Http   `yaml:"http"`
	FileSrc   *provider.File   `yaml:"file"`
}

// asFunc is a function which generates a function for scheduler.
// A returned function should be update data from source in in memory cache.
func (src *Source) asFunc(name string) func() {
	switch {
	case src.HttpSrc != nil:
		return src.HttpSrc.Pump(name)
	case src.FileSrc != nil:
		return src.FileSrc.Pump(name)
	case src.StaticSrc != nil:
		return src.StaticSrc.Pump(name)
	default:
		logger.ErrorF("Source with name %q is not supported", name)
		return func() {}
	}
}
