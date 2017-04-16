package fibgo

var cache = make(map[int]int)

// TODO create fibonacci interface to enable, and cache enable should through this interface

// N return fibonacci number on N position
// N should start from 0, otherwise panic will raised
func N(n int) int {
	if n < 0 {
		panic("n should not less than 0")
	}

	v, ok := cache[n]
	if ok {
		return v
	}

	if n < 2 {
		cache[n] = n
		return n
	}

	v = N(n-2) + N(n-1)
	cache[n] = v
	return v
}

// Seq will generate the fibonacci sequence
func Seq(length int) []int {
	out := make([]int, length)
	for i := 0; i < length; i++ {
		out[i] = N(i)
	}
	return out
}
