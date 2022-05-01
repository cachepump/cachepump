package pump

import (
	"github.com/cachepump/cachepump/cache"

	logger "github.com/AntonYurchenko/log-go"
	"github.com/robfig/cron/v3"
)

// if a rule is not defined for source in yaml file this value of rule will be used by default.
const defaultRule = "0 0 * * * *"

// runNew runs all cron jobs for updating of new sources.
// New source is a source updating of which is not ran but he is define in a current configuration.
func runNew(scheduler *cron.Cron) {
	for name, source := range getConfig().Sources {
		if _, ok := getWorkedID(name); ok {
			continue
		}
		rule := defaultRule
		if source.Rule != "" {
			rule = source.Rule
		}
		id, err := scheduler.AddFunc(rule, source.asFunc(name))
		if err != nil {
			logger.WarnF("Source with key %[1]q has not been added to scheduler, error: %[2]v", name, err)
			continue
		}
		setWorkedID(name, id)
		logger.InfoF("New source with key %q has been added to scheduler ", name)
	}
}

// stopOld stops all cron jobs for updating of old sources.
// Old source is a source updating of which is ran but he is not define in a current configuration.
func stopOld(scheduler *cron.Cron) {
	for _, key := range getWorkindSourceNames() {
		if _, ok := getConfig().Sources[key]; ok {
			continue
		}
		id, _ := getWorkedID(key)
		scheduler.Remove(id)
		delWorkindSource(key)
		cache.Del(key)
		logger.InfoF("Old source with key %q has been remowed from scheduler ", key)
	}
}
