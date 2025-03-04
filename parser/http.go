package parser

import (
	"io"
	"resty.dev/v3"
)

type HttpTestCaseParser struct {
	r      io.Reader
	parser ReaderParser
}

func NewHttpTestCaseParser(URL string) TestCaseParser {
	response, err := resty.New().R().Get(URL)
	if err != nil {
		return nil
	}
	return &HttpTestCaseParser{r: response.Body, parser: InputOutputParse(parse)}

}

func (p HttpTestCaseParser) Parse() (TestCaseSet, error) {
	return p.parser.Parse(p.r)
}
