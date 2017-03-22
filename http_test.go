package fibgo_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	fib "github.com/uudashr/fibgo"
)

func TestNumbers(t *testing.T) {
	cases := []struct {
		limit  string
		code   int
		result []int
	}{
		{limit: "", code: http.StatusOK, result: []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}},
		{limit: "-1", code: http.StatusBadRequest, result: []int{}},
		{limit: "non-number", code: http.StatusBadRequest, result: []int{}},
		{limit: "0", code: http.StatusOK, result: []int{}},
		{limit: "1", code: http.StatusOK, result: []int{0}},
		{limit: "2", code: http.StatusOK, result: []int{0, 1}},
		{limit: "4", code: http.StatusOK, result: []int{0, 1, 1, 2}},
		{limit: "9", code: http.StatusOK, result: []int{0, 1, 1, 2, 3, 5, 8, 13, 21}},
		{limit: "10", code: http.StatusOK, result: []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}},
	}

	handler := fib.NewHTTPHandler()

	for i, c := range cases {
		limit := c.limit
		expect := c.result
		code := c.code

		req, err := http.NewRequest(http.MethodGet, "/numbers", nil)
		if err != nil {
			t.Fatal("err:", err, "case#:", i)
		}

		if limit != "" {
			q := req.URL.Query()
			q.Add("limit", limit)
			req.URL.RawQuery = q.Encode()
		}

		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		if got, want := rec.Code, code; got != want {
			t.Fatal("got:", got, "want:", want, "limit:", limit)
		}

		if rec.Code != http.StatusOK {
			continue
		}

		var result []int
		if err := json.NewDecoder(rec.Body).Decode(&result); err != nil {
			t.Fatal("err:", err, "limit:", limit)
		}

		if got, want := len(result), len(expect); got != want {
			t.Fatal("len(result):", got, "len(expect):", want, "limit:", limit)
		}

		for ri, got := range result {
			if want := expect[ri]; got != want {
				t.Fatal("result[ri]:", got, "expect[ri]:", want, "ri:", i, "limit:", limit)
			}
		}
	}
}

func TestNumbers_N(t *testing.T) {
	nums := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}

	handler := fib.NewHTTPHandler()
	for i, v := range nums {

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/numbers/%d", i), nil)
		if err != nil {
			t.Fatal("err:", err, "i:", i)
		}

		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		if got, want := rec.Code, http.StatusOK; got != want {
			t.Fatal("got:", got, "want:", want, "i:", i)
		}

		var result int
		if err := json.NewDecoder(rec.Body).Decode(&result); err != nil {
			t.Fatal("err:", err, "i:", i)
		}

		if got, want := result, v; got != want {
			t.Fatal("got:", got, "want:", want, "i:", i)
		}
	}
}

func TestNumbers_nonNumberN_error(t *testing.T) {
	params := []string{"asdf", "872y34h", "2.4"}

	handler := fib.NewHTTPHandler()
	for i, p := range params {

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/numbers/%s", p), nil)
		if err != nil {
			t.Fatal("err:", err, "case#:", i)
		}

		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		if got, want := rec.Code, http.StatusBadRequest; got != want {
			t.Fatal("got:", got, "want:", want, "case#:", i)
		}
	}
}

func TestNumbers_negativeN_error(t *testing.T) {
	params := []string{"-2", "-10", "-1"}

	handler := fib.NewHTTPHandler()
	for i, p := range params {

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/numbers/%s", p), nil)
		if err != nil {
			t.Fatal("err:", err, "case#:", i)
		}

		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		if got, want := rec.Code, http.StatusBadRequest; got != want {
			t.Fatal("got:", got, "want:", want, "case#:", i)
		}
	}
}
