# Task Tracker

A simple CLI based task tracker for tracking what you need to do, what you have done, and what you are currently working on.


## Features

- Add, Update, and Delete tasks
- Mark a task as in progress or done
- List all tasks
- List all tasks that are done
- List all tasks that are not done
- List all tasks that are in progress

## Installation

### Using Go
```bash
go install github.com/TheRedScreen64/task-tracker@latest
```

### From Github Releases
You can download **pre-built binaries** for linux and windows from the [latest release](https://github.com/TheRedScreen64/task-tracker/releases/latest).

## Run Locally

1. Make sure you have Go 1.22.6 or later installed

2. Clone the project

```bash
git clone https://github.com/TheRedScreen64/task-tracker.git
cd task-tracker
```

3. Build the project

```bash
go build -o "bin/"
```

4. Run the CLI

```bash
# Linux
./bin/task-tracker --help

# Windows
./bin/task-tracker.exe --help
```

## Usage

For a list of all available commands
```bash
task-tracker --help
```

### Add Task
```bash
task-tracker add "Buy groceries"
```

### Update Task
```bash
task-tracker update 1 "Cook dinner"
```

### Delete Task
```bash
task-tracker delete 1
```

### Mark as In-Progress
```bash
task-tracker mark-in-progress 1
```

### Mark as Done
```bash
task-tracker mark-done 1
```

### List All Tasks
```bash
task-tracker list
```

### List Tasks with Status
```bash
task-tracker list todo
task-tracker list in-progress
task-tracker list done
```

## Project Reference

https://roadmap.sh/projects/task-tracker
