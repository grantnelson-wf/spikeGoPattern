package front

type Front interface {
	Input(i int) error
	Output() (int, error)
}
