package main

import (
	"fmt"
	"time"

	"github.com/bastiao/contributions/config"
	"github.com/bastiao/contributions/sourceCode"
	"github.com/uber/gonduit"
	"github.com/uber/gonduit/core"
)

// Manage the options for differential / arcanist requests

func ShowArc(confFile *config.ConfGoPath, arcList *bool, arcWatch *bool, arcParams *string) {
	fmt.Println("\n ⭐ Starting pha-go with arc command.")
	fmt.Println("\t List: ", *arcList)
	fmt.Println("\t Watch: ", *arcWatch)
	fmt.Println("\t Params: ", *arcParams)

	fmt.Println("\n🚒 Looking for the contributions for today. ")

	fmt.Println("📃 Endpoint: ", confFile.PhaConf.Endpoint)
	fmt.Println("⌛ Token: ", confFile.PhaConf.Token)

	if *arcWatch {
		watchArcMethod(confFile, arcList, arcWatch, arcParams)
	} else {
		showArcMethod(confFile, arcList, arcWatch, arcParams)
		fmt.Println("\n🚒 Done. ")
	}
}

func showArcMethod(confFile *config.ConfGoPath, arcList *bool, arcWatch *bool, arcParams *string) {

	client, err := gonduit.Dial(confFile.PhaConf.Endpoint,
		&core.ClientOptions{APIToken: confFile.PhaConf.Token})
	_ = err
	client.Connect()

	lstDiff := sourceCode.LookForDifferential(client)
	fmt.Println("\n\n🎆 Open or pending differentials:")

	for _, v := range lstDiff {

		fmt.Println("\t🐊 URI: ", v.URI)
		fmt.Println("\tBranch: ", v.Branch)
		fmt.Println("\tStatusName: ", v.StatusName)
		response := sourceCode.GetSourceCode(client, v.RepositoryPHID)
		fmt.Println("\tRepo: ", response[v.RepositoryPHID].FullName)
		fmt.Println()
		fmt.Println()
	}
}

func watchArcMethod(confFile *config.ConfGoPath,
	arcList *bool,
	arcWatch *bool,
	arcParams *string) {
	fmt.Println("\n🚒 Watching. ")

	retry := 30
	for retry > 0 {
		showArcMethod(confFile, arcList, arcWatch, arcParams)
		time.Sleep(60 * time.Second)
		retry--
	}

}
