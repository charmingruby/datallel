package simulationworker

import (
	"fmt"
)

// simulates a payload
type SimulationPayload struct {
	UserID int
	Status string
}

type SimulationResult struct {
	UserID      int
	Status      int
	ProcessedBy int
	Processed   bool
}

func SimulationWorker(
	workerID int,
	jobs <-chan SimulationPayload,
	result chan<- SimulationResult,
) {
	for job := range jobs {
		fmt.Printf("Processing - UserID `%d`\n", job.UserID)
		job.Status = "Done"
		res := SimulationResult{
			UserID:      job.UserID,
			Status:      job.UserID,
			ProcessedBy: workerID,
			Processed:   true,
		}
		result <- res
	}
}
