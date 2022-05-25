package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Jungle20m/golang-base/cmd/baseservice/app"
	"github.com/Jungle20m/golang-base/cmd/baseservice/component"
	"github.com/go-chi/chi/v5"
)

func main() {
	// Khởi tạo context cần thiết: database, config, ...
	config, err := component.LoadConfig()
	if err != nil {
		os.Exit(1)
	}

	db, err := component.NewDatabase(config)
	if err != nil {
		os.Exit(1)
	}

	appCtx := component.NewAppContext(db)

	// start service here...
	if err := runService(appCtx, config); err != nil {
		os.Exit(1)
	}
}

func runService(appCtx *component.AppCtx, conf *component.Config) error {
	r := chi.NewMux()

	app.InitRoutes(r, *appCtx)

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%d", conf.App.Host, conf.App.Port),
		Handler: r,
		// Get the default general of http.Server
		// todo: figure out each options and set it up correctly.
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	return server.ListenAndServe()
}
