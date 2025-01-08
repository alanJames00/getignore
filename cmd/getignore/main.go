package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/alanjames00/getignore/internal"
)

func main() {
	// define CLI flags
	lang := flag.String("lang", "", "programming language or platform for the .gitignore template")
	overwrite := flag.Bool("ow", false, "If true, overwrite existing gitignore, otherwise append to cuurret gitignore")
	help := flag.Bool("help", false, "show the help message")
	shortHelp := flag.Bool("h", false, "show the help message")
	allLang := flag.Bool("all", false, "print all available programming languages and platforms")

	flag.Parse()

	// handle help
	if *help || *shortHelp {
		internal.PrintHelp()
		os.Exit(0)
	}

	// handle print all languages
	if *allLang {
		internal.ShowAllLangs()
		os.Exit(0)
	}

	// validate flags
	if *lang == "" {
		fmt.Println("ERROR: missing requireed --lang flag")
		internal.PrintHelp()
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
