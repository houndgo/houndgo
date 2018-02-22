package itime

import "time"

// Today is
func Today() string {
	return time.Now().Format("2006-01-02")
}
