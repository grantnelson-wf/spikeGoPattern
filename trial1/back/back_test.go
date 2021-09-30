package back

import (
	"testing"

	"github.com/grantnelson-wf/spikeGoPattern/trial1/tester"
)

func TestBackBasics(t *testing.T) {
	tester := tester.New(t, `Bobby`)

	tester.Input(1)
	tester.OutputExpects(0)

	tester.Input(2)
	tester.OutputExpects(-1)

	tester.Input(3)
	tester.OutputExpects(-1)

	tester.Input(6)
	tester.OutputExpects(6)
}
