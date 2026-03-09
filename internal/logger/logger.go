package logger

import (
	"log/slog"
	"os"
)

var Log *slog.Logger

func Init(logFile string) error {

	file, err := os.OpenFile(
		logFile,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		return err
	}

	handler := slog.NewTextHandler(file, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	Log = slog.New(handler)

	return nil
}
