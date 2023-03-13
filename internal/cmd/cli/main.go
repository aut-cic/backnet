package cli

import (
	"context"
	"fmt"

	"github.com/aut-cic/backnet/internal/config"
	"github.com/aut-cic/backnet/internal/db"
	"github.com/aut-cic/backnet/internal/model"
	"github.com/aut-cic/backnet/internal/store/discount"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pterm/pterm"
)

type discountModel struct {
	discounts []model.Discount
	store     discount.Discount
	cursor    int
}

func (m *discountModel) Init() tea.Cmd {
	discounts, err := m.store.List(context.Background())
	if err != nil {
		pterm.Fatal.Printfln("loading discounts failed %s", err)
	}

	m.discounts = discounts

	return nil
}

func (m *discountModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.discounts)-1 {
				m.cursor++
			}
		}
	}

	return m, nil
}

func (m *discountModel) View() string {
	s := "AUT Internet discounts\n\n"

	for i, discount := range m.discounts {
		cursor := " "
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		s += fmt.Sprintf("%s [ ] %s %s %f\n", cursor, discount.Type(), discount.Value(), discount.Factor())
	}

	s += "\nPress q to quit.\n"

	return s
}

func Discount() {
	cfg := config.New()

	db, err := db.New(cfg.Database)
	if err != nil {
		pterm.Fatal.Printfln("database initiation failed %s", err)
	}

	dm := discountModel{
		store:     discount.NewSQL(db),
		discounts: nil,
		cursor:    0,
	}

	p := tea.NewProgram(&dm)
	if _, err := p.Run(); err != nil {
		pterm.Fatal.Println("could not start program:", err)
	}
}
