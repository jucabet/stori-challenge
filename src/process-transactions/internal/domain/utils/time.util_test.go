package utils

import "testing"

func TestGetCurrentTime(t *testing.T) {
	time := GetCurrentTime()
	if time.Year() == 0 {
		t.Fatal("errorGetCurrentTime")
	}
}
