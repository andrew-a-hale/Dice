package dice

type DiceSet[T any] interface {
	Add(T)
	Throw() T
}

func LoadAndThrow(ds DiceSet[any], faces int, n int) any {
	for i := 0; i < n; i++ {
		ds.Add(faces)
	}
	return ds.Throw()
}
