package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
	"sync"
	"time"
	"tutorial/pkg/config"
)

var (
	once          sync.Once
	defaultLogger *slog.Logger
)

type CustomHandler struct {
	slog.Handler
}

func GetLogger() *slog.Logger {
	return defaultLogger
}

func LoadLogger(cfg *config.Config) {
	once.Do(func() {
		var log slog.Level

		switch strings.ToLower(cfg.Logger.Level) {
		case "debug":
			log = slog.LevelDebug
		case "info":
			log = slog.LevelInfo
		case "warn":
			log = slog.LevelWarn
		case "error":
			log = slog.LevelError
		default:
			log = slog.LevelInfo
		}

		opts := &slog.HandlerOptions{
			Level: log,
		}

		var handler slog.Handler

		if strings.ToLower(cfg.Logger.Format) == "console" {
			handler = CustomHandler{slog.NewTextHandler(io.Discard, opts)}
		} else {
			handler = slog.NewJSONHandler(os.Stdout, opts)
		}

		defaultLogger = slog.New(handler)
		slog.SetDefault(defaultLogger)

	})
}

func (h CustomHandler) Handle(ctx context.Context, r slog.Record) error {
	level := "INF"
	color := ColorGreen
	switch r.Level {
	case slog.LevelDebug:
		level = "DBG"
		color = ColorCyan
	case slog.LevelWarn:
		level = "WRN"
		color = ColorYellow
	case slog.LevelError:
		level = "ERR"
		color = ColorRed
	}

	timeStr := r.Time.Format("2006/01/02 15:04:05")

	fmt.Printf("%s%s%s %s[%s]%s %s", ColorGray, timeStr, ColorReset, color, level, ColorReset, r.Message)

	var (
		method                           string
		status                           int
		latency                          time.Duration
		hasMethod, hasStatus, hasLatency bool
		otherAttrs                       []slog.Attr
	)

	r.Attrs(func(a slog.Attr) bool {
		switch a.Key {
		case "method":
			method = a.Value.String()
			hasMethod = true
		case "status":
			switch v := a.Value.Any().(type) {
			case int:
				status = v
			case int64:
				status = int(v)
			}
			hasStatus = true
		case "latency":
			latency = a.Value.Any().(time.Duration)
			hasLatency = true
		default:
			otherAttrs = append(otherAttrs, a)
		}
		return true
	})

	if hasMethod && hasStatus {
		methodColor := ColorCyan
		switch method {
		case "POST":
			method = "PST"
			methodColor = ColorGreen
		case "PATCH":
			method = "PTC"
			methodColor = ColorYellow
		case "DELETE":
			method = "DEL"
			methodColor = ColorRed
		case "PUT":
			methodColor = ColorYellow
		}

		statusColor := ColorGreen
		if status >= 400 && status < 500 {
			statusColor = ColorYellow
		} else if status >= 500 {
			statusColor = ColorRed
		}

		fmt.Printf(" %s%s%s %s%d%s", methodColor, method, ColorReset, statusColor, status, ColorReset)
	}

	//  if latency ≥ 1s print red color in latency section.
	if hasLatency {
		latencyColor := ColorGreen
		if latency >= time.Second {
			latencyColor = ColorRed
		}
		fmt.Printf(" %s%v%s", latencyColor, latency, ColorReset)
	}

	// other key-value attributes
	for _, attr := range otherAttrs {
		fmt.Printf(" %s%s:%v%s", ColorGray, attr.Key, attr.Value, ColorReset)
	}

	fmt.Println()
	return nil
}
