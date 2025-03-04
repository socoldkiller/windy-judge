package command

import "io"

type TestCase struct {
	Input  string
	Output string
}

type TestCaseSet = []TestCase

type ReaderParser interface {
	Parse(reader io.Reader) (TestCaseSet, error)
}

type TestCaseParser interface {
	Parse() (TestCaseSet, error)
}
