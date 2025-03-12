package render

import (
	"time"
	"windy-judge/internal/command"
)

type TestCaseResult = command.TestCaseResult

type SFCaseInfo struct {
	Total    int
	Passed   int
	Failed   int
	UsedTime time.Duration

	SuccessResult []TestCaseResult
	FailedResult  []TestCaseResult
}

type SFCaseHook interface {
	SuccessCaseHook(result TestCaseResult)
	FailCaseHook(result TestCaseResult)
	Info() SFCaseInfo
}
type Runner = command.TestCaseCommandRunner
