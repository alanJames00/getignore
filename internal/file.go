package internal

import (
	"fmt"
	"os"
)

// write to the gitignore file
func WriteGitIgnore(content string, overwrite bool, gitinoreExists bool) {
	// if file does not exist, create and if file exists and overwrite mode then overwrite
	if overwrite || !gitinoreExists {
		file, err := os.Create(".gitignore")
		if err != nil {
			fmt.Println("ERROR: cannot create or overwrite gitignore file", err.Error())
			os.Exit(1)
		}
		defer file.Close()

		// write content to the file
		_, err = file.WriteString(content)
		if err != nil {
			fmt.Println("ERROR: cannot write to gitignore file", err.Error())
			os.Exit(1)
		}

	} else if !overwrite {
		// open file and append to it
		file, err := os.OpenFile(".gitignore", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("ERROR: cannot write to gitignore file", err.Error())
			os.Exit(1)
		}

		_, err = file.WriteString(content)
		if err != nil {
			fmt.Println("ERROR: cannot write to gitignore file", err.Error())
			os.Exit(1)
		}
	}

	fmt.Println("Successfully written to .gitignore")
}
