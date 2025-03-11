package report

import "io"

func ReadAll(r io.Reader) string {
	var (
		s   []byte
		err error
	)
	if s, err = io.ReadAll(r); err != nil {
		return ""
	}
	return string(s)
}
