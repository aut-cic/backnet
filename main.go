package main

import (
	"embed"
	"html/template"

	"github.com/aut-cic/backnet/internal/config"
	"github.com/aut-cic/backnet/internal/db"
	"github.com/aut-cic/backnet/internal/http/handler"
	"github.com/aut-cic/backnet/internal/store/conference"
	"github.com/labstack/echo/v4"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

// nolint: gochecknoglobals
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

//go:embed web/template/*.html
var templates embed.FS

func main() {
	pterm.DefaultCenter.Println("in the name of god")

	s, _ := pterm.DefaultBigText.WithLetters(putils.LettersFromString("BackNet")).Srender()
	pterm.DefaultCenter.Println(s)

	pterm.DefaultCenter.
		WithCenterEachLineSeparately().
		Printfln(
			"backnet %s, commit %s, built at %s\n"+
				"The back door to the AUT internet\nDeveloped by @1995parham",
			version, commit, date,
		)

	cfg := config.New()

	db, err := db.New(cfg.Database)
	if err != nil {
		pterm.Fatal.Printfln("database initiation failed %s", err)
	}

	app := echo.New()

	t := &handler.Template{
		Templates: template.Must(template.ParseFS(templates)),
	}
	app.Renderer = t

	{
		h := handler.Conference{
			Store: conference.NewSQL(db),
		}

		h.Register(app.Group("/api/conference"))
	}

	{
		h := handler.Pages{}

		h.Register(app.Group(""))
	}

	if err := app.Start(":1378"); err != nil {
		pterm.Fatal.Printfln("http server initiation failed %s", err)
	}
}
