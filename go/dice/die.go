package dice

import "math/rand"

type Die[T any] interface {
	Roll() T
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

func MakeFudgeDie() FudgeDie {
	faces := []FudgeFace{PLUS, PLUS, NOTHING, NOTHING, MINUS, MINUS}
	return FudgeDie{faces: faces, numberOfFaces: len(faces)}
}

func (d *FudgeDie) Roll() FudgeFace {
	rng := rand.Int31n(int32(d.numberOfFaces))
	return d.faces[rng]
}
