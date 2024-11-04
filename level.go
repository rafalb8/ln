package log

import (
	"fmt"
)

type Level int

const (
	LevelTrace Level = iota*4 - 8
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

func (l Level) String() string {
	switch l {
	case LevelTrace:
		return "TRACE"
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return fmt.Sprintf("INFO%+d", l)
	}
}

// Colors used by TextHandler renderer
func (l Level) Color() string {
	switch l {
	case LevelTrace:
		return "\x1b[37;1mTRACE\x1b[0m"
	case LevelDebug:
		return "\x1b[35;1mDEBUG\x1b[0m"
	case LevelInfo:
		return "\x1b[34;1mINFO\x1b[0m"
	case LevelWarn:
		return "\x1b[33;1mWARN\x1b[0m"
	case LevelError:
		return "\x1b[31;1mERROR\x1b[0m"
	case LevelFatal:
		return "\x1b[30;41;1mFATAL\x1b[0m"
	default:
		return fmt.Sprintf("\x1b[34;1mINFO%+d\x1b[0m", l)
	}
}

func LevelFrom(str string) Level {
	switch str[0] | 32 {
	case 't':
		return LevelTrace
	case 'd':
		return LevelDebug
	case 'w':
		return LevelWarn
	case 'e':
		return LevelError
	case 'f':
		return LevelFatal
	default:
		return LevelInfo
	}
}
