package utils

import "time"

var GetCurrentTime = func() time.Time {
	now := time.Now()
	return now
}
