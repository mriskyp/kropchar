package infrastructure

import (
	"fmt"
	"time"
)

// EvaluateProcessTime is a func to check evaluate processing time
func EvaluateProcessTime(funcName string, start time.Time) string {
	// Please use defer in each function caller
	totalProcessingTime := fmt.Sprintf("Total processiong time [%s] : %v \n", funcName, time.Since(start))
	return totalProcessingTime
}
