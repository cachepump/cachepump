package provider

import logger "github.com/AntonYurchenko/log-go"

// Provider is a general interface for definition a new provider.
type Provider interface {
	Pump(name string) func()
	IsEmpty() bool
}

// EmptyProvider is empty provider. It is used if all supported providers if not defined.
type EmptyProvider struct{}

// Pump generate an job with empty function.
func (ep EmptyProvider) Pump(name string) func() {
	return func() {
		logger.ErrorF("Source with name %q is not supported", name)
	}
}

// IsEmpty returns true if a structure is empty.
func (ep EmptyProvider) IsEmpty() bool { return false }
