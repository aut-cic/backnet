package handler

import (
	"bytes"
	"encoding/csv"
	"net/http"

	"github.com/aut-cic/backnet/internal/http/request"
	"github.com/aut-cic/backnet/internal/store/conference"
	"github.com/labstack/echo/v4"
	"github.com/pterm/pterm"
)

type Conference struct {
	Store conference.Conference
}

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

	buff := new(bytes.Buffer)

	wc := csv.NewWriter(buff)

	wc.Write([]string{"username", "password"})
	for _, user := range users {
		wc.Write([]string{
			user.Username,
			user.Value,
		})
	}
	wc.Flush()

	return c.Blob(http.StatusOK, "application/csv", buff.Bytes())
}
