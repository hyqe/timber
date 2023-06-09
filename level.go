package timber

import "strings"

// Level indicates by convention the type of log being generated
// and how it would be handled.
//
// log levels range from SILENT to DEBUG, where each level will
// also permit the previous log levels to be written.
type Level int

const (
	// SILENT no logs.
	SILENT Level = iota

	// ALERT is used to log system level alerts. Alerts
	// are generally not generated directly by user level
	// events or requests.
	//
	// examples:
	//  - starting or stopping a service.
	//  - health status changes
	//  - config changes
	//  - os signals
	ALERT

	// ERROR logs errors usually generated at the user level
	// within your application.
	//
	// examples:
	// - 500 http status responses
	// - file read misses
	// - unexpected state
	ERROR

	// DEBUG logs all the things.
	DEBUG
)

func ParseLevel(lvl string) Level {
	switch {
	case strings.EqualFold(lvl, "DEBUG"):
		return DEBUG
	case strings.EqualFold(lvl, "ERROR"):
		return ERROR
	case strings.EqualFold(lvl, "ALERT"):
		return ALERT
	case strings.EqualFold(lvl, "SILENT"):
		return SILENT
	default:
		return SILENT
	}
}

func (l Level) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case ERROR:
		return "ERROR"
	case ALERT:
		return "ALERT"
	case SILENT:
		return "SILENT"
	default:
		return "SILENT"
	}
}

// Is reports whether a level is within another level
//
//	DEBUG.Is(ALERT) == true
//	ALERT.Is(ERROR) == false
func (l Level) GTE(lvl Level) bool {
	return l >= lvl
}
