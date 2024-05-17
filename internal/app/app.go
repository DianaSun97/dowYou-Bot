package app

import (
	"example.com/mod/config"
	"example.com/mod/internal/repository"
	"example.com/mod/internal/service"
	"example.com/mod/pkg/yt"
)

type App struct {
	service *service.Service
}

func New() (*App, error) {
	repo, err := repository.NewRepository(
		yt.NewYoutube(),
	)
	if err != nil {
		return nil, err
	}

	instanceService, err := service.New(config.GetConfig().Token, repo)
	if err != nil {
		return nil, err
	}

	return &App{
		service: instanceService,
	}, nil
}

func (a *App) Run() error {
	return a.service.Run()
}

func (a *App) Stop() error {
	return nil
}
