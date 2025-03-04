package main

import (
	"strings"
	T "test-cli/F"
	"time"
)

type Report struct {
	ID            string
	TestTime      time.Time
	ExecutionTime time.Duration
	Input         string
	Expected      string
	Output        string
	d             Differ
}

func printSection(title string, content string) {
	T.Infoln(title + ":")
	if content == "" {
		T.Warnln("<empty>")
	} else {
		lines := strings.Split(content, "\n")
		for _, line := range lines {
			T.Defaultln(line)
		}
	}
}

func (tr *Report) Print() {
	T.Infoln("# Test Case", tr.ID, "- Result:")
	T.Info("----------------------------------------------\n")
	T.Timeln("[Timestamps]")
	T.TitleTimeF("%s\n", "- Test Time: ", time.Now().Format(time.DateTime))
	T.TitleTimeF("%.2fs\n", "- Execution Time: ", tr.ExecutionTime.Seconds())
	printSection("Input", tr.Input)
	printSection("Expected Output", tr.Expected)
	printSection("Program Output", tr.Output)

	T.Info("[Comparison Result] ")
	if tr.d.IsAccept() {
		T.Successln("✅ Accept (Output Matched)")
	} else {
		T.Errorln("❌ Error (Output Mismatch)")
	}
	tr.d.Print()
}

func NewReport(ID string, ExecutionTime time.Duration, Excepted, Input, Output string) *Report {
	report := &Report{
		ID:            ID,
		TestTime:      time.Now(),
		ExecutionTime: ExecutionTime,
		Input:         Input,
		Expected:      Excepted,
		Output:        Output,
		d: Differ{
			expected: Excepted,
			actual:   Output,
		},
	}

	return report
}
