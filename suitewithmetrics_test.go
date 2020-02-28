package testifystats

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// To test a suite, ironically, these tests can't be in a suite.Suite.
// Bad things happen when you test a suite.Suite inside a suite!

type suiteUnderTest struct {
	suite.Suite
	SuiteWithMetrics
}

func (s *suiteUnderTest) Test_Something() {
	s.Equal(1, 1)
}

func Test_ANewSuiteHasAnExecutionIdAndStartTime(t *testing.T) {
	suiteUnderTest := new(suiteUnderTest)
	suiteUnderTest.SetupSuite()

	assert.NotZero(t, suiteUnderTest.Execution.Start)
	assert.Zero(t, suiteUnderTest.Execution.End)
	assert.NotNil(t, suiteUnderTest.Stats)
}

func Test_TearDownSuite_UpdatesTheExecutionEndTime(t *testing.T) {
	suiteUnderTest := new(suiteUnderTest)
	suiteUnderTest.TearDownSuite()

	assert.NotZero(t, suiteUnderTest.Execution.End)
}

func Test_TearDownSuite_CallsCallbackIfDefined(t *testing.T) {
	callbackWasCalled := false
	suiteUnderTest := new(suiteUnderTest)
	suiteUnderTest.OnFinish = func(execution Execution, stats map[string]Stats) {
		callbackWasCalled = true
		assert.Equal(t, suiteUnderTest.Execution, execution)
		assert.Equal(t, suiteUnderTest.Stats, stats)
	}

	suiteUnderTest.SetupSuite()
	suiteUnderTest.TearDownSuite()

	assert.True(t, callbackWasCalled)
}

func Test_BeforeTest_CreatesNewStats(t *testing.T) {
	suiteUnderTest := new(suiteUnderTest)
	suiteUnderTest.SetupSuite()
	suiteUnderTest.BeforeTest("mySuite", "myTest")

	actualStats := suiteUnderTest.Stats["myTest"]

	assert.Equal(t, "mySuite", actualStats.SuiteName)
	assert.Equal(t, "myTest", actualStats.Testname)
	assert.NotZero(t, actualStats.Start)
	assert.Zero(t, actualStats.End)
}

func Test_AfterTest_UpdatesStats(t *testing.T) {
	suiteUnderTest := new(suiteUnderTest)
	suiteUnderTest.SetupSuite()
	suiteUnderTest.BeforeTest("mySuite", "myTest")
	suiteUnderTest.AfterTest("mySuite", "myTest", true)

	actualStats := suiteUnderTest.Stats["myTest"]

	assert.NotZero(t, actualStats.End)
	assert.True(t, actualStats.Success)
}

func Test_RunningTheSuiteWillGenerateStats(t *testing.T) {
	suiteToTest := new(suiteUnderTest)

	suite.Run(t, suiteToTest)

	assert.NotZero(t, suiteToTest.Execution.Start)
	assert.NotZero(t, suiteToTest.Execution.End)
	assert.NotNil(t, suiteToTest.Stats["Test_Something"])
}
