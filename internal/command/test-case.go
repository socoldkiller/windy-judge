package command

import (
	"fmt"
	"strconv"
	"windy-judge/internal"
	"windy-judge/internal/F"
	"windy-judge/internal/runner"
)

type TestCaseCommand struct {
	putter F.OutPutter
	r      runner.Runner[TestCase, Result]

	cmd  string
	args []string
}

func (t *TestCaseCommand) PreRun(input []TestCase) {
}

func (t *TestCaseCommand) PostRun(input []TestCase, output []internal.Result) {
	var testCaseResult = DefaultTestResult{
		putter: t.putter,
	}
	testCaseResult.Write(count)
	testCaseResult.Beauty()

}

func NewTestCaseCommand(opts ...TestCaseOption) runner.Runner[[]TestCase, []internal.Result] {
	testCaseCmd := &TestCaseCommand{}
	for _, opt := range opts {
		opt(testCaseCmd)
	}

	testCaseCmd.r = NewCmd(
		WithCmd(testCaseCmd.cmd, testCaseCmd.args...),
		WithPrinter(testCaseCmd.putter),
	)

	return runner.NewBatchContextualRunner(testCaseCmd.r, testCaseCmd)
}

type TestCaseOption func(c *TestCaseCommand)

func WithTestCaseCmd(cmd string, args ...string) TestCaseOption {
	return func(c *TestCaseCommand) {
		c.cmd = cmd
		c.args = args
	}
}

func WithTestCasePrinter(p F.OutPutter) TestCaseOption {
	return func(c *TestCaseCommand) {
		c.putter = p
	}
}

type DefaultTestResult struct {
	TestCaseResultCount
	putter F.OutPutter
}

func (r *DefaultTestResult) Write(data any) error {
	r.TestCaseResultCount = data.(TestCaseResultCount)
	return nil
}

func (r *DefaultTestResult) Beauty() {
	p := r.putter
	usedTime := fmt.Sprintf("%.2fs", r.usedTime.Seconds())
	switch {
	case r.total == 0:
		p.Defaultln("🔍 No test cases were found or executed. Please ensure your tests are correctly set up. 🛠️")
	case r.failed == 0:
		p.Defaultln("🎉 Congratulations! All "+strconv.Itoa(r.total)+" test cases passed successfully! ✅🎯", "Execution time:", usedTime, "Keep up the great work! 🚀🔥")
	default:
		p.Errorln("❌ "+strconv.Itoa(r.failed)+"/"+strconv.Itoa(r.total)+" test cases failed!", "Execution time:", usedTime, "Please check the errors above.")
	}
}
