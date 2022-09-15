package main

import (
	"context"

	"github.com/aut-cic/backnet/internal/config"
	"github.com/aut-cic/backnet/internal/db"
	"github.com/aut-cic/backnet/internal/model"
	"github.com/aut-cic/backnet/internal/store/conference"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	pterm.DefaultCenter.Println("in the name of god")

	s, _ := pterm.DefaultBigText.WithLetters(putils.LettersFromString("BackNet")).Srender()
	pterm.DefaultCenter.Println(s)

	pterm.DefaultCenter.
		WithCenterEachLineSeparately().
		Println("The back door to the AUT internet\nDeveloped by @1995parham")

	cfg := config.New()

	db, err := db.New(cfg.Database)
	if err != nil {
		pterm.Fatal.Printfln("database initiation failed %s", err)
	}

	rc := new(model.Check)

	if err := db.WithContext(context.Background()).First(rc).Error; err != nil {
		pterm.Fatal.Printfln("database query failed %s", err)
	}

	pterm.Info.Printfln("%+v\n", rc)

	users, err := conference.NewSQL(db).Create(context.Background(), "parham", 10)
	if err != nil {
		pterm.Fatal.Printfln("conference creation failed %s", err)
	}

	pterm.Info.Printfln("%+v\n", users)
}
