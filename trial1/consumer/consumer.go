package main

import (
	"fmt"

	"github.com/grantnelson-wf/spikeGoPattern/trial1/back"
	"github.com/grantnelson-wf/spikeGoPattern/trial1/front"
)

func main() {
	back.Prepare()

	tim := front.New(`Tim`)

	if err := tim.Input(18); err != nil {
		panic(err)
	}

	if val, err := tim.Output(); err != nil {
		panic(err)
	} else {
		fmt.Println("Tim has the value of", val)
	}
}
