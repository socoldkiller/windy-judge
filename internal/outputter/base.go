package outputter

type OutPutter interface {
	Error(a ...any)
	Errorln(a ...any)
	Success(a ...any)
	Successln(a ...any)
	Info(a ...any)
	Infoln(a ...any)
	Warn(a ...any)
	Warnln(a ...any)
	Time(a ...any)
	Timeln(a ...any)
	KeyValueFormat(format string, key string, value ...any)
	Defaultln(a ...any)
	Beauty()
}

type Beauty interface {
	Beauty()
}
