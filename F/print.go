// T/color.go
package T

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

// 带样式的打印方法
func Error(a ...any) {
	ErrorStyle.Print(a...)
}

func Errorln(a ...any) {
	ErrorStyle.Println(a...)
}

func Success(a ...any) {
	SuccessStyle.Print(a...)
}

func Successln(a ...any) {
	SuccessStyle.Println(a...)
}

func Info(a ...any) {
	InfoStyle.Print(a...)
}

func Infoln(a ...any) {
	InfoStyle.Println(a...)
}

func Warn(a ...any) {
	WarnStyle.Print(a...)
}

func Warnln(a ...any) {
	WarnStyle.Println(a...)
}

func Timef(format string, a ...any) {
	TimeStyle.Printf(format, a...)
}

func Timeln(a ...any) {
	TimeStyle.Println(a...)
}

func Diff(a ...any) {
	DiffStyle.Print(a...)
}

func Diffln(a ...any) {
	DiffStyle.Println(a...)
}

func Defaultln(a ...any) {
	DefaultStyle.Println(a...)
}

func TitleTimeF(format, title string, a ...any) {
	Timef("%s", title)
	DefaultStyle.Printf(format, a...)
}
