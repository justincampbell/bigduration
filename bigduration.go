package bigduration

import (
	"strconv"
	"strings"
	"time"
)

// Additional time.Duration constants
const (
	Day   = time.Hour * 24
	Week  = Day * 7
	Month = Day * 30
	Year  = Day * 365
)

// ParseDuration takes a string in the same format as time.ParseDuration, but
// also supports additonal duration suffixes:
//
//   D day
//   W week
//   M month
//   Y year
//
func ParseDuration(s string) (time.Duration, error) {
	var nilDuration time.Duration

	s, years, err := splitDuration(s, "Y")
	if err != nil {
		return nilDuration, err
	}
	s, months, err := splitDuration(s, "M")
	if err != nil {
		return nilDuration, err
	}
	s, weeks, err := splitDuration(s, "W")
	if err != nil {
		return nilDuration, err
	}
	s, days, err := splitDuration(s, "D")
	if err != nil {
		return nilDuration, err
	}

	big := time.Duration(years)*Year +
		time.Duration(months)*Month +
		time.Duration(weeks)*Week +
		time.Duration(days)*Day

	if s == "" {
		return big, nil
	}

	little, err := time.ParseDuration(s)
	if err != nil {
		return nilDuration, err
	}

	return big + little, nil
}

func splitDuration(s string, delim string) (string, int, error) {
	elems := strings.Split(s, delim)
	if len(elems) == 2 {
		i64, err := strconv.ParseInt(elems[0], 10, 64)
		if err != nil {
			return "", 0, err
		}
		return elems[len(elems)-1], int(i64), nil
	}

	return s, 0, nil
}
