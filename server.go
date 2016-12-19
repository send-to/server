package main

import (
	"fmt"

	"github.com/fragmenta/server"

	"github.com/send-to/server/src/app"
)

func main() {

	// Setup server
	server, err := server.New()
	if err != nil {
		fmt.Printf("Error creating server %s", err)
		return
	}

	app.Setup(server)

	// Inform user of server setup try
	server.Logf("#info Starting server in %s mode on port %d", server.Mode(), server.Port())

	if server.Production() {

		// Redirect all :80 traffic to :443 -
		server.StartRedirectAll(80, server.Config("root_url"))

		// If in production, serve over tls with autocerts from let's encrypt
		err = server.StartTLSAutocert(server.Config("autocert_email"), server.Config("autocert_domains"))
		if err != nil {
			server.Fatalf("Error starting server %s", err)
		}

	} else {
		// else just serve on local port without https
		err = server.Start()
		if err != nil {
			server.Fatalf("Error starting server %s", err)
		}

	}

}
