package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ananrafs/descartes-ui/config"
)

func Create() {
	// Serve static files from the frontend/dist directory.
	fs := http.FileServer(http.Dir(config.Configuration.Frontend.Dist))
	http.Handle("/", fs)

	// Start the server.
	fmt.Println("Server listening on port", config.Configuration.Http.Port)
	log.Panic(
		http.ListenAndServe(fmt.Sprintf(":%s", config.Configuration.Http.Port), nil),
	)
}
