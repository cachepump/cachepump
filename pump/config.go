package pump

import (
	"os"
	"sync"

	logger "github.com/AntonYurchenko/log-go"
	"gopkg.in/yaml.v2"
)

var config Config
var cfgMtx sync.RWMutex

// getConfig returns a current configuration.
func getConfig() (cfg Config) {
	cfgMtx.RLock()
	defer cfgMtx.RUnlock()
	cfg = config
	return cfg
}

// setConfig sets a new configuration.
func setConfig(cfg Config) {
	cfgMtx.Lock()
	defer cfgMtx.Unlock()
	config = cfg
}

// Config is a structure for serialisation of a yaml file with definition of configurations all data sources.
type Config struct {
	Version string            `yaml:"version"`
	Sources map[string]Source `yaml:"sources"`
}

// uploadConfig reads a new configuration from yaml file and updates current configuration if it is possible.
func uploadConfig(file string) (err error) {
	fl, err := os.Open(file)
	if err != nil {
		return err
	}

	var newConfig Config
	err = yaml.NewDecoder(fl).Decode(&newConfig)
	switch {
	case err != nil:
		logger.ErrorF("Configuration has not been updated, error: %v", err)
	case newConfig.Version != "1.0":
		logger.ErrorF("Version %s is not supported", newConfig.Version)
	default:
		setConfig(newConfig)
		logger.DebugF("Configuretion has been updated, source: %+v", getConfig())
	}

	return nil
}
