package main

import (
	"fmt"

	"github.com/bastiao/contributions/config"
	"github.com/bastiao/contributions/jenkins"
)

func JenkinsAction(confFile *config.ConfGoPath,
	jenkinsBranch *string,
	jenkinsRepo *string,
	jenkinsParams *string) {
	fmt.Println("\n ⭐ Starting pha-go with jenkins command.")
	fmt.Println("\t List: ", *jenkinsBranch)
	fmt.Println("\t Watch: ", *jenkinsRepo)
	fmt.Println("\t Params: ", *jenkinsParams)

	jenkins.StartPipeline(confFile, *jenkinsParams)

}
