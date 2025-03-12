package report

import (
	"windy-judge/internal/F"
)

type Render interface {
	Printer() F.Printer
	Write(p []byte) (n int, err error)
	Beauty()
}
