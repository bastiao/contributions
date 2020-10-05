package jenkins

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bastiao/contributions/config"
	"github.com/bastiao/contributions/sourceCode"
	"github.com/bndr/gojenkins"
	"github.com/uber/gonduit"
	"github.com/uber/gonduit/core"
)

// This module do the integration with
// defined pipeline and start jobs and watch ther status in queue.
// The results will also be available in phabricator.

//func StartPipeline(string endpoint, string username, string token, string pipeline) {
func StartPipeline(confFile *config.ConfGoPath, jenkinsBranch string, jenkinsRepo string,
	paramsStr string, revision int) {
	jenkins := gojenkins.CreateJenkins(nil, confFile.PhaJenkins.Endpoint,
		confFile.PhaJenkins.Username, confFile.PhaJenkins.Token)
	// Provide CA certificate if server is using self-signed certificate
	// caCert, _ := ioutil.ReadFile("/tmp/ca.crt")
	// jenkins.Requester.CACert = caCert

	branchParam := ""
	if revision != 0 {
		client, err := gonduit.Dial(confFile.PhaConf.Endpoint,
			&core.ClientOptions{APIToken: confFile.PhaConf.Token})
		_ = err
		client.Connect()
		diff := sourceCode.LookForDifferentialById(client, revision)
		if diff.Branch != "" {
			branchParam := diff.Branch
			fmt.Println("\n\t⏲ - Branch changed: ", branchParam)
		}

	}
	log.Println("Starting Jenkins..")
	jenkins.Init()
	log.Println("Starting Jenkins, with pipeline ", confFile.PhaJenkins.Pipeline)
	build, err := jenkins.GetJob(confFile.PhaJenkins.Pipeline)
	log.Println("Jenkins Build ", build)
	if err != nil {
		panic("Job Does Not Exist")
	}
	lastSuccessBuild, err := build.GetLastSuccessfulBuild()
	log.Println("Jenkins Build ", lastSuccessBuild)
	if err != nil {
		panic("Last SuccessBuild does not exist")
	}

	fmt.Println("\n🏃 Jenkins mode.")

	fmt.Println("\n🙅 Jenkins Nodes:")
	fmt.Println()

	nodes, _ := jenkins.GetAllNodes()

	for _, node := range nodes {

		// Fetch Node Data
		//node.Poll()
		status, _ := node.IsOnline()
		if status {
			fmt.Println("\t📗 Node is online", node.Raw.DisplayName)
		} else {
			fmt.Println("\t📕 Node is offline", node.Raw.DisplayName)
		}
	}

	fmt.Println("\n🎃 Latest job:")
	fmt.Println()
	fmt.Println()

	duration := lastSuccessBuild.GetDuration()
	fmt.Println("\t - Last Success Build: ", lastSuccessBuild.GetParameters())
	fmt.Println("\t - Duration: ", duration/1000, "seconds")

	paramsDynamic := make(map[string]string)
	sParams := strings.Split(paramsStr, ",")

	for _, s := range sParams {
		sTmp := strings.Split(s, "=")
		paramsDynamic[sTmp[0]] = sTmp[1]
		log.Println(sTmp[0], sTmp[1])
	}
	if jenkinsBranch != "" {
		paramsDynamic[jenkinsBranch] = branchParam
	}

	fmt.Println("\n🎃 Current build:")
	fmt.Println()
	fmt.Println()
	fmt.Println("\t - Params: ", paramsDynamic)

	id, _ := jenkins.BuildJob(confFile.PhaJenkins.Pipeline, paramsDynamic)
	fmt.Println("\t 📕 Jenkins Build Id: ", id)
	task, err := jenkins.GetQueueItem(id)
	if err != nil {
		log.Fatal(err)
	}

	retry := 30
	for retry > 0 {
		if task.Raw.Executable.URL != "" {
			break
		}
		time.Sleep(1 * time.Second)
		task.Poll()
		retry--
	}

	// get the build using the build number
	build2, err := jenkins.GetBuild(confFile.PhaJenkins.Pipeline, task.Raw.Executable.Number)
	fmt.Println("\t - Job: ", build2.Job)
	fmt.Println("\t - Building Number: ", task.Raw.Executable.Number)

	fmt.Println("\t - Params: ", build2.GetParameters())
	fmt.Println("\t - Duration: ", build2.GetDuration()/1000, "seconds")
	fmt.Println("\t - Running: ", build2.IsRunning())
	consoleResponse, _ := build2.GetConsoleOutputFromIndex(1)
	fmt.Println("\t - Output:\n\n\n ", consoleResponse)
}
