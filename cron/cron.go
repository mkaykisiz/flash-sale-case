package cron

import (
	"flash_sale/pkg/logging"
	"github.com/robfig/cron/v3"
)

// SetupCronJobs initializes and starts cron jobs
func SetupCronJobs() {
	c := cron.New()

	_, err := c.AddFunc("* * * * *", SyncFlashSale)

	if err != nil {
		logging.Error(err)
	}

	c.Start()
}
