package internal

import "time"

type Result struct {
	Input   string
	Output  string
	Error   string
	Elapsed time.Duration
}

type TestCaseResult struct {
	ID       string
	Excepted string
	Result
}
