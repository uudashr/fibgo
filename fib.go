package fibgo

import "fmt"

// N return fibonacci number on N position
// N should start from 0, otherwise panic will raised
func N(n int) int {
	if n < 0 {
		panic(fmt.Sprintf("n should be greather than 0, but found %d", n))
	}
	if n < 2 {
		return n
	}

	return N(n-2) + N(n-1)
}

// Seq return
func Seq(length int) []int {
	out := make([]int, length)
	for i := 0; i < length; i++ {
		out[i] = N(i)
	}
	return out
}
