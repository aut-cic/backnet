package handler

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"net/http"

	"github.com/aut-cic/backnet/internal/http/request"
	"github.com/aut-cic/backnet/internal/store/conference"
	"github.com/labstack/echo/v4"
	"github.com/pterm/pterm"
)

type Conference struct {
	Store conference.Conference
}

// nolint: wrapcheck
func (con Conference) Create(c echo.Context) error {
	req := new(request.Conference)

	if err := c.Bind(req); err != nil {
		pterm.Error.Printfln("conference bind request failed %s", err)

		return echo.ErrBadRequest
	}

	users, err := con.Store.Create(c.Request().Context(), req.Name, req.Count, req.Group)
	if err != nil {
		pterm.Error.Printfln("conference creation failed %s", err)

		return echo.ErrInternalServerError
	}

	return c.Render(http.StatusOK, "conference.html", map[string]interface{}{
		"users": users,
		"name":  req.Name,
	})
}

// nolint: wrapcheck
func (con Conference) List(c echo.Context) error {
	name := c.Param("name")

	users, err := con.Store.List(c.Request().Context(), name)
	if err != nil {
		pterm.Error.Printfln("conference creation failed %s", err)

		return echo.ErrInternalServerError
	}

	buff := new(bytes.Buffer)

	wc := csv.NewWriter(buff)

	if err := wc.Write([]string{"username", "password"}); err != nil {
		pterm.Error.Printfln("write to csv failed %s", err)

		return echo.ErrInternalServerError
	}

	for _, user := range users {
		if err := wc.Write([]string{
			user.Username,
			user.Value,
		}); err != nil {
			pterm.Error.Printfln("write to csv failed %s", err)

			return echo.ErrInternalServerError
		}
	}

	wc.Flush()

	c.Response().Header().Add("Content-Disposition", fmt.Sprintf("attachment;filename=%s.csv", name))

	return c.Blob(http.StatusOK, "text/csv", buff.Bytes())
}

func (con Conference) Register(g *echo.Group) {
	g.POST("/", con.Create)
	g.GET("/list/:name", con.List)
}
