package main

import (
	"flag"
	"fmt"
	"getignore/internal"
	"os"
)

func main() {
	// define CLI flags
	lang := flag.String("lang", "", "programming language or platform for the .gitignore template")
	overwrite := flag.Bool("ow", false, "If true, overwrite existing gitignore, otherwise append to cuurret gitignore")

	flag.Parse()

	// validate flags
	if *lang == "" {
		fmt.Println("ERROR: missing requireed --lang flag")
		// print all available langs
		// TODO: send help
		os.Exit(1)
	}

	fmt.Println("overwriting:", *overwrite)

	// check if its a git repo root
	isGitRoot := internal.IsGitRoot("./")
	if !isGitRoot {
		fmt.Println("ERROR: current working directory is not a git repository root")
		os.Exit(1)
	}

	// get the template content 
	templateContent := internal.GetTemplate(*lang)

	// check if gitignore file exists
	gitignoreExists := internal.IgnoreExists("./")

	// write or append to gitignore file
	internal.WriteGitIgnore(templateContent, *overwrite, gitignoreExists)
}
