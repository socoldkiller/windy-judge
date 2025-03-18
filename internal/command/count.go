package command

import (
	"time"
)

var (
	exitCode int
)

type TestCaseResultCount struct {
	total    int
	passed   int
	failed   int
	usedTime time.Duration
}

var count TestCaseResultCount
