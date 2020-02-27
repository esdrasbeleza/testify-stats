package testifystats

import "time"

type Stats struct {
	Id        string
	SuiteName string
	Testname  string
	Start     time.Time
	End       time.Time
}
