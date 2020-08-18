package interfaces

import (
	"time"
)

type Evaluate interface {
	// config time
	EvaluateProcessTime(funcName string, start time.Time) string
}
