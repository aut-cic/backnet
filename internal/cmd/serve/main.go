package serve

import (
	"html/template"
	"io/fs"
	"net/http"

	"github.com/aut-cic/backnet/internal/config"
	"github.com/aut-cic/backnet/internal/db"
	"github.com/aut-cic/backnet/internal/http/handler"
	"github.com/aut-cic/backnet/internal/store/conference"
	"github.com/aut-cic/backnet/internal/store/group"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pterm/pterm"
)

func Main(assets fs.FS, version string) {
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
			Store:   group.NewSQL(db),
		}

		h.Register(app.Group(""))
	}

	if err := app.Start(":1378"); err != nil {
		pterm.Fatal.Printfln("http server initiation failed %s", err)
	}
}
