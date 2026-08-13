# GoToDo — CLI Todo App in Go

A command-line todo app written in Go. Built with the standard Go project layout, pointer-based state mutation, explicit error handling, unit tests, and an interactive shell mode

## Features
- **Persistence**: saves and loads tasks as JSON (`todos.json`);
- **Interactive mode**: a shell loop (`todo -i`) built on `bufio.Scanner` for running commands without restarting;
- **Error handling**: bounds checking on indices, explicit edge cases, and a `scanner.Err()` check after the input loop;
- **Unit tested**: core package is covered by tests for task mutation, data isolation, and boundary conditions.

## Project Structure
Standard Go project layout:
```text
├── cmd/
│   └── todo/
│       └── main.go       # Entry point: flag parsing & CLI layer
├── internal/
│   └── todo/
│       ├── todo.go       # Business logic: Task/List structs & JSON I/O
│       └── todo_test.go  # Unit tests
├── go.mod
└── README.md
```

### Key Go concepts used
- **Pointer receivers (`*List`)**: methods change the original list instead of a copy;
- **Stringer interface**: `List` implements `String() string`, so it prints directly with `fmt.Println`;
- **File permissions (`0644`)**: owner read/write, group/others read.

## Installation
Requires Go 1.22+. Clone the repo and enter the directory:
```bash
git clone https://github.com/ove4lo/GoToDo
cd GoToDo
```
Sync dependencies:
```bash
go mod tidy
```

## Usage
Run with `go run`, or compile to a binary (see below).

```bash
# Show help
go run cmd/todo/main.go -h

# Add a task (trailing words are joined into one task)
go run cmd/todo/main.go add Buy a fresh pack of green tea

# List tasks with checkboxes
go run cmd/todo/main.go list

# Complete a task by its displayed index
go run cmd/todo/main.go complete 1
```

### Interactive mode
```bash
go run cmd/todo/main.go -i
```
Inside the shell: `add <text>`, `complete <id>`, `help`, `quit`.

## Build & install
Compile to a native binary and install it into your Go bin directory:
```bash
go install ./cmd/todo
```
Go places the binary in `$(go env GOPATH)/bin`. If that directory is in your `PATH`, you can run the app from anywhere:
```bash
todo add "Learn Go"
todo list
todo -i
```

## Testing
```bash
go test -v ./...
```

## Local documentation
View package docs locally with `pkgsite`:
```bash
go install golang.org/x/pkgsite/cmd/pkgsite@latest
pkgsite -open .
```
Then open `http://localhost:8080`