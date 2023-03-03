package log

import (
	"os"
	"runtime"
	"strings"

	"github.com/rs/zerolog"
)

var root zerolog.Logger

func SetupLogging(options ...Option) {
	lc := &logConfig{}
	for _, o := range options {
		o(lc)
	}

	stream := lc.stream
	if stream == nil {
		// default to stderr
		stream = os.Stderr
	}

	if lc.console {
		stream = zerolog.ConsoleWriter{Out: stream}
	}

	// TODO better timestamp formats

	l := zerolog.New(stream)

	x := l.Level(lc.level).With().Caller().Timestamp()
	// TODO appname
	root = x.Logger()
}

func GetNamed(name string) zerolog.Logger {
	return root.With().Str("name", name).Logger()
}

func Get() zerolog.Logger {
	pc, _, _, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()

	parts := strings.Split(funcName, "/")
	last := parts[len(parts)-1]
	return GetNamed(last)
}
