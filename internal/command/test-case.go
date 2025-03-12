package command

import (
	"strconv"
	"windy-judge/internal/parser"
)

func ReadTestCaseSet(caseParser TestCaseParser) (TestCaseSet, error) {
	return caseParser.Parse()
}

type TestCaseCommand struct {
	parser     TestCaseParser
	taskRunner TestCaseCommandRunner
	c          *Cmd
}

func (cmd *TestCaseCommand) runOneTestCase(testCase TestCase) Result {
	runner := cmd.c
	res := runner.Run(testCase.Input)
	return res
}

func (cmd *TestCaseCommand) Run(in string) Result {

	var (
		err         error
		testCaseSet TestCaseSet
	)

	if testCaseSet, err = ReadTestCaseSet(cmd.parser); err != nil {
		//TODO
	}

	var taskIOResult []TestCaseResult
	for idx, testCase := range testCaseSet {

		res := cmd.runOneTestCase(testCase)

		ts := TestCaseResult{
			ID:       strconv.Itoa(idx),
			Excepted: testCase.Output,
			Result:   res,
		}
		cmd.taskRunner.TestCaseTask(ts)

		taskIOResult = append(taskIOResult, ts)
	}
	cmd.taskRunner.TestCaseSetTask(taskIOResult)

	return Result{}
}

type Option func(*TestCaseCommand)

func NewTestCaseCommand(opts ...Option) CmdResultRunner {
	c := &TestCaseCommand{}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func WithCommand(cmd string, args ...string) Option {
	return func(c *TestCaseCommand) {
		c.c = NewCmd(cmd, args...)
	}
}

func WithTestCaseParser(p parser.TestCaseParser) Option {

	return func(c *TestCaseCommand) {
		c.parser = p
	}
}

func WithTaskRunner(task TestCaseCommandRunner) Option {
	return func(c *TestCaseCommand) {
		c.taskRunner = task

	}
}
