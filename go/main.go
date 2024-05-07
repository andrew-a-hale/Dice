package main

import (
	"fmt"

	"dice/dice"
)

func main() {
	nds := &dice.NumericDiceSet{}
	nds.Add(dice.MakeNumericDie(20))
	nds.Add(dice.MakeNumericDie(20))
	nds.Add(dice.MakeNumericDie(20))
	fmt.Println(nds.Throw())

	fds := &dice.FudgeDiceSet{}
	fds.Add(dice.MakeFudgeDie())
	fds.Add(dice.MakeFudgeDie())
	fds.Add(dice.MakeFudgeDie())
	fmt.Println(fds.Throw())
}
