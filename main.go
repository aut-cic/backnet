package main

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/aut-cic/backnet/internal/config"
	"github.com/aut-cic/backnet/internal/db"
	"github.com/aut-cic/backnet/internal/http/handler"
	"github.com/aut-cic/backnet/internal/store/conference"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

//go:embed web
var assets embed.FS

// nolint: gochecknoglobals
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

// nolint: funlen
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
		Templates: template.Must(template.ParseFS(assets, "web/template/*.html")),
	}
	pterm.Info.Println(t.Templates.DefinedTemplates())
	app.Renderer = t
	app.Debug = cfg.Debug

	fsys, err := fs.Sub(assets, "web/app")
	if err != nil {
		pterm.Fatal.Printfln("cannot find web/app folder %s", err)
	}

	distHandler := http.FileServer(http.FS(fsys))
	app.GET("/dist/*", echo.WrapHandler(distHandler))

	{
		h := handler.Conference{
			Store: conference.NewSQL(db),
		}

		h.Register(app.Group("/api/conference", middleware.BasicAuth(
			func(username string, password string, _ echo.Context) (bool, error) {
				if username == cfg.Auth.Username && password == cfg.Auth.Password {
					return true, nil
				}

				return false, nil
			})))
	}

	{
		h := handler.Pages{
			Version: version,
		}

		h.Register(app.Group(""))
	}

	if err := app.Start(":1378"); err != nil {
		pterm.Fatal.Printfln("http server initiation failed %s", err)
	}
}
