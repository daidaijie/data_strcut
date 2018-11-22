package comparator

import "strconv"

type ComparableInt int

func (i ComparableInt) String() string {
	return strconv.Itoa(int(i))
}

func (i ComparableInt) CompareTo(c Comparator) int {
	v, ok := c.(ComparableInt)
	if !ok {
		panic("Can't compare to not a ComparableInt comparator")
	}
	return int(i) - int(v)
}
