package main

import (
	"context"
	"embed"
	"os"

	backnetCli "github.com/aut-cic/backnet/internal/cmd/cli"
	"github.com/aut-cic/backnet/internal/cmd/serve"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"github.com/urfave/cli/v3"
)

// nolint: gochecknoglobals
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

//go:embed web
var assets embed.FS

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

	// nolint: exhaustruct
	app := &cli.Command{
		Name:           "Backnet",
		Usage:          "The backdoor to the AUT Internet",
		Description:    "The backdoor to the AUT Internet",
		DefaultCommand: "serve",
		Version:        version,
		Commands: []*cli.Command{
			{
				Name:        "serve",
				Description: "Run a webserver on :1378",
				Action: func(_ context.Context, _ *cli.Command) error {
					serve.Main(assets, version)

					return nil
				},
			},
			{
				Name:        "cli",
				Description: "Comamnd-line interface",
				Commands: []*cli.Command{
					{
						Name: "discount",
						Action: func(_ context.Context, _ *cli.Command) error {
							backnetCli.Discount()

							return nil
						},
					},
				},
			},
		},
	}

	_ = app.Run(context.Background(), os.Args)
}
