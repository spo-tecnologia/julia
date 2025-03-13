package jobs

import (
	"context"
	"encoding/json"
	"fmt"
)

func RunSampleJob(ctx context.Context, payload string) error {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(payload), &data)
	if err != nil {
		return fmt.Errorf("invalid payload: %s", payload)
	}

	fmt.Println("Running SampleJob with payload:", data)
	return nil
}
