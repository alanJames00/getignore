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
	overwrite := flag.Bool("ow", false, "If true, overwrite existing gitignore, otherwise append to current gitignore")
	help := flag.Bool("help", false, "show the help message")
	shortHelp := flag.Bool("h", false, "show the help message")
	allLang := flag.Bool("all", false, "print all available programming languages and platforms")

	flag.Parse()

	switch {
	case *help || *shortHelp:
		internal.PrintHelp()
		os.Exit(0)

	case *allLang:
		internal.ShowAllLangs()
		os.Exit(0)

	case *lang == "":
		fmt.Println("ERROR: missing required --lang flag")
		internal.PrintHelp()
		os.Exit(1)

	default:
		fmt.Println("overwriting:", *overwrite)

		// check if it's a git repo root
		if !internal.IsGitRoot("./") {
			fmt.Println("ERROR: Are you in the Git repository root? Please navigate to the root directory and try again.")
			os.Exit(1)
		}

		// get the template content
		templateContent := internal.GetTemplate(*lang)

		// check if gitignore file exists
		gitignoreExists := internal.IgnoreExists("./")

		// write or append to gitignore file
		internal.WriteGitIgnore(templateContent, *overwrite, gitignoreExists)
	}
}
