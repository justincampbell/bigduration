package bigduration

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	minute = 1
	hour   = minute * 60
	day    = hour * 24
	week   = day * 7
	month  = day * 30
	year   = day * 365
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

	minutes := years*year + months*month + weeks*week + days*day

	add, err := time.ParseDuration(fmt.Sprintf("%0dm", minutes))
	if err != nil {
		return nilDuration, err
	}

	if s == "" {
		return add, nil
	}

	d, err := time.ParseDuration(s)
	if err != nil {
		return nilDuration, err
	}

	return d + add, nil
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
