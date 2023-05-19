package function

import (
	"time"

	"github.com/zenmuharom/zenlogger"
)

func (assigner *DefaultAssigner) DateNow(format string) (generated string, err error) {
	assigner.logger.Debug("DateNow", zenlogger.ZenField{Key: "format", Value: format})
	date := time.Now()
	if format == "" {
		format = time.RFC3339
	}
	generated = date.Format(format)
	return
}
