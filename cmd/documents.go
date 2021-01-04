package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/bastiao/contributions/config"
	"github.com/bastiao/contributions/documents"
	"github.com/uber/gonduit"
	"github.com/uber/gonduit/core"
)

// Manage the options for differential / arcanist requests

func ShowDocuments(confFile *config.ConfGoPath, documentList *bool, documentQuery *string, documentFilter *string) {
	fmt.Println("\n⭐ Starting pha-go with documents command.")
	fmt.Println("\t List: ", *documentList)

	fmt.Println("📃 Endpoint: ", confFile.PhaConf.Endpoint)
	fmt.Println("⌛ Token: ", confFile.PhaConf.Token)

	client, err := gonduit.Dial(confFile.PhaConf.Endpoint,
		&core.ClientOptions{APIToken: confFile.PhaConf.Token})
	_ = err
	client.Connect()

	docs := documents.LookForDocument(client, documentQuery)
	for _, v := range docs.Data {
		fmt.Println("\t🐊 Phid: ", v.Phid)
		for _, line := range strings.Split(strings.TrimSuffix(v.Attachments.Content.Content.Raw, "\n"), "\n") {
			fmt.Println(line)
			re := regexp.MustCompile(*documentFilter)
			fmt.Println(re.FindStringSubmatch(line))
		}
	}
}
