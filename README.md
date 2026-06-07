# Todo CLI

A simple todo CLI built with Go and [Cobra](https://github.com/spf13/cobra).

## Requirements

- Go 1.26+

## Build

```bash
go build -o todo .
```

## Usage

```bash
todo add [todo description]            # Add a new todo
todo list                              # List todos
todo done [todo ID]                    # Mark a todo as done
todo update [todo ID] [description]    # Update a todo description
todo delete [todo ID]                  # Delete a todo
```

## Quick Start

```bash
# 1) Add a todo
./todo add "Finish Go CLI project" -p high -d 2026-06-10

# 2) List current todos
./todo list -d

# 3) Mark a todo as done (replace 1 with your todo ID)
./todo done 1

# 4) List pending todos only
./todo list -t pending
```

## Command Flags

| Command | Flag | Description |
| --- | --- | --- |
| `add` | `-p, --priority` (default: `medium`) | Set priority: `low`, `medium`, `high` |
| `add` | `-d, --deadline` | Set deadline (`YYYY-MM-DD`) |
| `list` | `-s, --sortby` (default: `updated`) | Sort by: `updated`, `priority`, `deadline` |
| `list` | `-q, --search` | Search in todo description |
| `list` | `-t, --status` | Filter by status: `pending`, `completed` |
| `list` | `-d, --details` | Show detailed output |
| `list` | `-c, --complete` | Include completed todos section |
| `update` | `-p, --priority` | Update priority: `low`, `medium`, `high` |
| `update` | `-d, --deadline` | Update deadline (`YYYY-MM-DD`) |

## Data Storage

Todos are stored in `todos.json` in the same directory as the compiled binary.
