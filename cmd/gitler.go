package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/bastiao/contributions/config"
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
func GitlerIteration(confFile *config.ConfGoPath, csvFile *string) string {
	// Try to count like this
	//  git log --tags --simplify-by-decoration --pretty="format:%ai %d"|grep tag|grep 2021 | wc -l
	fmt.Println("Entering in counter", *csvFile)
	var reposContent string

	reposContent = ReadFileWithRepos(*csvFile)
	//fmt.Println("{}", reposContent)

	for _, line := range strings.Split(reposContent, "\n") {
		fmt.Println("Repo {}", line)
		GitClone(&line)
	}

	return reposContent
}
func GitClone(gitRepo *string) string {
	if err := ensureDir(".tmp"); err != nil {
		fmt.Println("Directory creation failed with error: " + err.Error())
		os.Exit(1)
	}
	cmd := exec.Command("git", "clone", *gitRepo)
	cmd.Dir = ".tmp"
	_, _ = cmd.Output()
	return ""
}
func ensureDir(dirName string) error {
	err := os.Mkdir(dirName, os.ModeDir)
	if err == nil {
		return nil
	}
	if os.IsExist(err) {
		// check that the existing path is a directory
		info, err := os.Stat(dirName)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return errors.New("path exists but is not a directory")
		}
		return nil
	}
	return err
}
