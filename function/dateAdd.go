package function

import (
	"fmt"
	"strings"
	"time"

	"github.com/zenmuharom/zenlogger"
)

func (assigner *DefaultAssigner) DateAdd(format string, theDate string, add int, duration string) (added string, err error) {
	assigner.logger.Debug("DateAdd", zenlogger.ZenField{Key: "format", Value: format})

	// default layout
	layoutPatterns := []string{
		time.RFC3339,          // "2006-01-02T15:04:05Z07:00"
		time.RFC3339Nano,      // "2006-01-02T15:04:05.999999999Z07:00"
		time.RFC1123,          // "Mon, 02 Jan 2006 15:04:05 MST"
		time.RFC1123Z,         // "Mon, 02 Jan 2006 15:04:05 -0700"
		time.RFC822,           // "02 Jan 06 15:04 MST"
		time.RFC822Z,          // "02 Jan 06 15:04 -0700"
		time.RFC850,           // "Monday, 02-Jan-06 15:04:05 MST"
		time.ANSIC,            // "Mon Jan _2 15:04:05 2006"
		"2006",                // 2006
		"01",                  // 01
		"02",                  // 02
		"15",                  // 05
		"04",                  // 04
		"05",                  // 05
		"200601",              // 200601
		"2006-01",             // 2006-01
		"2006/01",             // 2006/01
		"2006-01-02",          // 2006-01-02
		"2006/01/02",          // 2006/01/02
		"20060102",            // 20060102
		"2006-01-02 15:04:05", // 2006-01-02 15:04:05
		"2006-01-02 15-04-05", // 2006-01-02 15:04:05
		"2006-01-02 15/04/05", // 2006-01-02 15:04:05
		"2006/01/02 15:04:05", // 2006/01/02 15:04:05
		"2006/01/02 15-04-05", // 2006/01/02 15-04-05
		"2006/01/02 15/04/05", // 2006/01/02 15/04/05
		"2006/01/02 15 04 05", // 2006/01/02 15 04 05
		"2006 01 02 15:04:05", // 2006 01 02 15:04:05
		"2006 01 02 15 04 05", // 2006 01 02 15 04 05
		"20060102 150405",     // 20060102 150405
		"20060102150405",      // 20060102150405
		"15:04:05",            // 15:04:05 hh:mm:ss
		"15/04/05",            // 15/04/05 hh/mm/ss
		"15-04-05",            // 15-04-05
		"15 04 05",            // 15 04 05 hh mm ss
		"15:04",               // 15:04 hh:mm
		"15/04",               // 15/04 hh/mm
		"15-04",               // 15-04 hh-mm
		"05 04",               // 15 04 hh mm
		"04:05",               // 04:05 mm:ss
		"04/05",               // 04:05 mm/ss
		"04-05",               // 04:05 mm-ss
		"04 05",               // 04:05 mm ss
	}

	// Parse the ISO date string
	invalid := false
	date, err := time.Parse(strings.TrimSpace(format), theDate)
	if err != nil {
		assigner.logger.Error("DateAdd", zenlogger.ZenField{Key: "error", Value: err.Error()})
		invalid = true
	}

	if invalid {
		// Attempt to parse the date string using predefined layouts
		for _, layout := range layoutPatterns {
			date, err = time.Parse(layout, theDate)
			if err == nil {
				invalid = false
				break
			}
		}
	}

	if err != nil {
		date = time.Now()
		assigner.logger.Error("DateAdd", zenlogger.ZenField{Key: "error", Value: err.Error()}, zenlogger.ZenField{Key: "handling", Value: fmt.Sprintf("set date to %v", date)})
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
