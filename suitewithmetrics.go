package testifystats

import (
	"time"
)

type SuiteWithMetrics struct {
	Execution Execution
	Stats     map[string]Stats
	OnFinish  func(Execution, map[string]Stats)
}

func (sm *SuiteWithMetrics) SetupSuite() {
	sm.Execution = Execution{
		Start: time.Now(),
		End:   time.Time{},
	}

	sm.Stats = make(map[string]Stats)
}

func (sm *SuiteWithMetrics) TearDownSuite() {
	sm.Execution.End = time.Now()

	if sm.OnFinish != nil {
		sm.OnFinish(sm.Execution, sm.Stats)
	}
}

func (sm *SuiteWithMetrics) BeforeTest(suiteName, testName string) {
	sm.Stats[testName] = Stats{
		SuiteName: suiteName,
		Testname:  testName,
		Start:     time.Now(),
	}
}

func (sm *SuiteWithMetrics) AfterTest(suiteName, testName string, success bool) {
	stats := sm.Stats[testName]
	stats.End = time.Now()
	stats.Success = success

	sm.Stats[testName] = stats
}
