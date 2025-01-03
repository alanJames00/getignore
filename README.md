# getignore

`getignore` is a command-line tool to simplify the management of `.gitignore` files in your Git repositories. With `getignore`, you can easily create or update `.gitignore` files based on programming languages or platforms.

## Features

- Generate `.gitignore` files for any programming language or platform.
- Append or overwrite existing `.gitignore` files.
- Easy-to-use command-line interface.
- Works directly in the root of your Git repositories.

## Installation

1. Clone the repository:

```bash
   git clone https://github.com/yourusername/getignore.git
```

2. Build the project with make:

```bash
    make build
```

3. Use the binary wherever you want

## Usage

### Basic Command Structure

```bash
getignore [OPTIONS]
```

### Options

Option Description

| Option     | Description                                                                           |
| ---------- | ------------------------------------------------------------------------------------- |
| --lang     | Specify the programming language or platform for the .gitignore template (required).  |
| --ow       | overwrite the existing .gitignore file if it exists. If omitted, appends to the file. |
| -h, --help | Show the help menu and exit.                                                          |
| -all       | Print all available programming languages and platforms                               |

### Examples

- Append a Python .gitignore template to the existing .gitignore:

```bash
getignore --lang=python
```

- Overwrite the .gitignore file with a Go template:

```bash
getignore --lang=go --ow
```

Show the help menu:

```bash
getignore --help
```

- **Notes**:
  The tool must be run in the root directory of a Git repository.
  If the --lang flag is missing, the program will exit with an error.
