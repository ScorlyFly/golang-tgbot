package logger

type IsDebug struct {
	Debug bool
}

type LevelLogger interface {
	Info()
}

func LoadLogger() *IsDebug {
	return &IsDebug{
		Debug: true,
	}
}

func Info(msg string, v ...interface{}) {

}
