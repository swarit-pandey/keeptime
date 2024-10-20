package common

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// GetLocation fetches the location for the given zone
func GetLocation(tz string) (*time.Location, error) {
	loc, exists := TimeZones[strings.ToUpper(tz)]
	if !exists {
		return nil, fmt.Errorf("%v - time zone not supported", tz)
	}

	tzLoc, err := time.LoadLocation(loc)
	if err != nil {
		return nil, fmt.Errorf("failed to load location %v: %v", tzLoc.String(), err)
	}

	return tzLoc, nil
}

// ParseTimeLoc parses time with default layout and given time zone
// layout is YYYYMMDDHHMMSS
func ParseTimeLoc(input string, tzLoc *time.Location) (time.Time, error) {
	if strings.ToLower(input) == "now" {
		return time.Now().In(tzLoc), nil
	}

	t, err := time.ParseInLocation(DefaultLayout, input, tzLoc)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse time: %v", err)
	}
	return t, nil
}

// ParseRelativeTime will parse relative time given in input looking back
// at time, for example if 1y is given it means we need to look back a year
func ParseRelativeTime(input string, tzLoc *time.Location) (time.Time, error) {
	re := regexp.MustCompile(`^(\d+)([smhdy])$`)
	matches := re.FindStringSubmatch(strings.ToLower(input))
	if len(matches) != 3 {
		return time.Time{}, fmt.Errorf("invalid relative time format: %s", input)
	}

	value, err := strconv.Atoi(matches[1])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid time in relative number: %s", input)
	}

	unit := matches[2]
	now := time.Now().In(tzLoc)

	switch unit {
	case "s":
		return now.Add(time.Duration(-value) * time.Second), nil
	case "h":
		return now.Add(time.Duration(-value) * time.Hour), nil
	case "d":
		return now.AddDate(0, 0, -value), nil
	case "m":
		return now.AddDate(0, -value, 0), nil
	case "y":
		return now.AddDate(-value, 0, 0), nil
	default:
		return time.Time{}, fmt.Errorf("unsupported unit in relative time: %v", unit)
	}
}
