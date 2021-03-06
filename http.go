package fibgo

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// NewHTTPHandler create new http.Handler
func NewHTTPHandler() http.Handler {
	e := echo.New()

	e.GET("/numbers", func(c echo.Context) error {
		limitS := c.QueryParam("limit")
		if limitS == "" {
			limitS = "10"
		}

		limit, err := strconv.Atoi(limitS)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		if limit < 0 {
			return c.String(http.StatusBadRequest, "Limit should be not less than 0")
		}

		return c.JSON(http.StatusOK, Seq(limit))
	})

	e.GET("/numbers/:n", func(c echo.Context) error {
		n, err := strconv.Atoi(c.Param("n"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid n (should number)")
		}
		if n < 0 {
			return c.String(http.StatusBadRequest, "Invalid n (should be non-negative)")
		}
		return c.JSON(http.StatusOK, N(n))
	})

	return e
}
