package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bastiao/contributions/config"
)

// This is the main module and CLI to access all funcionality available
// Global idea is to check differentials changes and send them to CI.

func main() {

	arcCmd := flag.NewFlagSet("arc", flag.ExitOnError)
	arcList := arcCmd.Bool("list", false, "false")
	arcWatch := arcCmd.Bool("watch", false, "false")
	arcParams := arcCmd.String("params-ci", "", "")

	jenkinsCmd := flag.NewFlagSet("jenkins", flag.ExitOnError)
	jenkinsBranch := jenkinsCmd.String("branch", "", "")
	jenkinsRepo := jenkinsCmd.String("repo", "", "")
	jenkinsParams := jenkinsCmd.String("params-ci", "", "")

	if len(os.Args) < 2 {
		fmt.Println("\n🚒 pha-go does not recognize your command. ")
		fmt.Println("It is expecting 'arc', 'jenkins' or 'help' subcommands.")
		fmt.Println("\t arc: allow to verify the pending differentials.")
		fmt.Println("\t jenkins: allow to run the build remotely.")
		os.Exit(1)
	}

	var confFile config.ConfGoPath
	confFile.FromFile("conf/pha.yml")

	switch os.Args[1] {

	case "arc":
		arcCmd.Parse(os.Args[2:])
		ShowArc(&confFile, arcList, arcWatch, arcParams)

	case "jenkins":
		jenkinsCmd.Parse(os.Args[2:])
		JenkinsAction(&confFile, jenkinsBranch, jenkinsRepo, jenkinsParams)

	default:
		fmt.Println("Error: not available.")
		os.Exit(1)
	}

}
