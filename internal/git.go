package main

import (
	"os"
	"path/filepath"
)

// check if the given path is a git repo
func IsGitRoot(dir string) bool {
	// build the full path to the .git directory
	gitPath := filepath.Join(dir, ".git")

	// check if .git directory exists
	info, err := os.Stat(gitPath)
	if err != nil {
		return false
	}

	// check if .git is a directory
	return info.IsDir()
}


