package main

import (
	"fmt"

	"dice/dice"
)

func main() {
	nds := &dice.NumericDiceSet{}
	fmt.Println(dice.LoadAndThrow(nds, 200, 3))

	fds := &dice.FudgeDiceSet{}
	fmt.Println(dice.LoadAndThrow(fds, 3, 10))
}
