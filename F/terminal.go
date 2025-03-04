package F

import (
	"github.com/gookit/color"
)

// 定义全局颜色样式
var (
	ErrorStyle   = color.New(color.FgRed, color.Bold)    // 错误样式（红色粗体）
	SuccessStyle = color.New(color.FgGreen, color.Bold)  // 成功样式（绿色粗体）
	InfoStyle    = color.New(color.FgBlue, color.Bold)   // 信息样式（蓝色粗体）
	WarnStyle    = color.New(color.FgYellow)             // 警告样式（黄色）
	TimeStyle    = color.New(color.FgCyan)               // 时间样式（青色）
	DiffStyle    = color.New(color.BgRed, color.FgWhite) // 差异样式（红底白字）
	DefaultStyle = color.New(color.White)
)

type Terminal struct {
}

func (t Terminal) Beauty() {
	//TODO implement me
	panic("implement me")
}

func (t Terminal) KeyValueFormat(format string, key string, value ...any) {
	t.Timef("%s", key)
	DefaultStyle.Printf(format, value...)
}

func (t Terminal) Time(a ...any) {
	//TODO implement me
	panic("implement me")
}

// 带样式的打印方法
func (t Terminal) Error(a ...any) {
	ErrorStyle.Print(a...)
}

func (t Terminal) Errorln(a ...any) {
	ErrorStyle.Println(a...)
}

func (t Terminal) Success(a ...any) {
	SuccessStyle.Print(a...)
}

func (t Terminal) Successln(a ...any) {
	SuccessStyle.Println(a...)
}

func (t Terminal) Info(a ...any) {
	InfoStyle.Print(a...)
}

func (t Terminal) Infoln(a ...any) {
	InfoStyle.Println(a...)
}

func (t Terminal) Warn(a ...any) {
	WarnStyle.Print(a...)
}

func (t Terminal) Warnln(a ...any) {
	WarnStyle.Println(a...)
}

func (t Terminal) Timef(format string, a ...any) {
	TimeStyle.Printf(format, a...)
}

func (t Terminal) Timeln(a ...any) {
	TimeStyle.Println(a...)
}

func (t Terminal) Diff(a ...any) {
	DiffStyle.Print(a...)
}

func (t Terminal) Diffln(a ...any) {
	DiffStyle.Println(a...)
}

func (t Terminal) Defaultln(a ...any) {
	DefaultStyle.Println(a...)
}

func (t Terminal) TitleTimeF(format, title string, a ...any) {
	t.Timef("%s", title)
	DefaultStyle.Printf(format, a...)
}
