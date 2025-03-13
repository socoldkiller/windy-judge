package renderTask

import (
	"windy-judge/internal/F"
	"windy-judge/internal/command"
	"windy-judge/internal/report"
)

type TestCaseRunner struct {
	p     F.Printer
	count *TestCaseResultCount
}

type Option func(runner *TestCaseRunner)

func (c *TestCaseRunner) TestCaseTask(result TestCaseResult) {
	r := report.NewRender(report.WithPrinter(c.p))
	if err := r.Write(&result); err != nil {
		return
	}

	if r.IsAccept() {
		c.count.passed += 1
	}
	r.Beauty()
	c.count.usedTime += result.Elapsed
}

func (c *TestCaseRunner) TestCaseSetTask(result []command.TestCaseResult) {
	r := DefaultTestResult{
		p:                   c.p,
		TestCaseResultCount: c.count,
	}
	if err := r.Write(result); err != nil {
		return
	}
	r.Beauty()

}

func (c *TestCaseRunner) ErrCode() int {
	return exitCode
}

func WithPrinter(p F.Printer) Option {
	return func(runner *TestCaseRunner) {
		runner.p = p
	}
}

func NewTestCaseRunner(opts ...Option) *TestCaseRunner {

	runner := &TestCaseRunner{
		count: &count,
	}

	for _, opt := range opts {
		opt(runner)
	}
	return runner
}
