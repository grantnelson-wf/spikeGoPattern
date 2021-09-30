// +build !prod

package tester

import (
	"testing"

	"github.com/grantnelson-wf/spikeGoPattern/trial1/front"
)

func New(tb testing.TB, name string) *Tester {
	return NewFactory(tb).New(name).(*Tester)
}

type Factory struct {
	tb testing.TB
}

var _ front.Factory = (*Factory)(nil)

func NewFactory(tb testing.TB) *Factory {
	return &Factory{
		tb: tb,
	}
}

func (f *Factory) New(name string) front.Front {
	return &Tester{
		tb:   f.tb,
		real: front.New(name),
	}
}
