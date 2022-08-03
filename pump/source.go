package pump

import (
	"reflect"

	"github.com/cachepump/cachepump/provider"
)

// Source define available fields for description a source.
type Source struct {
	Rule string `yaml:"rule"`

	// If you create a new provider you should describe his mapping here.
	StaticSrc provider.Static `yaml:"static"`
	HttpSrc   provider.Http   `yaml:"http"`
	FileSrc   provider.File   `yaml:"file"`
}

// getProviders returns list of all providers.
func (src Source) getProviders() (prvs []provider.Provider) {

	// Supported providers.
	source := reflect.ValueOf(src)
	for i := 0; i < source.NumField(); i++ {
		value := source.Field(i).Interface()
		prv, ok := value.(provider.Provider)
		if ok {
			prvs = append(prvs, prv)
		}
	}

	// Default provider if all is empty.
	return append(prvs, provider.EmptyProvider{})
}

// asFunc is a function which generates a function for scheduler.
// A returned function should be update data from source in in memory cache.
func (src Source) asFunc(name string) (fn func()) {
	for _, p := range src.getProviders() {
		if !p.IsEmpty() {
			fn = p.Pump(name)
			break
		}
	}
	return fn
}
