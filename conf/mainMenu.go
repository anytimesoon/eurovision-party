package conf

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type MainMenu struct {
	choice   string
	quitting bool
	list     list.Model
}

func NewMainMenu() MainMenu {
	return MainMenu{}
}

func (m MainMenu) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m MainMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = string(i)
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m MainMenu) View() string {
	if m.choice != "" {
		return quitTextStyle.Render(fmt.Sprintf("%s? Sounds good to me.", m.choice))
	}
	if m.quitting {
		return quitTextStyle.Render("Not hungry? That’s cool.")
	}
	return "\n" + m.list.View()
}
