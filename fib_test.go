package fib_test

import (
	"testing"

	fib "github.com/uudashr/fibgo"
)

func TestFibN(t *testing.T) {
	data := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	for n, v := range data {
		if actual, expect := fib.N(n), v; actual != expect {
			t.Error("actual:", actual, "expect:", expect, "n:", n)
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
	if actualLen, expectedLen := len(actual), len(expect); actualLen != expectedLen {
		t.Error("len(actual):", actualLen, "len(expect):", expectedLen)
	}
	for i := 0; i < len(expect); i++ {
		if actual[i] != expect[i] {
			t.Error("actual[i]:", actual, "expect[i]:", expect, "i:", i)
		}
	}
}
