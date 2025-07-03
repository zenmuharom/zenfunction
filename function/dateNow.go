package function

import (
	"fmt"
	"time"

	"github.com/zenmuharom/zenlogger"
)

func (assigner *DefaultAssigner) DateNow(format string) (generated string, err error) {
	assigner.Logger.Debug("DateNow", zenlogger.ZenField{Key: "format", Value: format})
	now := time.Now()
	if format == "" {
		format = time.RFC3339
	}

	switch format {
	case "", "rfc3339":
		return now.Format(time.RFC3339), nil
	case "unix":
		return fmt.Sprintf("%d", now.Unix()), nil
	case "unixMilli":
		return fmt.Sprintf("%d", now.UnixMilli()), nil
	case "unixNano":
		return fmt.Sprintf("%d", now.UnixNano()), nil
	default:
		return now.Format(format), nil
	}

	return
}
