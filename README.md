# 🛠️ tasker-cli

A simple and fast command-line tool to manage your tasks and todos right from your terminal.

Built with [Go](https://golang.org) and [Cobra](https://github.com/spf13/cobra).

---

## 📦 Features

- 📋 Add multiple tasks at once
- ✏️ Update task status
- 🗑️ Remove tasks by ID
- 📃 List all, done, or pending tasks
- ⚡ Simple and efficient CLI experience

---

## 🧰 Tech Stack

- [Go](https://golang.org/)
- [Cobra CLI Library](https://github.com/spf13/cobra)

---

## 🚀 Installation

```bash
git clone https://github.com/your-username/tasker-cli.git
cd tasker-cli
go build -o tasker
./tasker --help
```

## Usage

```bash
# Adding a new task
tasker-cli add "Buy groceries"

# Updating and deleting tasks
tasker-cli update 1 "Buy groceries and cook dinner"
tasker-cli remove 1

# Marking a task as in progress or done
tasker-cli mark-pending 1
tasker-cli mark-done 1

# Listing all tasks
tasker-cli list

# Listing tasks by status
tasker-cli list done
tasker-cli list pending
```
