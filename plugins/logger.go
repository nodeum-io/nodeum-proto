package plugins

import (
	"context"
	"os"

	"github.com/hashicorp/go-hclog"
	commonv1 "github.com/nodeum-io/nodeum-proto/nodeum/common/v1"
)

func NewLogger(exec *commonv1.Execution) hclog.Logger {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: true,
	})

	if id := exec.GetId(); id != "" {
		logger = logger.With("execution_id", id)
	}
	if id := exec.GetTask().GetId(); id != "" {
		logger = logger.With("task_id", id)
	}

	return logger
}

type loggerKey struct{}

func ExtractLogger(ctx context.Context) hclog.Logger {
	if logger, ok := ctx.Value(loggerKey{}).(hclog.Logger); ok {
		return logger
	}
	return NewLogger(nil)
}

func InjectLogger(ctx context.Context, logger hclog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}
