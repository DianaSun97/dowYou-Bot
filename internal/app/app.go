package app

import "example.com/mod/internal/service"

type App struct {
	service *service.Service
}

func New() *App {
	return &App{}
}

func (app *App) Run() error {
	return nil
}

func (a *App) Stop() error {
	return nil
}
