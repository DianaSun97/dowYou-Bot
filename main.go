package main

import (
	"example.com/mod/config"
	"example.com/mod/internal/app"
)

func main() {
	if err := config.ParseConfig(); err != nil {
		panic(err)
	}
	instanceApp, err := app.New()
	if err != nil {
		panic(err)
	}
	if err := instanceApp.Run(); err != nil {
		panic(err)
	}

	select {}
}
