package fibgo_test

import (
	"testing"

	fib "github.com/uudashr/fibgo"
)

func TestFibN(t *testing.T) {
	data := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	for n, v := range data {
		if got, want := fib.N(n), v; got != want {
			t.Error("got:", got, "want:", want, "n:", n)
		}
	}
}

func TestFibN_panic(t *testing.T) {
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

func TestFibSeq(t *testing.T) {
	expect := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	actual := fib.Seq(10)
	if got, want := len(actual), len(expect); got != want {
		t.Error("len(actual):", got, "len(expect):", want)
	}
	for i := 0; i < len(expect); i++ {
		if got, want := actual[i], expect[i]; got != want {
			t.Error("actual[i]:", got, "expect[i]:", want, "i:", i)
		}
	}
}
