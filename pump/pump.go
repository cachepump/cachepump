package pump

import (
	logger "github.com/AntonYurchenko/log-go"
	"github.com/robfig/cron/v3"
)

var scheduler *cron.Cron

// Start scheduler.
func Start(srcPath string) {

	scheduler = cron.New(cron.WithSeconds())
	scheduler.AddFunc("*/5 * * * * *", func() {
		uploadConfig(srcPath)
		runNew(scheduler)
		stopOld(scheduler)
	})
	scheduler.Start()
	logger.Info("Scheduler has been started")
}

// Stop scheduler.
func Stop() {
	if scheduler != nil {
		logger.Info("Stopping a scheduler ...")
		scheduler.Stop()
	}
}
