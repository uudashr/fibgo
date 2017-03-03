package fibgo_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
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

		req, err := http.NewRequest(echo.GET, "/numbers", nil)
		if err != nil {
			t.Error("err:", err, "case#:", i)
			t.FailNow()
		}

		if limit != "" {
			params := req.URL.Query()
			params.Add("limit", limit)
			req.URL.RawQuery = params.Encode()
		}

		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		if got, want := rec.Code, code; got != want {
			t.Error("got:", got, "want:", want, "limit:", limit)
			t.FailNow()
		}

		if got, want := rec.Code, code; got != want {
			t.Error("got:", got, "want:", want)
			t.FailNow()
		}

		if rec.Code != http.StatusOK {
			continue
		}

		var result []int
		if err := json.NewDecoder(rec.Body).Decode(&result); err != nil {
			t.Error("err:", err, "limit:", limit)
			t.FailNow()
		}

		if got, want := len(result), len(expect); got != want {
			t.Error("len(result):", got, "len(expect):", want, "limit:", limit)
			t.FailNow()
		}

		for ri, got := range result {
			if want := expect[ri]; got != want {
				t.Error("result[ri]:", got, "expect[ri]:", want, "ri:", i, "limit:", limit)
				t.FailNow()
			}
		}
	}
}

func TestNumbers_N(t *testing.T) {
	nums := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}

	handler := fib.NewHTTPHandler()
	for i, v := range nums {

		req, err := http.NewRequest(echo.GET, fmt.Sprintf("/numbers/%d", i), nil)
		if err != nil {
			t.Error("err:", err, "i:", i)
			t.FailNow()
		}

		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		if got, want := rec.Code, http.StatusOK; got != want {
			t.Error("got:", got, "want:", want, "i:", i)
			t.FailNow()
		}

		var result int
		if err := json.NewDecoder(rec.Body).Decode(&result); err != nil {
			t.Error("err:", err, "i:", i)
			t.FailNow()
		}

		if got, want := result, v; got != want {
			t.Error("got:", got, "want:", want, "i:", i)
			t.FailNow()
		}
	}
}

func TestNumbers_nonNumberN_error(t *testing.T) {
	params := []string{"asdf", "872y34h", "2.4"}

	handler := fib.NewHTTPHandler()
	for i, p := range params {

		req, err := http.NewRequest(echo.GET, fmt.Sprintf("/numbers/%s", p), nil)
		if err != nil {
			t.Error("err:", err, "case#:", i)
			t.FailNow()
		}

		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		if got, want := rec.Code, http.StatusBadRequest; got != want {
			t.Error("got:", got, "want:", want, "case#:", i)
			t.FailNow()
		}
	}
}

func TestNumbers_negativeN_error(t *testing.T) {
	params := []string{"-2", "-10", "-1"}

	handler := fib.NewHTTPHandler()
	for i, p := range params {

		req, err := http.NewRequest(echo.GET, fmt.Sprintf("/numbers/%s", p), nil)
		if err != nil {
			t.Error("err:", err, "case#:", i)
			t.FailNow()
		}

		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		if got, want := rec.Code, http.StatusBadRequest; got != want {
			t.Error("got:", got, "want:", want, "case#:", i)
			t.FailNow()
		}
	}
}
