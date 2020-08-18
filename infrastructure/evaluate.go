package infrastructure

import (
	"fmt"
	"time"
)

// Please use defer in each function caller
func EvaluateProcessTime(funcName string, start time.Time) string {
	totalProcessingTime := fmt.Sprintf("Total processiong time [%s] : %v \n", funcName, time.Since(start))
	return totalProcessingTime
}
