package testifystats

import "time"

type Stats struct {
	SuiteName string
	Testname  string
	Start     time.Time
	End       time.Time
	Success   bool
}
