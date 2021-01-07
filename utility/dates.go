package utility

import "time"

// MonthToTime creates a time from a month
func MonthToTime(month int) time.Time {
	now := time.Now()
	return time.Date(
		now.Year(),
		time.Month(month),
		1,
		0, 0, 0, 0, time.UTC)
}
