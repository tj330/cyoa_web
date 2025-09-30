package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/cyoa_web/internal/app"
	"github.com/cyoa_web/internal/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "server port")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	r := routes.SetupRoutes(app)
	server := &http.Server{
		Addr:        fmt.Sprintf(":%d", port),
		Handler:     r,
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
	}

	app.Logger.Printf("server running on port: %d\n", port)
	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
