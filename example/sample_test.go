package example

import (
	"fmt"
	"os"
	"testing"
	"time"

	testifystats "github.com/esdrasbeleza/testify-stats"
	"github.com/stretchr/testify/suite"
)

func Test_RunSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

type TestSuite struct {
	suite.Suite
	testifystats.SuiteWithMetrics
}

func (s *TestSuite) SetupSuite() {
	s.SuiteWithMetrics.SetupSuite()

	s.SuiteWithMetrics.OnFinish = func(e testifystats.Execution, s map[string]testifystats.Stats) {
		// In this example, we're creating a log file with the execution results.

		filename := fmt.Sprintf("test_%d.log", time.Now().Unix())
		file, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer file.Close()

		for testName, testStats := range s {
			duration := testStats.End.Sub(testStats.Start)
			fmt.Fprintf(file, "%s: %s ==> %v\n", testName, duration.String(), testStats.Success)
		}
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
