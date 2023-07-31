package function

import (
	"strings"
	"time"

	"github.com/zenmuharom/zenlogger"
)

// default layout
var dateLayoutPatterns = []string{
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

func (assigner *DefaultAssigner) DateFormat(oldFormat string, theDate string, newFormat string) (formatted string, err error) {
	// Parse date
	date, err := time.Parse(strings.TrimSpace(oldFormat), theDate)
	if err != nil {
		assigner.Logger.Error("DateFormat", zenlogger.ZenField{Key: "error", Value: err.Error()})
	}

	formatted = date.Format(newFormat)

	return
}
