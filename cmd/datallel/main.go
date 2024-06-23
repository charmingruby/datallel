package main

import simulationworker "github.com/charmingruby/datallel/internal/worker/simulation_worker"

func main() {
	// simulates data
	data := []simulationworker.SimulationPayload{}
	for i := 0; i < 1000; i++ {
		payload := simulationworker.SimulationPayload{
			UserID: i,
			Status: "Waiting",
		}
		data = append(data, payload)
	}
	concurrency := 10

	simulationworker.ProcessSimulationData(data, concurrency)
}
