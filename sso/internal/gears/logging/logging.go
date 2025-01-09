package logging

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/shotowon/kon.io/sso/internal/gears/config"
)

func SetupLogger(env string) (*slog.Logger, error) {
	var logger *slog.Logger

	switch env {
	case config.EnvLocal:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	case config.EnvDev:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	case config.EnvProd:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	default:
		return nil, fmt.Errorf("logging: no log setup for %s env", env)
	}

	return logger, nil
}
