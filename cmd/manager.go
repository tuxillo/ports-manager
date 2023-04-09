package main

import (
	"fmt"
	"quantumachine.net/ports_utils"
)

func main() {
	repo := ports_utils.NewGitRepo(
		"https://github.com/dragonflybsd/mirrorselect",
		"/tmp/mirror")
	repo.Clone()
	fmt.Println("Hello world", repo.Exists())
}
