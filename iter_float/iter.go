package iter

type Pair struct {
	Index int
	Value float64
}

type Chan chan float64

type FilterCallback func(Chan, Chan) Chan

type ReduceCallback func(Chan) float64

//
func Map(in Chan, cb FilterCallback) Chan {
	out := make(Chan)
	go func() {
		defer close(out)
		out = cb(in, out)
	}()
	return out
}

var Filter = Map

//
func Reduce(in Chan, cb ReduceCallback) float64 {
	word := cb(in)
	return word
}

//
func Take(in Chan, nmax int) Chan {
	out := make(Chan)
	go func() {
		defer close(out)
		index := 0
		for i := range in {
			index++
			if index <= nmax {
				out <- i
			}
		}
	}()
	return out
}

//
func Enumerate(in Chan) chan Pair {
	out := make(chan Pair)
	go func() {
		defer close(out)
		index := 0
		for i := range in {
			out <- Pair{index, i}
			index++
		}
	}()
	return out
}

//
func Range(nmax int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for index := 0; index < nmax; index++ {
			out <- index
		}
	}()
	return out
}

// func Slice(int n, int m) {
// }
