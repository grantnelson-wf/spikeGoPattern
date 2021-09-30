package back

import "github.com/grantnelson-wf/spikeGoPattern/trial1/front"

type factory struct{}

var _ front.Factory = (*factory)(nil)

func (f *factory) New(name string) front.Front {
	return &back{
		name:  name,
		value: 0,
	}
}

func Prepare() {
	front.DefaultFactory = &factory{}
}
