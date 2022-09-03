package cron

import (
	"OctaneServer/config"
	"OctaneServer/handlers"
	"OctaneServer/storage"
	"time"

	"github.com/go-co-op/gocron"
)

var cron = gocron.NewScheduler(time.Local)

func init() {
	cron.Every(1).Minute().Do(handlers.HouseKeeping)
	if config.CurrentConfig.Storage.HouseKeeping != "" {
		cron.Every(config.CurrentConfig.Storage.HouseKeeping).Do(storage.Storage.HouseKeeping)
	}
	cron.StartAsync()
}
