package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Pages struct {
	Version string
}

// nolint: wrapcheck
func (p Pages) Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]string{
		"version": p.Version,
	})
}

// nolint: wrapcheck
func (Pages) Conference(c echo.Context) error {
	return c.Render(http.StatusOK, "conference.html", nil)
}

func (p Pages) Register(g *echo.Group) {
	g.GET("/", p.Index)
	g.GET("/conference", p.Conference)
}
