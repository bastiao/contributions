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

// Verify if a string match a specific criterias with filters such as str1|str2
func matchSequenceStr(strEntry *string, filter *string) bool {
	for _, line := range strings.Split(*filter, "|") {
		if strings.Contains(*strEntry, line) {
			return true
		}
	}
	return false

}

// Management options for documents
func ShowDocuments(confFile *config.ConfGoPath, documentList *bool, documentQuery *string,
	documentFilter *string, documentFilterBoolean *string, documentTitle *string, documentShowAll *bool) {
	fmt.Println("\n⭐ Starting pha-go with documents command.")
	fmt.Println("\t Options: ")
	fmt.Println("📃 DocumentList: ", *documentList)
	fmt.Println("⌛ Document Query (--query): ", *documentQuery)
	fmt.Println("📃 documentFilter: (--filter)", *documentFilter)
	fmt.Println("📃 Document Match (--match): ", *documentFilterBoolean)
	fmt.Println("📃 Document Title: (--title)", *documentTitle)
	fmt.Println("📃 Document Show All (--show-all): ", *documentShowAll)

	client, err := gonduit.Dial(confFile.PhaConf.Endpoint,
		&core.ClientOptions{APIToken: confFile.PhaConf.Token})
	_ = err
	client.Connect()

	docs := documents.LookForDocument(client, documentQuery)
	for _, v := range docs.Data {

		reTitle := regexp.MustCompile(*documentTitle)
		matchTitle := reTitle.FindStringSubmatch(v.Attachments.Content.Title)

		if len(matchTitle) > 1 && matchSequenceStr(&matchTitle[1], documentFilterBoolean) {
			fmt.Println("\t🐊 Customer: ", v.Attachments.Content.Title)
			for _, line := range strings.Split(strings.TrimSuffix(v.Attachments.Content.Content.Raw, "\n"), "\n") {
				if *documentShowAll {
					fmt.Println(line)
				}

				re := regexp.MustCompile("^## (.*?)$")
				re2 := regexp.MustCompile("^\\=\\=\\ (.*?) \\=\\=$")

				match := re.FindStringSubmatch(line)

				if len(match) > 1 && strings.Contains(match[1], *documentFilter) == false {
					break

				} else {
					match2 := re2.FindStringSubmatch(line)
					if len(match2) > 1 && strings.Contains(match2[1], *documentFilter) == false {
						break
					}
				}

			}
		}
	}
}
