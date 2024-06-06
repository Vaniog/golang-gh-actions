package app

import (
	"github.com/charmbracelet/bubbletea"
	_ "github.com/charmbracelet/bubbletea"
)

type App struct {
	tea *tea.Program
}

func New() *App {
	return &App{
		tea: tea.NewProgram(
			newModel(),
		),
	}
}

func (a *App) Run() error {
	_, err := a.tea.Run()
	return err
}
