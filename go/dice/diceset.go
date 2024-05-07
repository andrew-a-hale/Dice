package dice

import "fmt"

type DiceSet[T any, U any] interface {
	Add(T)
	Throw() U
}

type NumericDiceSet struct {
	dice []NumericDie
}

func (ds *NumericDiceSet) Add(d NumericDie) {
	ds.dice = append(ds.dice, d)
}

func (ds NumericDiceSet) Throw() int {
	var total int
	for _, d := range ds.dice {
		v := d.Roll()
		fmt.Printf("Rolled: %d\n", v)
		total += v
	}

	return total
}

type FudgeDiceSet struct {
	dice []FudgeDie
}

func (ds *FudgeDiceSet) Add(d FudgeDie) {
	ds.dice = append(ds.dice, d)
}

func (ds FudgeDiceSet) Throw() FudgeFace {
	var total FudgeFace
	for _, d := range ds.dice {
		v := d.Roll()
		fmt.Printf("Rolled: %d\n", v)
		total += v
	}

	return total
}
