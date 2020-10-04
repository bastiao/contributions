package main

import (
	"fmt"

	"github.com/bastiao/contributions/config"
	"github.com/bastiao/contributions/jenkins"
)

func JenkinsAction(confFile *config.ConfGoPath,
	jenkinsBranch *string,
	jenkinsRepo *string,
	jenkinsParams *string,
	revision *string) {
	fmt.Println("\n ⭐ Starting pha-go with jenkins command.")
	fmt.Println("\t List: ", *jenkinsBranch)
	fmt.Println("\t Watch: ", *jenkinsRepo)
	fmt.Println("\t Params: ", *jenkinsParams)
	fmt.Println("\t Revision: ", *revision)

	jenkins.StartPipeline(confFile, *jenkinsBranch, *jenkinsRepo, *jenkinsParams, *revision)

}
