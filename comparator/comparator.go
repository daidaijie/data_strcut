package comparator

type Comparator interface {
	CompareTo(c Comparator) int
}

func Convert(i interface{}) Comparator {
	var c Comparator
	switch v := i.(type) {
	case int:
		c = ComparableInt(v)
	case Comparator:
		c = v
	default:
		panic("Unsupported type")
	}
	return c
}
