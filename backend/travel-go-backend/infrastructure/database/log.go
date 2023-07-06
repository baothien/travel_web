package pg

type LogLevel int

const (
	Debug LogLevel = iota + 1
	Info
	Warn
	Error
	Trace
)
