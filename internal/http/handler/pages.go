package handler

import (
	"net/http"

	"github.com/aut-cic/backnet/internal/store/group"
	"github.com/labstack/echo/v4"
	"github.com/pterm/pterm"
)

type Pages struct {
	Version string

	Store group.Group
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

// nolint: wrapcheck
func (p Pages) Groups(c echo.Context) error {
	groups, err := p.Store.List(c.Request().Context())
	if err != nil {
		pterm.Error.Printfln("cannot read groups %s", err)
	}

	groupNames := make([]string, len(groups))
	for i, g := range groups {
		groupNames[i] = g.Groupname
	}

	return c.Render(http.StatusOK, "groups.html", map[string]interface{}{
		"groups": groupNames,
	})
}

func (p Pages) Register(g *echo.Group) {
	g.GET("/", p.Index)
	g.GET("/conference", p.Conference)
	g.GET("/groups", p.Groups)
}
