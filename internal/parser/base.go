package parser

import (
	"io"
	"strings"
)

type InputOutputParse func(input io.Reader) ([]TestCase, error)

func (p InputOutputParse) Parse(reader io.Reader) ([]TestCase, error) {
	return p(reader)
}

func parse(r io.Reader) (cases []TestCase, err error) {

	textBytes, err := io.ReadAll(r)
	if err != nil {
		return nil, err

	}

	content := string(textBytes)

	for {
		inputStart := strings.Index(content, "input:")
		if inputStart == -1 {
			break
		}

		inputContent := content[inputStart+len("input:"):]

		outputStart := strings.Index(inputContent, "output:")
		if outputStart == -1 {
			break
		}

		inputVal := strings.TrimSpace(inputContent[:outputStart])

		outputContent := inputContent[outputStart+len("output:"):]

		nextInput := strings.Index(outputContent, "input:")

		var outputVal string
		if nextInput == -1 {
			outputVal = strings.TrimSpace(outputContent)
		} else {
			outputVal = strings.TrimSpace(outputContent[:nextInput])
		}

		cases = append(cases, TestCase{
			Input:  inputVal,
			Output: outputVal,
		})

		if nextInput == -1 {
			break
		}
		content = outputContent[nextInput:]
	}

	return cases, nil

}
