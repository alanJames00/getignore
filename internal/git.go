package internal

import (
	"os"
	"path/filepath"
)

// IsGitRoot checks if the given path is a Git repository root.
func IsGitRoot(dir string) bool {
	gitPath := filepath.Join(dir, ".git")

	info, err := os.Stat(gitPath)
	if err != nil || !info.IsDir() {
		return false
	}

	return true
}

// IgnoreExists checks if a .gitignore file exists in the given directory.
func IgnoreExists(dir string) bool {
	gitignorePath := filepath.Join(dir, ".gitignore")

	info, err := os.Stat(gitignorePath)
	if err != nil {
		return false
	}

	return info.Mode().IsRegular()
}
