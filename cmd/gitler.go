package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bastiao/contributions/config"
	"github.com/bastiao/contributions/gitler"
)

// Read List of Repositories
func ReadFileWithRepos(repos string) string {
	f, err := os.Open(repos)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var content string
	content = ""
	for scanner.Scan() {
		content = content + scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return content
}

// Command Iterations for Counter
func GitlerIteration(confFile *config.ConfGoPath, csvFile *string, grep *string, gitlerExportFile *string) string {
	// Try to count like this
	//  git log --tags --simplify-by-decoration --pretty="format:%ai %d"|grep tag|grep 2021 | wc -l
	fmt.Println("Entering in counter", *csvFile)
	var reposContent string

	reposContent = ReadFileWithRepos(*csvFile)
	//fmt.Println("{}", reposContent)

	for _, line := range strings.Split(reposContent, "\n") {
		gitler.GitCloneAndCounter(&line, grep, gitlerExportFile)
	}

	return reposContent
}
