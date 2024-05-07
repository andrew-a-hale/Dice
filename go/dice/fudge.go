package dice

import (
	"fmt"
	"math/rand"
)

type FudgeDiceSet struct {
	dice []FudgeDie
}

func (ds *FudgeDiceSet) Add(d any) {
	if d.(int)%3 != 0 {
		panic("Error: Fudge Die must have a number of faces that is divisible by 3")
	}
	ds.dice = append(ds.dice, MakeFudgeDie(d.(int)))
}

func (ds FudgeDiceSet) Throw() any {
	var total FudgeFace
	for _, d := range ds.dice {
		v := d.Roll()
		fmt.Printf("Rolled: %d\n", v)
		total += v
	}

	return total
}

type FudgeFace int

const (
	PLUS    FudgeFace = 1
	MINUS   FudgeFace = -1
	NOTHING FudgeFace = 0
)

type FudgeDie struct {
	faces         []FudgeFace
	numberOfFaces int
}

func MakeFudgeDie(nfaces int) FudgeDie {
	var faces []FudgeFace
	for i := 0; i < nfaces; i += 3 {
		faces = append(faces, []FudgeFace{PLUS, NOTHING, MINUS}...)
	}
	return FudgeDie{faces: faces, numberOfFaces: len(faces)}
}

func (d *FudgeDie) Roll() FudgeFace {
	rng := rand.Int31n(int32(d.numberOfFaces))
	return d.faces[rng]
}
