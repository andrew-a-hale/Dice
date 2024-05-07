package dice

import (
	"fmt"
	"math/rand"
)

type NumericDiceSet struct {
	dice []NumericDie
}

func (ds *NumericDiceSet) Add(faces any) {
	ds.dice = append(ds.dice, MakeNumericDie(faces.(int)))
}

func (ds NumericDiceSet) Throw() any {
	var total int
	for _, d := range ds.dice {
		v := d.Roll()
		fmt.Printf("Rolled: %d\n", v)
		total += v
	}

	return total
}

type NumericDie struct {
	faces         []int
	numberOfFaces int
}

func MakeNumericDie(n int) NumericDie {
	var faces []int
	for i := 0; i < n; i++ {
		faces = append(faces, i)
	}
	return NumericDie{faces: faces, numberOfFaces: n}
}

func (d *NumericDie) Roll() int {
	rng := rand.Int31n(int32(d.numberOfFaces))
	return d.faces[rng]
}
