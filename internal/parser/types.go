package parser

import (
	"io"
	"windy-judge/internal"
)

type TestCase = internal.TestCase

type ReaderParser[O any] interface {
	Parse(reader io.Reader) ([]O, error)
}

type TestCaseParser[O any] interface {
	Parse() ([]O, error)
}
