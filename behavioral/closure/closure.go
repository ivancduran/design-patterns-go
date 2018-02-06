package closure

func ClosureCounter() func() int {
	i := 0

	return func() int {
		i++
		return i
	}

}
