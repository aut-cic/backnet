package cli

import (
	"context"
	"fmt"

	"github.com/aut-cic/backnet/internal/config"
	"github.com/aut-cic/backnet/internal/db"
	"github.com/aut-cic/backnet/internal/model"
	"github.com/aut-cic/backnet/internal/store/discount"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/pterm/pterm"
)

// nolint gochecknoglobals
var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type discountModel struct {
	store discount.Discount

	isLoading bool

	table   table.Model
	spinner spinner.Model
}

type discountListMsg struct {
	discounts []model.Discount
}

func (m discountModel) List() tea.Msg {
	discounts, err := m.store.List(context.Background())
	if err != nil {
		pterm.Fatal.Printfln("reading from database failed %s", err)
	}

	return discountListMsg{
		discounts: discounts,
	}
}

func (m discountModel) Init() tea.Cmd {
	return m.List
}

func (m discountModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case discountListMsg:
		rows := make([]table.Row, len(msg.discounts))
		for i, discount := range msg.discounts {
			rows[i] = table.Row{
				discount.Type(),
				discount.Value(),
				fmt.Sprintf("%g", discount.Factor()),
			}
		}

		m.isLoading = false
		m.table.SetRows(rows)
	}

	m.table, cmd = m.table.Update(msg)

	return m, cmd
}

func (m discountModel) View() string {
	if m.isLoading {
		return fmt.Sprintf("\n\n   %s Loading from database...\n\n", m.spinner.View())
	}

	return baseStyle.Render(m.table.View()) + "\n"
}

func Discount() {
	cfg := config.New()

	db, err := db.New(cfg.Database)
	if err != nil {
		pterm.Fatal.Printfln("database initiation failed %s", err)
	}

	// nolint: gomnd
	columns := []table.Column{
		{Title: "Type", Width: 20},
		{Title: "Value", Width: 20},
		{Title: "Factor", Width: 10},
	}

	// nolint: gomnd
	dm := discountModel{
		store:     discount.NewSQL(db),
		isLoading: true,
		spinner:   spinner.New(),
		table: table.New(
			table.WithColumns(columns),
			table.WithFocused(true),
			table.WithHeight(7),
		),
	}

	p := tea.NewProgram(dm)
	if _, err := p.Run(); err != nil {
		pterm.Fatal.Println("could not start program:", err)
	}
}
