package plugins

import (
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
