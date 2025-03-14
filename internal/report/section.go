package report

import "strings"

type Section struct {
	*Report
}

func (p Section) printSection(title, content string) {
	p.Infoln(title + ":")
	if content == "" {
		p.Warnln("<empty>")
	} else {

		lines := strings.Split(content, "\n")
		for _, line := range lines {
			p.Defaultln(line)
		}
	}
}

func (p Section) Beauty() {
	p.printSection("Input", p.Input)
	p.printSection("Expected Output", p.Excepted)
	p.printSection("Program Output", p.Output)
}
