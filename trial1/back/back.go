package back

import "fmt"

type back struct {
	name  string
	value int
}

func (b *back) Input(i int) error {
	if i%2 == 0 {
		b.value = i
		return nil
	}

	return fmt.Errorf(`%s may not accept an odd value for input: %d`, b.name, i)
}

func (b *back) Output() (int, error) {
	if b.value%3 == 0 {
		return b.value, nil
	}

	return -1, fmt.Errorf(`%s may only output values which are multiples of 3: %d`, b.name, b.value)
}
