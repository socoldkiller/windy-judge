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

	return Result{
		Input:   input,
		Output:  strings.TrimSpace(buf.String()),
		Error:   MaybeErrorMessage(e),
		Elapsed: end.Sub(start),
	}

}
