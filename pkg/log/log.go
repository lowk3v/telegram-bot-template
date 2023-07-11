package log

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

type Logger struct {
	*logrus.Logger
	symbol SymbolConfig
}

type SymbolConfig struct {
	Success string
	Error   string
	Info    string
	Debug   string
}

func New(level string) *Logger {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		lvl = logrus.InfoLevel
	}

	l := &Logger{
		Logger: logrus.New(),
		symbol: SymbolConfig{
			Success: color.GreenString("≠"),
			Error:   color.RedString("¿"),
			Info:    color.BlueString("ℹ"),
			Debug:   color.MagentaString("☣"),
		},
	}
	l.SetReportCaller(true)
	l.SetLevel(lvl)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			fileParts := strings.Split(f.File, "/")
			funcParts := strings.Split(f.Function, "/")
			return fmt.Sprintf("%s()", funcParts[len(funcParts)-1]),
				fmt.Sprintf("%s:%d", fileParts[len(fileParts)-1], f.Line)
		},
	}
	return l
}
