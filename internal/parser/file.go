package parser

import (
	"os"
)

type FileTestCaseParser struct {
	file   *os.File
	parser ReaderParser[TestCase]
}

func NewFileTestCaseParser(filepath string) TestCaseParser[TestCase] {
	file, err := os.Open(filepath)
	if err != nil {
		return nil
	}
	return &FileTestCaseParser{file: file, parser: InputOutputParse(parse)}

}

func (p FileTestCaseParser) Parse() (cases []TestCase, err error) {
	return p.parser.Parse(p.file)
}
