package jobs

import (
	"context"
	"fmt"
	"time"

	"github.com/OdairPianta/julia/models"
	"github.com/OdairPianta/julia/repository"
)

func StartClient() {
	for {
		jobs, err := repository.FindPendingJobs()
		if err != nil {
			fmt.Println("Error fetching pending jobs:", err)
			time.Sleep(5 * time.Second)
			continue
		}

		for _, job := range *jobs {
			err = ProcessJob(job)
			if err != nil {
				fmt.Println("Error processing job:", err)
				err = repository.CreateFailedJob(&models.FailedJob{
					JobID:     job.ID,
					Exception: err.Error(),
					FailedAt:  time.Now(),
				})
				if err != nil {
					fmt.Println("Error saving failed job:", err)
				}
			} else {
				err = repository.DeleteJob(&job)
				if err != nil {
					fmt.Println("Error deleting job:", err)
				}
			}
		}

		time.Sleep(5 * time.Second)
	}
}

func ProcessJob(job models.Job) error {
	switch job.Type {
	case "SampleJob":
		ctx := context.Background()
		err := RunSampleJob(ctx, job.Payload)
		if err != nil {
			return err
		}
		return nil

	default:
		return fmt.Errorf("unsupported job type: %s", job.Type)
	}
}
