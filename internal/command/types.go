package command

import (
	"windy-judge/internal/F"
	"windy-judge/internal/parser"
)

type Printer = F.Printer

type TestCase = parser.TestCase
type TestCaseSet = parser.TestCaseSet
type TestCaseParser = parser.TestCaseParser

type CmdResultRunner interface {
	Run(input string) Result
	ErrCode() int
}

type TestCaseResult struct {
	ID       string
	Excepted string
	Result
}

type TestCaseTaskRunner interface {
	TestCaseTask(result TestCaseResult)
	TestCaseSetTask(result []TestCaseResult)
	ErrCode() int
}
