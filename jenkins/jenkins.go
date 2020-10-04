package jenkins

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bastiao/phago/config"
	"github.com/bndr/gojenkins"
)

// This module do the integration with
// defined pipeline and start jobs and watch ther status in queue.
// The results will also be available in phabricator.

//func StartPipeline(string endpoint, string username, string token, string pipeline) {
func StartPipeline(confFile *config.ConfGoPath, paramsStr string) {
	jenkins := gojenkins.CreateJenkins(nil, confFile.PhaJenkins.Endpoint,
		confFile.PhaJenkins.Username, confFile.PhaJenkins.Token)
	// Provide CA certificate if server is using self-signed certificate
	// caCert, _ := ioutil.ReadFile("/tmp/ca.crt")
	// jenkins.Requester.CACert = caCert
	jenkins.Init()
	build, err := jenkins.GetJob(confFile.PhaJenkins.Pipeline)
	if err != nil {
		panic("Job Does Not Exist")
	}
	lastSuccessBuild, err := build.GetLastSuccessfulBuild()
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
	}
	fmt.Println("\n🎃 Current build:")
	fmt.Println()
	fmt.Println()

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
