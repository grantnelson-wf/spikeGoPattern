package front

func New(name string) Front {
	return DefaultFactory.New(name)
}

type Factory interface {
	New(name string) Front
}

var DefaultFactory Factory
