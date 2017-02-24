package fibgo

var cache = make(map[int]int)

func noCache(n int) int {
	if n < 0 {
		panic("n should not less than 0")
	}
	if n < 2 {
		return n
	}

	return N(n-2) + N(n-1)
}

// N return fibonacci number on N position
// N should start from 0, otherwise panic will raised
func N(n int) int {
	v, ok := cache[n]
	if ok {
		return v
	}

	v = noCache(n)
	cache[n] = v
	return v
}

// Seq return
func Seq(length int) []int {
	out := make([]int, length)
	for i := 0; i < length; i++ {
		out[i] = N(i)
	}
	return out
}
