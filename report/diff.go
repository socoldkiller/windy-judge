package report

import (
	"fmt"
	"github.com/pmezard/go-difflib/difflib"
	"io"
	"strings"
	"test-cli/F"
)

func generateDiffs(expected, actual string) string {
	diff := difflib.UnifiedDiff{
		A:        difflib.SplitLines(expected),
		B:        difflib.SplitLines(actual),
		FromFile: "Expected",
		ToFile:   "Actual",
	}
	text, _ := difflib.GetUnifiedDiffString(diff)
	return text
}

type Accept interface {
	IsAccept() bool
}

type Diff interface {
	Accept
	ReportPrinter
	Diff() string
	ExceptLines() []string
	ActualLines() []string
	Lines() []string
}

type Differ struct {
	expected    string
	actual      string
	lines       []string
	actualLines []string
	exceptLines []string
	ReportPrinter
}

func NewDiffer(expected, actual io.Reader, printer ReportPrinter) *Differ {
	return &Differ{
		expected:      ReadAll(expected),
		actual:        ReadAll(actual),
		ReportPrinter: printer,
	}
}

func (d *Differ) ExceptLines() []string {
	if len(d.exceptLines) != 0 {
		return d.exceptLines
	}

	var actualLines []string
	lines := d.Lines()
	for {

		if len(lines) == 0 {
			break
		}

		if line := lines[0]; len(line) != 0 && line[0] == '-' {
			actualLines = append(actualLines, line)
		}
		lines = lines[1:]
	}

	d.exceptLines = actualLines
	return d.exceptLines
}

func (d *Differ) ActualLines() []string {
	lines := d.Lines()

	if len(d.actualLines) != 0 {
		return d.actualLines
	}

	var actualLines []string
	for {

		if len(lines) == 0 {
			break
		}

		if line := lines[0]; len(line) != 0 && line[0] == '+' {
			actualLines = append(actualLines, line)
		}
		lines = lines[1:]
	}

	d.actualLines = actualLines
	return d.actualLines
}

func (d *Differ) Lines() []string {
	if len(d.lines) == 0 {
		d.lines = strings.Split(generateDiffs(d.expected, d.actual), "\n")[2:]
	}
	return d.lines
}

func (d *Differ) IsAccept() bool {
	return d.expected == d.actual
}

func (d *Differ) Diff() string {
	text := generateDiffs(d.expected, d.actual)
	return text
}

func (d *Differ) Beauty() {

	text := d.Diff()
	lines := strings.Split(text, "\n")

	if len(lines) < 2 {
		return
	}
	d.Infoln("[Diff Details]")
	d.Defaultln(lines[2])
	d.Successln("Expected: ")

	for _, line := range d.ExceptLines() {
		d.Defaultln(line)
	}

	d.Errorln("Actual: ")

	for idx, line := range d.ActualLines() {
		var tokens []string
		if idx < len(d.ExceptLines()) {
			tokens = exceptLineTokens(d.ExceptLines()[idx])
		}

		printLine(tokens, line, d)
		fmt.Println()
	}

	lastIdx := len(d.ExceptLines()) - len(d.ActualLines())
	if lastIdx > 0 {
		lastIdx = len(d.ExceptLines()) - lastIdx
		for _, line := range d.ExceptLines()[lastIdx:] {
			d.Errorln(line)
		}

	}

}

func exceptLineTokens(line string) []string {
	return strings.Fields(line)
}

func printLine(tokens []string, line string, p F.Printer) {
	for {
		if len(line) == 0 {
			break
		}

		switch line[0] {

		case ' ':
			fmt.Printf("%c", line[0])
			line = line[1:]

		default:
			var newToken string
			if index := strings.Index(line, " "); index != -1 {
				newToken = line[:index]
				line = line[index:]
			} else {
				newToken = line
				line = ""
			}

			if len(tokens) == 0 {
				p.Error(newToken)
				continue
			}

			if judgeToken(tokens[0], newToken) {
				p.Success(newToken)
			} else {
				p.Error(newToken)
			}
			tokens = tokens[1:]

		}

	}
}

func judgeToken(oldToken, newToken string) bool {

	if oldToken == newToken {
		return true
	}

	if oldToken[0] == '-' && newToken[0] == '+' && oldToken[1:] == newToken[1:] {
		return true
	}
	return false

}
