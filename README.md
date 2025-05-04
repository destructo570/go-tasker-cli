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


# Usage

# Adding a new task
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
task-cli update 1 "Buy groceries and cook dinner"
task-cli delete 1

# Marking a task as in progress or done
task-cli mark-in-progress 1
task-cli mark-done 1

# Listing all tasks
task-cli list

# Listing tasks by status
task-cli list done
task-cli list todo
task-cli list in-progress
