package report

import (
	"fmt"
	"github.com/pmezard/go-difflib/difflib"
	"slices"
	"strings"
	"windy-judge/internal/outputter"
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

func getExceptLines(lines []string) []string {

	var exceptLines []string
	for {

		if len(lines) == 0 {
			break
		}

		if line := lines[0]; len(line) != 0 && line[0] == '-' {
			exceptLines = append(exceptLines, line)
		}
		lines = lines[1:]
	}
	return exceptLines

}

func getActualLines(lines []string) []string {

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

	return actualLines
}

type Accept interface {
	IsAccept() bool
}

type Diff interface {
	Accept
	OutPutter
	Diff() string
}

type Differ struct {
	OutPutter

	expected    string
	actual      string
	lines       []string
	actualLines []string
	exceptLines []string
}

func NewDiffer(expected, actual string, printer OutPutter) *Differ {
	lines := strings.Split(generateDiffs(expected, actual), "\n")

	if len(lines) > 2 {
		lines = lines[2:]
	}

	differ := &Differ{
		OutPutter:   printer,
		expected:    expected,
		actual:      actual,
		lines:       lines,
		actualLines: getActualLines(lines),
		exceptLines: getExceptLines(lines),
	}
	return differ
}

func (d *Differ) IsAccept() bool {

	exceptTokens := strings.Fields(strings.Join(d.exceptLines, " "))
	actualTokens := strings.Fields(strings.Join(d.actualLines, " "))

	return slices.EqualFunc(exceptTokens, actualTokens, judgeToken)
}

func (d *Differ) Diff() string {
	text := generateDiffs(d.expected, d.actual)
	return text
}

func (d *Differ) Beauty() {

	if d.IsAccept() {
		return
	}

	lines := d.lines

	if len(lines) < 2 {
		return
	}
	d.Infoln("[Diff Details]")
	d.Defaultln(lines[0])
	d.Successln("Expected: ")

	for _, line := range d.exceptLines {
		d.Defaultln(line)
	}

	d.Errorln("Actual: ")

	for idx, line := range d.actualLines {
		var tokens []string
		if idx < len(d.exceptLines) {
			tokens = exceptLineTokens(d.exceptLines[idx])
		}

		printLine(tokens, line, d)
		fmt.Println()
	}

	lastIdx := len(d.exceptLines) - len(d.actualLines)
	if lastIdx > 0 {
		lastIdx = len(d.exceptLines) - lastIdx
		for _, line := range d.exceptLines[lastIdx:] {
			d.Errorln(line)
		}

	}

}

func exceptLineTokens(line string) []string {
	return strings.Fields(line)
}

func printLine(tokens []string, line string, p outputter.OutPutter) {
	for {
		if len(line) == 0 {
			break
		}

		switch line[0] {

		case ' ', '\t':
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
