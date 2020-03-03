package example

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

func Test_RunSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

type TestSuite struct {
	suite.Suite
}

func (s *TestSuite) HandleStats(suiteName string, stats *suite.SuiteStats) {
	// In this example, we're creating a log file with the execution results.
	filename := fmt.Sprintf("test_%d.log", time.Now().Unix())
	file, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	fmt.Fprintf(file, "Suite: %s\n", suiteName)
	fmt.Fprintf(file, "Total time: %s\n\n", stats.EndTime.Sub(stats.StartTime).String())

	for testName, testStats := range stats.TestStats {
		duration := testStats.EndTime.Sub(testStats.StartTime)
		fmt.Fprintf(file, "%s: %s ==> %v\n", testName, duration.String(), testStats.Passed)
	}
}

func (s *TestSuite) TestThatTakes100ms() {
	time.Sleep(time.Millisecond * 100)
	s.Equal(1, 1)
}

func (s *TestSuite) TestThatTakes400ms() {
	time.Sleep(time.Millisecond * 400)
	s.Equal(1, 1)
}

func (s *TestSuite) TestThatIsQuick() {
	s.Equal(1, 1)
}

func (s *TestSuite) TestThatFails() {
	s.Equal(1, 2)
}
