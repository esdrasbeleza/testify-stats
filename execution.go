package testifystats

import "time"

type Execution struct {
	Id        string
	SuiteName string
	Start     time.Time
	End       time.Time
}
