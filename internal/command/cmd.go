package command

import (
	"bytes"
	"os/exec"
	"strings"
	"time"
)

func MaybeErrorMessage(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}

type Result struct {
	Input   string
	Output  string
	Error   string
	Elapsed time.Duration
}

type Cmd struct {
	cmd  string
	args []string

	errCode int
}

func (c *Cmd) ErrCode() int {
	return c.errCode
}

func NewCmd(cmd string, args ...string) *Cmd {
	return &Cmd{cmd: cmd, args: args}
}

func (c *Cmd) Run(input string) Result {
	var buf bytes.Buffer

	cmd := exec.Command(c.cmd, c.args...)
	cmd.Stdin = strings.NewReader(input)
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	start := time.Now()
	e := cmd.Run()
	end := time.Now()

	errMessage := MaybeErrorMessage(e)
	if errMessage != "" {
		c.errCode = 1
	}

	return Result{
		Input:   input,
		Output:  strings.TrimSpace(buf.String()),
		Error:   errMessage,
		Elapsed: end.Sub(start),
	}

}
