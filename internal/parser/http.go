package parser

import (
	"io"
	"resty.dev/v3"
)

type HttpTestCaseParser struct {
	r      io.Reader
	parser ReaderParser[TestCase]
}

func NewHttpTestCaseParser(URL string) TestCaseParser[TestCase] {
	response, err := resty.New().R().Get(URL)
	if err != nil {
		return nil
	}
	return &HttpTestCaseParser{r: response.Body, parser: InputOutputParse(parse)}

}

func (p HttpTestCaseParser) Parse() ([]TestCase, error) {
	return p.parser.Parse(p.r)
}
