package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) (r Ints) {
	for _, v := range i {
		if filter(v) {
			r = append(r, v)
		}
	}
	return r
}

func (i Ints) Discard(filter func(int) bool) (r Ints) {
	for _, v := range i {
		if !filter(v) {
			r = append(r, v)
		}
	}
	return r
}

func (l Lists) Keep(filter func([]int) bool) (r Lists) {
	for _, v := range l {
		if filter(v) {
			r = append(r, v)
		}
	}
	return r
}

func (s Strings) Keep(filter func(string) bool) (r Strings) {
	for _, v := range s {
		if filter(v) {
			r = append(r, v)
		}
	}
	return r
}
