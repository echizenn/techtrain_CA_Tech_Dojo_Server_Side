package errors

type Level string

const (
	Debug  Level = "Debug"
	Info   Level = "Info"
	Warn   Level = "Warn"
	Error  Level = "Error"
	DPanic Level = "DPanic"
	Panic  Level = "Panic"
	Fatal  Level = "Fatal"
)
