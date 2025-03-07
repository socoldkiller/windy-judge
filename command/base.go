package command

import (
	"errors"
	"io"
	"os/exec"
	"time"
	"windy-judge/F"
)

type Printer = F.Printer

type CmdRunner interface {
	Run(r io.Reader, w io.Writer) error
}

type Render interface {
	Printer() Printer
	Write(p []byte) (n int, err error)
	Beauty()
}

type Cmd struct {
	cmd  string
	args []string
}

func NewCmd(cmd string, args ...string) CmdRunner {
	return &Cmd{cmd: cmd, args: args}
}

func (c *Cmd) Run(r io.Reader, w io.Writer) error {
	cmd := exec.Command(c.cmd, c.args...)
	cmd.Stdin = r
	cmd.Stdout = w
	return cmd.Run()
}

type Result struct {
	Input  string
	Output string
	Error  string
}

type TestCaseResult struct {
	ID            string
	ExecutionTime time.Duration
	Excepted      string
	Result
}

func NewTestCaseResult(ID string, executionTime time.Duration, excepted, input, output string, e error) *TestCaseResult {
	if e == nil {
		e = errors.New("")
	}
	return &TestCaseResult{ID: ID, ExecutionTime: executionTime, Excepted: excepted, Result: Result{Input: input, Output: output, Error: e.Error()}}
}
