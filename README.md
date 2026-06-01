# Todo CLI

A simple CLI todo application built with Go and [Cobra](https://github.com/spf13/cobra).

## Installation

```bash
go build -o todo .
```

## Usage

```
todo add [description]             Add a new todo
todo list                          List all todos
todo complete [id]                 Mark a todo as completed
todo delete [id]                   Permanently delete a todo
todo update [id] [description]     Update a todo
```

### Flags

| Command    | Flag                        | Description                     |
| ---------- | --------------------------- | ------------------------------- |
| `add`      | `-p, --priority` (medium)   | Priority: low, medium, high     |
| `add`      | `-d, --deadline`            | Deadline (YYYY-MM-DD)           |
| `list`     | `-s, --sortby` (updated)    | Sort: updated, priority, deadline|
| `list`     | `-q, --search`              | Search by description           |
| `list`     | `-t, --status`              | Filter: pending, completed      |
| `update`   | `-p, --priority`            | Update priority                 |
| `update`   | `-d, --deadline`            | Update deadline                 |

## Data

Todos are stored in `todos.json` in the same directory as the binary.

## Build

Requires Go 1.26+.

```bash
go build -o todo .
```
