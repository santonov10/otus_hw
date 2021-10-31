package logger

import (
	"log"
	"os"
	"sync"

	"github.com/santonov10/otus_hw/hw12_13_14_15_calendar/internal/config"

	"github.com/rs/zerolog"
)

type Logger struct {
	*zerolog.Logger
}

var (
	logger      Logger
	once        sync.Once
	logFilePath = "./logfile.txt"
	useFile     = true
)

// Get reads config from environment. Once.
func Get() *Logger {
	once.Do(func() {
		cfg := config.Get()
		var zeroLogger zerolog.Logger
		if useFile {
			f, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
			if err != nil {
				log.Fatalf("ошибка инициализации логгера:%v", err)
			}
			zeroLogger = zerolog.New(f).With().Timestamp().Logger()
		} else {
			zeroLogger = zerolog.New(os.Stderr).With().Timestamp().Logger()
		}

		switch cfg.Logger.Level {
		case "debug":
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		case "info":
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		case "warn", "warning":
			zerolog.SetGlobalLevel(zerolog.WarnLevel)
		case "err", "error":
			zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		case "fatal":
			zerolog.SetGlobalLevel(zerolog.FatalLevel)
		case "panic":
			zerolog.SetGlobalLevel(zerolog.PanicLevel)
		default:
			zerolog.SetGlobalLevel(zerolog.InfoLevel) // log info and above by default
		}
		logger = Logger{&zeroLogger}
	})
	return &logger
}

func UseFile(value bool) {
	useFile = value
}
