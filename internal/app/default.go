package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dustin-ward/termtyping/internal/statusbar"
)

func defaultHandler(m AppModel, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c", "esc":
			m.quitting = true
			cmds = append(cmds, tea.ClearScreen)
			cmds = append(cmds, tea.Quit)

		case "enter":
			m.CurState = StateTyping
			m.status_bar.CurState = statusbar.StateTyping
		}
	}

	return m, tea.Batch(cmds...)
}
