package main

import (
	"embed"
	"io/fs"
	"net/http"
	"todo/__htmgo"

	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
)

//go:embed assets/dist/*
var StaticAssets embed.FS

func main() {
	locator := service.NewLocator()

	h.Start(h.AppOpts{
		ServiceLocator: locator,
		LiveReload:     true,
		Register: func(app *h.App) {
			sub, err := fs.Sub(StaticAssets, "assets/dist")

			if err != nil {
				panic(err)
			}

			http.FileServerFS(sub)

			app.Router.Handle("/public/*", http.StripPrefix("/public", http.FileServerFS(sub)))
			__htmgo.Register(app.Router)
		},
	})
}
