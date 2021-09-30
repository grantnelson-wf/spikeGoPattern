// +build !prod

package tester

import (
	"testing"

	"github.com/grantnelson-wf/spikeGoPattern/trial1/front"
)

type Tester struct {
	tb   testing.TB
	real front.Front
}

var _ front.Front = (*Tester)(nil)

func (t *Tester) Input(i int) error {
	t.tb.Logf("calling Input(%d)", i)
	err := t.real.Input(i)

	if err != nil {
		t.tb.Errorf("Input(%d) failed: %s", i, err.Error())
		t.tb.Fail()
		return err
	}

	t.tb.Logf("Input(%d) succeeded", i)
	return nil
}

func (t *Tester) Output() (int, error) {
	t.tb.Logf("calling Output()")
	val, err := t.real.Output()

	if err != nil {
		t.tb.Errorf("Output() failed: %s", err.Error())
		if val != -1 {
			t.tb.Errorf("Output() should return -1 on error but had %d", val)
			t.tb.Fail()
		}
		return val, err
	}

	if val < 0 {
		t.tb.Errorf("Output() should return a positive value but got %d", val)
		t.tb.Fail()
		return val, nil
	}

	t.tb.Logf("Output() succeeded")
	return val, nil
}

func (t *Tester) OutputExpects(exp int) error {
	val, err := t.Output()
	if val != exp {
		t.tb.Errorf(`Output() was expected to return %d but it returned %d`, exp, val)
		t.tb.Fail()
	}

	return err
}
