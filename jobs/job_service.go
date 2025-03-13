package jobs

import (
	"encoding/json"

	"github.com/OdairPianta/julia/models"
	"github.com/OdairPianta/julia/repository"
)

func RegisterJob(jobType string, payload interface{}) error {
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	job := models.Job{
		Type:    jobType,
		Payload: string(payloadJSON),
		Status:  models.JobStatusPending,
	}

	err = repository.CreateJob(&job)
	if err != nil {
		return err
	}

	return nil
}
