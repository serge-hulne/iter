package iter

type Pair struct {
	Index int
	Value int
}

func Fib(n int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i, j := 0, 1; i < n; i, j = i+j, i {
			out <- i
		}
	}()
	return out
}

func Filter(in chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range in {
			if i%2 == 0 {
				out <- i
			}
		}
	}()
	return out
}

func Map(in chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range in {
			out <- i * i
		}
	}()
	return out
}

func Reduce(in chan int) int {
	sum := 0
	for i := range in {
		sum += i
	}
	return sum
}

func Take(in chan int, nmax int) chan int {
	out := make(chan int)
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

func Enumerate(in chan int) chan Pair {
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
