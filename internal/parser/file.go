package parser

import (
	"os"
)

type FileTestCaseParser struct {
	file   *os.File
	parser ReaderParser
}

func NewFileTestCaseParser(filepath string) TestCaseParser {
	file, err := os.Open(filepath)
	if err != nil {
		return nil
	}
	return &FileTestCaseParser{file: file, parser: InputOutputParse(parse)}

}

func (p FileTestCaseParser) Parse() (cases TestCaseSet, err error) {
	return p.parser.Parse(p.file)
}
