package simulationworker

import (
	"fmt"
	"strconv"
	"sync"
)

func ProcessSimulationData(data []SimulationPayload, concurrency int) []SimulationResult {
	dataSize := len(data)
	dataCh := make(chan SimulationPayload, dataSize)
	resultCh := make(chan SimulationResult, dataSize)
	var dataWg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i < concurrency; i++ {
		dataWg.Add(1)
		go func() {
			defer dataWg.Done()
			SimulationWorker(i, dataCh, resultCh)
		}()
	}

	// enqueue jobs
	for i := 0; i < dataSize-1; i++ {
		dataCh <- data[i]
	}
	close(dataCh)

	go func() {
		dataWg.Wait()
		close(resultCh)
	}()

	resultsWithErr := []SimulationResult{}

	for result := range resultCh {
		mutex.Lock()
		if !result.Processed {
			resultsWithErr = append(resultsWithErr, result)
		}
		mutex.Unlock()

		isProcessed := strconv.FormatBool(result.Processed)
		fmt.Printf("UserID `%d`, Processed? %s, by: Worker #%d\n", result.UserID, isProcessed, result.ProcessedBy)
	}

	return resultsWithErr
}
