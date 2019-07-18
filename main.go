package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"regexp"
	"strings"
)

func main() {
	path := getInput()

	home, err := getHomeDir()
	if err != nil {
		fmt.Println(path)
		os.Exit(1)
	}

	path = stripHomeDir(path, home)
	path, prefix := stripLeadingSlash(path)
	fmt.Println(shortenPath(path, prefix))
}

func getInput() (input string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = scanner.Text()
	}
	return
}

func getHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.HomeDir, nil
}

func stripHomeDir(path string, homeDir string) string {
	re := regexp.MustCompile("^"+homeDir)
	if re.MatchString(path) {
		return re.ReplaceAllString(path, "~")
	}
	return path
}

func stripLeadingSlash(path string) (string, string) {
	re := regexp.MustCompile("^/")
	if re.MatchString(path) {
		return re.ReplaceAllString(path, ""), "/"
	}
	return path, ""
}

func shortenPath(path string, prefix string) string {
	parts := strings.Split(path, "/")
	shorts := make([]string, len(parts))

	for i := 0; i < len(parts) - 1; i++ {
		shorts[i] = string(parts[i][0])
	}
	shorts[len(shorts)-1] = parts[len(parts)-1]

	return prefix + strings.Join(shorts, "/")
}
