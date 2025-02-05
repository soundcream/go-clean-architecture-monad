package datetimeutil

import "time"

func Now() time.Time {
	return time.Now()
}

func NowPtr() *time.Time {
	result := Now()
	return &result
}
