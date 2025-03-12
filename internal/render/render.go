package render

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"windy-judge/internal/F"
	"windy-judge/internal/command"
	"windy-judge/internal/report"
)

type Runner = command.TestCaseCommandRunner

type CaseRunner struct {
	p           F.Printer
	passedCount int
	usedTime    time.Duration
}

type Option func(runner *CaseRunner)

func (c *CaseRunner) TestCaseTask(result command.TestCaseResult) {
	var (
		jsonData []byte
		err      error
	)
	r := report.NewRender(report.WithPrinter(c.p))
	if jsonData, err = json.Marshal(result); err != nil {
		return
	}

	if _, err = r.Write(jsonData); err != nil {
		return
	}

	c.usedTime += r.Elapsed
	if r.IsAccept() {
		c.passedCount++
	}

	r.Beauty()

}

func (c *CaseRunner) TestCaseSetTask(result []command.TestCaseResult) {
	var (
		jsonData []byte
		err      error
	)

	res := DefaultTestResult{
		Total:    len(result),
		Passed:   c.passedCount,
		Failed:   len(result) - c.passedCount,
		UsedTime: c.usedTime,
		p:        c.p,
	}

	if jsonData, err = json.Marshal(res); err != nil {
		return
	}
	_, err = res.Write(jsonData)

	res.Beauty()
	return

}

type DefaultTestResult struct {
	Total    int
	Passed   int
	Failed   int
	UsedTime time.Duration
	p        F.Printer
}

func (r *DefaultTestResult) Write(p []byte) (n int, err error) {
	if err := json.Unmarshal(p, r); err != nil {
		return 0, err
	}
	r.Failed = r.Total - r.Passed
	return len(p), nil
}

func (r *DefaultTestResult) Beauty() {
	p := r.p
	usedTime := fmt.Sprintf("%.2fs", r.UsedTime.Seconds())
	switch {
	case r.Total == 0:
		p.Defaultln("🔍 No test cases were found or executed. Please ensure your tests are correctly set up. 🛠️")
	case r.Failed == 0:
		p.Defaultln("🎉 Congratulations! All "+strconv.Itoa(r.Total)+" test cases passed successfully! ✅🎯", "Execution time:", usedTime, "Keep up the great work! 🚀🔥")
	default:
		p.Errorln("❌ "+strconv.Itoa(r.Failed)+"/"+strconv.Itoa(r.Total)+" test cases failed!", "Execution time:", usedTime, "Please check the errors above.")
	}
}

func WithPrinter(p F.Printer) Option {
	return func(runner *CaseRunner) {
		runner.p = p
	}
}

func NewCaseRunner(opts ...Option) *CaseRunner {
	runner := &CaseRunner{}

	for _, opt := range opts {
		opt(runner)
	}
	return runner
}
