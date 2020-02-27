package testifystats

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSampleMetrics(t *testing.T) {
	suite.Run(t, new(SampleMetricsSuite))
}

type SampleMetricsSuite struct {
	suite.Suite
}

type suiteUnderTest struct {
	SuiteWithMetrics
}

func (s *SampleMetricsSuite) Test_ANewSuiteHasAnExecutionIdAndStartTime() {
	suiteUnderTest := suiteUnderTest{}
	suiteUnderTest.SetupSuite()

	s.NotEmpty(suiteUnderTest.Execution.Id)
	s.NotZero(suiteUnderTest.Execution.Start)
	s.Zero(suiteUnderTest.Execution.End)
	s.NotNil(suiteUnderTest.Stats)
}

func (s *SampleMetricsSuite) Test_TearDownSuiteUpdatesTheExecutionEndTime() {
	suiteUnderTest := suiteUnderTest{SuiteWithMetrics{}}
	suiteUnderTest.TearDownSuite()

	s.NotZero(suiteUnderTest.Execution.End)
}

func (s *SampleMetricsSuite) Test_BeforeTest_CreatesNewStats() {
	suiteUnderTest := suiteUnderTest{SuiteWithMetrics{}}
	suiteUnderTest.SetupSuite()
	suiteUnderTest.BeforeTest("mySuite", "myTest")

	actualStats := suiteUnderTest.Stats["myTest"]

	s.NotEmpty(actualStats.Id)
	s.Equal("mySuite", actualStats.SuiteName)
	s.Equal("myTest", actualStats.Testname)
	s.NotZero(actualStats.Start)
	s.Zero(actualStats.End)
}

func (s *SampleMetricsSuite) Test_AfterTest_UpdatesStats() {
	suiteUnderTest := suiteUnderTest{SuiteWithMetrics{}}
	suiteUnderTest.SetupSuite()
	suiteUnderTest.BeforeTest("mySuite", "myTest")
	suiteUnderTest.AfterTest("mySuite", "myTest")

	actualStats := suiteUnderTest.Stats["myTest"]

	s.NotZero(actualStats.End)
}
