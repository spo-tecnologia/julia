package jobs

import (
	"time"

	"github.com/go-co-op/gocron"
)

func RunAsync() error {
	scheduler := gocron.NewScheduler(time.Local)
	scheduler.Every(15).Minutes().Do(RunSampleJob)
	scheduler.StartAsync()
	return nil
}
