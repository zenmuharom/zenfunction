package function

import (
	"fmt"
	"strings"
	"time"

	"github.com/zenmuharom/zenlogger"
)

func (assigner *DefaultAssigner) DateAdd(format string, theDate string, add int, duration string) (added string, err error) {
	assigner.Logger.Debug("DateAdd", zenlogger.ZenField{Key: "format", Value: format})

	// Parse the ISO date string
	invalid := false
	date, err := time.Parse(strings.TrimSpace(format), theDate)
	if err != nil {
		assigner.Logger.Error("DateAdd", zenlogger.ZenField{Key: "error", Value: err.Error()})
		invalid = true
	}

	if invalid {
		// Attempt to parse the date string using predefined layouts
		for _, layout := range dateLayoutPatterns {
			date, err = time.Parse(layout, theDate)
			if err == nil {
				invalid = false
				break
			}
		}
	}

	if err != nil {
		date = time.Now()
		assigner.Logger.Error("DateAdd", zenlogger.ZenField{Key: "error", Value: err.Error()}, zenlogger.ZenField{Key: "handling", Value: fmt.Sprintf("set date to %v", date)})
	}

	var newDate time.Time

	switch duration {
	case "year":
		newDate = date.AddDate(add, 0, 0)
	case "month":
		newDate = date.AddDate(0, add, 0)
	case "day":
		newDate = date.AddDate(0, 0, add)
	case "hour":
		addDuration := time.Duration(int(time.Hour) * add)
		// Add duration to the date
		newDate = date.Add(addDuration)
	case "minute":
		addDuration := time.Duration(int(time.Minute) * add)
		// Add duration to the date
		newDate = date.Add(addDuration)
	case "second":
		addDuration := time.Duration(int(time.Second) * add)
		// Add duration to the date
		newDate = date.Add(addDuration)

	}

	// Format the new date back to ISO format
	added = newDate.Format(format)

	return
}
