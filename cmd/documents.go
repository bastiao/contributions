package main

import (
	"fmt"

	"github.com/bastiao/contributions/config"
	"github.com/bastiao/contributions/documents"
	"github.com/uber/gonduit"
	"github.com/uber/gonduit/core"
)

// Manage the options for differential / arcanist requests

func ShowDocuments(confFile *config.ConfGoPath, documentList *bool) {
	fmt.Println("\n⭐ Starting pha-go with documents command.")
	fmt.Println("\t List: ", *documentList)

	fmt.Println("📃 Endpoint: ", confFile.PhaConf.Endpoint)
	fmt.Println("⌛ Token: ", confFile.PhaConf.Token)

	client, err := gonduit.Dial(confFile.PhaConf.Endpoint,
		&core.ClientOptions{APIToken: confFile.PhaConf.Token})
	_ = err
	client.Connect()

	documents.LookForDocument(client)
}
