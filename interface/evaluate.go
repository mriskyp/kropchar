package interface

import "time"

type IEvaluate interface {
	// config time
	EvaluateProcessTime(funcName string, start time.Time) string
}