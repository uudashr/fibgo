package fibgo_test

import (
	"fmt"
	"testing"

	fib "github.com/uudashr/fibgo"
)

func TestN(t *testing.T) {
	data := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	for n, v := range data {
		if got, want := fib.N(n), v; got != want {
			t.Error("got:", got, "want:", want, "n:", n)
		}
	}
}

func TestThatFail(t *testing.T) {
	t.Error("always fail")
}

func ExampleN() {
	fmt.Println(fib.N(0))
	fmt.Println(fib.N(6))
	fmt.Println(fib.N(9))
	// Output:
	// 0
	// 8
	// 34
}

func BenchmarkN(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fib.N(50)
	}
}

func TestN_panic(t *testing.T) {
	fibN := func(n int) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expecting panic")
			}
		}()
		fib.N(n)
	}

	for i := -1; i >= -10; i-- {
		fibN(i)
	}
}

func ExampleSeq() {
	fmt.Println(fib.Seq(10))
	// Output: [0 1 1 2 3 5 8 13 21 34]
}
