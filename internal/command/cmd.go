package command

import (
	"bytes"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"windy-judge/internal"
	"windy-judge/internal/F"
	"windy-judge/internal/report"
	"windy-judge/internal/runner"
)

func MaybeErrorMessage(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}

func registerCount(count *TestCaseResultCount, report *report.Report) {
	switch report.IsAccept() {
	case false:
		count.failed += 1
	case true:
		count.passed += 1
	}
}

var idx int

type TestCase = internal.TestCase
type Result = internal.Result

type Cmd struct {
	cmd  string
	args []string

	count *TestCaseResultCount
	p     F.OutPutter
}

func (c *Cmd) PreRun(input TestCase) {
	idx = idx + 1
}

func (c *Cmd) PostRun(input TestCase, output Result) {
	count := c.count
	ts := internal.TestCaseResult{
		ID:       strconv.Itoa(idx),
		Excepted: input.Output,
		Result:   output,
	}

	re := report.NewRender(report.WithPrinter(c.p))
	re.Write(&ts)
	count.total += 1
	registerCount(count, re)
	count.usedTime += output.Elapsed
	re.Beauty()
}

func NewCmd(opts ...CmdOption) runner.Runner[TestCase, Result] {
	c := &Cmd{
		count: &count,
	}

	for _, opt := range opts {
		opt(c)
	}
	return runner.NewContextualRunner(c, c)
}

func (c *Cmd) Run(input TestCase) Result {
	var buf bytes.Buffer
	cmd := exec.Command(c.cmd, c.args...)
	cmd.Stdin = strings.NewReader(input.Input)
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	start := time.Now()
	e := cmd.Run()
	end := time.Now()
	errMessage := MaybeErrorMessage(e)
	return Result{
		Input:   input.Input,
		Output:  strings.TrimSpace(buf.String()),
		Error:   errMessage,
		Elapsed: end.Sub(start),
	}
}

type CmdOption func(*Cmd)

func WithCmd(cmd string, args ...string) CmdOption {
	return func(c *Cmd) {
		c.cmd = cmd
		c.args = args
	}
}

func WithPrinter(putter F.OutPutter) CmdOption {
	return func(c *Cmd) {
		c.p = putter
	}
}
