package interfaces

import (
	"time"
)

// Evaluate is a interface
type Evaluate interface {
	// config time
	EvaluateProcessTime(funcName string, start time.Time) string
}
