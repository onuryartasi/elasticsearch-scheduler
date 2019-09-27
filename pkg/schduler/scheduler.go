package schduler

import (
	"github.com/robfig/cron/v3"
	"log"
	"os"
)


func Cron() *cron.Cron {
	var c *cron.Cron
	DEBUG := os.Getenv("DEBUG")

	if DEBUG == "true"{
		c = cron.New(
			cron.WithLogger(
				cron.VerbosePrintfLogger(log.New(os.Stdout, "Schedule: ", log.LstdFlags))))
	}else {
		c = cron.New()
	}

	return c
}
