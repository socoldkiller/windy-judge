package report

type Render interface {
	Write(data any) error
	Beauty()
}
