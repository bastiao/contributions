package gitler

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Git Clone and Counter
func GitCloneAndCounter(gitRepo *string, grep *string) string {
	if err := ensureDir(".tmp"); err != nil {
		fmt.Println("Directory creation failed with error: " + err.Error())
		os.Exit(1)
	}
	var dirName string
	fmt.Println(dirName)
	dirName = strings.Replace(*gitRepo, ".git", "", -1)
	fmt.Println(dirName)
	var dirs []string
	dirs = strings.Split(dirName, "/")
	dirName = dirs[len(dirs)-1]
	fmt.Println(dirName)
	cmd := exec.Command("git", "clone", *gitRepo)
	cmd.Dir = ".tmp"
	_, _ = cmd.Output()

	cmd2 := exec.Command("git", "log", "--tags", "--simplify-by-decoration", "--pretty='format:%ai %d'")
	cmd2.Dir = ".tmp/" + dirName
	output, err := cmd2.Output()
	if err != nil {
		return "nil"
	}
	//myString := string(output)
	lineBytes := bytes.Split(output, []byte{'\n'})
	// The last split is just an empty string, right?
	lineBytes = lineBytes[0 : len(lineBytes)-1]
	tags := make([]*Tag, len(lineBytes))

	for x := 0; x < len(lineBytes); x++ {
		tag, tagErr := ParseTag(".tmp/"+dirName, string(lineBytes[x]), *grep)
		if tagErr != nil {
			return "error"
		}
		tags[x] = tag
	}

	//fmt.Println("out:", myString)

	return ""
}
