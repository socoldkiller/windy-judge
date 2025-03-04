package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

func ReadTestCaseSet(caseParser TestCaseParser) (TestCaseSet, error) {
	return caseParser.Parse()
}

type TestCaseCommand struct {
	parser     TestCaseParser
	mainRender Render
	CmdRunner
	bottomRender Render
}

type Option func(*TestCaseCommand)

func NewTestCaseCommand(opts ...Option) CmdRunner {
	c := &TestCaseCommand{}

	for _, opt := range opts {
		opt(c)
	}

	if c.bottomRender == nil {
		c.bottomRender = &DefaultTestResult{
			p: c.mainRender.Printer(),
		}
	}
	return c
}

func WithCommand(runner CmdRunner) Option {
	return func(cmd *TestCaseCommand) {
		cmd.CmdRunner = runner
	}

}

func WithTestCase(parser TestCaseParser) Option {
	return func(cmd *TestCaseCommand) {
		cmd.parser = parser
	}
}

func WithRender(render Render) Option {
	return func(cmd *TestCaseCommand) {
		cmd.mainRender = render
	}
}

func WithBottomRender(bottomRender Render) Option {
	return func(cmd *TestCaseCommand) {
		cmd.bottomRender = bottomRender
	}
}

func (cmd *TestCaseCommand) runOneTestCase(r Render, idx int, testCase *TestCase) (Render, bool, time.Duration, error) {
	var (
		buf      bytes.Buffer
		err      error
		jsonData []byte
	)
	input := testCase.Input
	output := testCase.Output
	start := time.Now()
	warnErr := cmd.CmdRunner.Run(strings.NewReader(input), &buf)
	end := time.Now()
	elapsed := end.Sub(start)

	res := NewTestCaseResult(
		strconv.Itoa(idx),
		elapsed,
		output,
		input,
		buf.String(),
		warnErr,
	)

	if jsonData, err = json.Marshal(res); err != nil {
		return nil, false, elapsed, err
	}

	if _, err = r.Write(jsonData); err != nil {
		return nil, false, elapsed, err
	}

	// Check if the test case passed using the IsAccept method if available
	isAccept := true
	if ac, ok := r.(interface{ IsAccept() bool }); ok {
		isAccept = ac.IsAccept()
	}
	return r, isAccept, elapsed, nil
}

func (cmd *TestCaseCommand) writeTestCaseSet(r Render, total, passedCount int, usedTime time.Duration) error {

	var (
		jsonData []byte
		err      error
	)

	res := DefaultTestResult{
		Total:    total,
		Passed:   passedCount,
		Failed:   passedCount - passedCount,
		UsedTime: usedTime,
	}

	if jsonData, err = json.Marshal(res); err != nil {
		return err
	}

	_, err = r.Write(jsonData)

	return err
}
func (cmd *TestCaseCommand) Run(r io.Reader, w io.Writer) error {

	var (
		err         error
		testCaseSet TestCaseSet
	)

	if testCaseSet, err = ReadTestCaseSet(cmd.parser); err != nil {
		return nil
	}

	render := cmd.mainRender
	passedCount := 0
	totalCount := len(testCaseSet)
	var usedTime time.Duration
	for idx, testCase := range testCaseSet {
		var (
			testPassed bool
			elapsed    time.Duration
		)
		if render, testPassed, elapsed, err = cmd.runOneTestCase(render, idx, &testCase); err != nil {
			continue
		}
		usedTime += elapsed
		if testPassed {
			passedCount++
		}
		render.Beauty()
	}

	// Display test results

	if err = cmd.writeTestCaseSet(cmd.bottomRender, totalCount, passedCount, usedTime); err != nil {
		return err
	}
	cmd.bottomRender.Beauty()
	return nil

}

type DefaultTestResult struct {
	Total    int
	Passed   int
	Failed   int
	UsedTime time.Duration
	p        Printer
}

func (r *DefaultTestResult) Printer() Printer {
	return r.p
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
	case r.Failed == 0:
		p.Defaultln("🎉 Congratulations! All "+strconv.Itoa(r.Total)+" test cases passed successfully! ✅🎯", "used time:", usedTime, "Keep up the great work! 🚀🔥")
	default:
		p.Errorln("❌ "+strconv.Itoa(r.Failed)+" of "+strconv.Itoa(r.Total)+" test cases failed!", "used time:", usedTime, "Please check the errors above.")
	}
}
