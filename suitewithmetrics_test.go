package testifystats

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestSampleMetrics(t *testing.T) {
	suite.Run(t, new(SampleMetricsSuite))
}

type SampleMetricsSuite struct {
	suite.Suite
}

type suiteUnderTest struct {
	suite.Suite
	SuiteWithMetrics
}

func (s *suiteUnderTest) Test_Something() {
	s.Equal(1, 1)
}

func (s *SampleMetricsSuite) Test_ANewSuiteHasAnExecutionIdAndStartTime() {
	suiteUnderTest := new(suiteUnderTest)
	suiteUnderTest.SetupSuite()

	s.NotEmpty(suiteUnderTest.Execution.Id)
	s.NotZero(suiteUnderTest.Execution.Start)
	s.Zero(suiteUnderTest.Execution.End)
	s.NotNil(suiteUnderTest.Stats)
}

func (s *SampleMetricsSuite) Test_TearDownSuiteUpdatesTheExecutionEndTime() {
	suiteUnderTest := new(suiteUnderTest)
	suiteUnderTest.TearDownSuite()

	s.NotZero(suiteUnderTest.Execution.End)
}

func (s *SampleMetricsSuite) Test_BeforeTest_CreatesNewStats() {
	suiteUnderTest := new(suiteUnderTest)
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
	suiteUnderTest := new(suiteUnderTest)
	suiteUnderTest.SetupSuite()
	suiteUnderTest.BeforeTest("mySuite", "myTest")
	suiteUnderTest.AfterTest("mySuite", "myTest")

	actualStats := suiteUnderTest.Stats["myTest"]

	s.NotZero(actualStats.End)
}

// To test a suite, ironically, this test can't be in a suite.Suite.
// Bad things happen when you test a suite.Suite inside a suite.
func Test_RunningTheSuiteWillGenerateStats(t *testing.T) {
	suiteToTest := new(suiteUnderTest)

	suite.Run(t, suiteToTest)

	assert.NotEmpty(t, suiteToTest.Execution.Id)
	assert.NotZero(t, suiteToTest.Execution.Start)
	assert.NotZero(t, suiteToTest.Execution.End)
	assert.NotNil(t, suiteToTest.Stats)
}
