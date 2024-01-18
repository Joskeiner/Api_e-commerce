package helper

import (
	"time"
)

const (
	ddmmyyyyLayout = "02/01/2006"
)

func ParseTime(date string) (time.Time, error) {
	t, err := time.Parse(ddmmyyyyLayout, date)
	if err != nil {
		return t, err
	}
	return t, nil
}
