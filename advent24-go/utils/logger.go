package utils

import (
	"testing"
	"time"
)

func LogResult(t *testing.T, result int) {
	t.Logf("Result: \033[34m%d\033[0m", result)
}

// Time is in ns by default
func LogTime(t *testing.T, time time.Duration) {
	t.Logf("Elapsed: \033[32m%.2f\033[0m µs", float64(time)/1000.0)
}
