package report

import (
	"errors"
	"fmt"
	"time"
	F2 "windy-judge/internal/F"
	"windy-judge/internal/command"
)

type Options func(*Report)

func WithPrinter(p F2.Printer) Options {
	return func(r *Report) {
		r.ReportPrinter = p
	}
}

type ReportPrinter = F2.Printer
type TestCaseResult = command.TestCaseResult

type Report struct {
	TestTime time.Time
	TestCaseResult
	d Diff
	*Title
	*Judge
	*Section
	ReportPrinter
}

func (r *Report) Beauty() {
	r.Title.Beauty()
	r.Section.Beauty()
	r.Judge.Beauty()
	r.d.Beauty()

	if r.TestCaseResult.Error != "" {
		r.Warn(errors.New(r.TestCaseResult.Error))
	}

}

func (r *Report) IsAccept() bool {
	return r.d.IsAccept()
}

func NewRender(opts ...Options,
) *Report {
	r := &Report{
		ReportPrinter: new(F2.Terminal),
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

func (r *Report) Write(data any) (err error) {
	result := data.(*TestCaseResult)
	r.TestCaseResult = *result
	*r = Report{
		TestTime:       time.Now(),
		d:              NewDiffer(r.Excepted, r.Output, r.ReportPrinter),
		Title:          &Title{r},
		Section:        &Section{r},
		Judge:          &Judge{r},
		ReportPrinter:  r.ReportPrinter,
		TestCaseResult: r.TestCaseResult,
	}
	return nil
}

func (r *Report) Warn(err error) {
	r.Errorln("[Warning]")
	warningInfo := fmt.Sprintf("⚠️ warning!  %s", err.Error())
	r.Warnln(warningInfo)
}

func (r *Report) Printer() F2.Printer {
	return r.ReportPrinter
}
