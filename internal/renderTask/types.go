package renderTask

import (
	"time"
	"windy-judge/internal/command"
)

type TestCaseResult = command.TestCaseResult

type Runner = command.TestCaseTaskRunner

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
