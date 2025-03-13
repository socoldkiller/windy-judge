package report

import (
	"windy-judge/internal/F"
)

type Render interface {
	Printer() F.Printer
	Write(data any) error
	Beauty()
}
