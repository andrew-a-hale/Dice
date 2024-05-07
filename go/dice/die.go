package dice

type Die[T any] interface {
	Roll() T
}

func Roll(d Die[any]) any {
	return d.Roll()
}
