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

var exitCode = 0

type CaseRunner struct {
	p           F.Printer
	passedCount int
	usedTime    time.Duration
	hook        SFCaseHook
}

type Option func(runner *CaseRunner)

func (c *CaseRunner) TestCaseTask(result TestCaseResult) {
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

	switch r.IsAccept() {
	case true:
		c.hook.SuccessCaseHook(result)
	case false:
		c.hook.FailCaseHook(result)

	}
	r.Beauty()
}

func (c *CaseRunner) TestCaseSetTask(result []command.TestCaseResult) {
	var (
		jsonData []byte
		err      error
	)

	res := DefaultTestResult{
		SFCaseInfo: c.hook.Info(),
		p:          c.p,
	}

	if jsonData, err = json.Marshal(res); err != nil {
		return
	}
	_, err = res.Write(jsonData)

	res.Beauty()
}

func (c *CaseRunner) ErrCode() int {
	return exitCode
}

type DefaultTestResult struct {
	SFCaseInfo
	p F.Printer
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
	runner := &CaseRunner{
		hook: &DefaultSFCaseHook{},
	}

	for _, opt := range opts {
		opt(runner)
	}
	return runner
}
