package gitler

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Git Clone and Counter
func GitCloneAndCounter(gitRepo *string, grep *string, gitlerExportFile *string) string {
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
	var i int
	i = 0
	for x := 0; x < len(lineBytes); x++ {
		tag, tagErr := ParseTag(".tmp/"+dirName, string(lineBytes[x]), *grep)
		if tagErr == nil {
			// No error
			tags[x] = tag
			i++
		}

	}
	f, err := os.OpenFile(*gitlerExportFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(dirName + ", " + strconv.Itoa(i) + "\n")

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("Number of tags: ", i)

	//fmt.Println("out:", myString)

	return ""
}
