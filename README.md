# Learning Golang

Learning fundamentals of the Google's Go programming language.

## Getting Started

### Installing Go Using Homebrew On Mac

    brew install go

### Setting up VS Code

### Go CLI commands

1. **go build**
    Compiles Go source code files.
2. **go run**
    Compiles and executes Go source code files.
3. **go fmt**
    Formats all the code in the Go source files in the current directory.
4. **go install**
    Compiles and installs a package.
5. **go get**
    Downloads the raw Go source code from someone else's package.
6. **go test**
    Runs any tests associated with the current project.

## Basics Of Go

### Basic Data Types

- bool
- string
- integer
- float

### Other Data Types

- **Array**
  - A fixed length container for holding the values of a singular data type.
    - An example: Array of strings
    - Declaration:

            names := [20]string

- **Slice**
  - A variable length container for holding the values of a singular data type.
    - An example: Slice of strings
    - Declaration:

            names := [] string{"Amitesh Rai", "Go"}

- **Map**
  - A container for holding the key-value pairs.

- **Struct**
  - A container for holding the key-value pairs.

### Types Of Packages

- **Executable**
  - Defines a package that can be compiled and executed. It must have a function called ***main***.
- **Reusable**
  - Defines a package that can be used as a dependency(helper code).

### Receiver Functions

### Resources For Learning

1. [Go By Example](https://gobyexample.com/)
2. [Exercism's Go Track](https://exercism.org/tracks/go)
3. [Go: The Complete Developer's Guide](https://www.udemy.com/course/go-the-complete-developers-guide/)
