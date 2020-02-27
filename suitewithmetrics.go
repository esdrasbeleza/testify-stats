package testifystats

import (
	"time"

	"github.com/google/uuid"
)

type SuiteWithMetrics struct {
	Execution Execution
	Stats     map[string]*Stats
}

func (sm *SuiteWithMetrics) SetupSuite() {
	sm.Execution = Execution{
		Id:    uuid.New().String(),
		Start: time.Now(),
		End:   time.Time{},
	}

	sm.Stats = make(map[string]*Stats)
}

func (sm *SuiteWithMetrics) TearDownSuite() {
	sm.Execution.End = time.Now()
}

func (sm *SuiteWithMetrics) BeforeTest(suiteName, testName string) {
	stats := &Stats{
		Id:        uuid.New().String(),
		SuiteName: suiteName,
		Testname:  testName,
		Start:     time.Now(),
	}

	sm.Stats[testName] = stats
}

func (sm *SuiteWithMetrics) AfterTest(suiteName, testName string) {
	sm.Stats[testName].End = time.Now()
}
