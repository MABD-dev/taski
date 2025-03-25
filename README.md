# Taski 
A simple task manager for the terminal. Manage your to-dos with ease, right from the terminal


## Features
- Add, Update, and Delete Tasks
- List Tasks: show tasks in kanban style table.
  - search tasks based on keyword(s) and highlight in kanban table/board
- Data is saved locally on your machine in a json file


## Installation
1.  **Clone the repository:**
    ```bash
    git clone https://github.com/MABD-dev/taski
    cd taski
    ```
2.  **Ensure Go is installed:**
    * Make sure Go `1.24.1` or later is installed on your machine.
    * You can check using:
        ```bash
        go version
        ``` 
3.  **Build the project:**
    ```bash
    go build
    ```
    * This will generate the `taski` executable.

4.  **Add to your PATH (Linux/macOS):**
    * Move the `taski` executable to a directory in your `PATH` (e.g., `~/bin` or `/usr/local/bin`).
    * Alternatively, add the directory containing the executable to your `PATH` environment variable.
5.  **Data Location:**
    * Your data is saved locally at `~/.taski/tasks.json`.


## Usage
1. **Listing Tasks**
    ```sh
    taski list # get list of all tasks 
    taski list -s "bug" # show all tasks and highlight 'bug' keyword
    ```
2. **Add Task**
    ```sh
    taski add "update taski readme file" # this will create new task with name "update taski reamde file"
    taski add "task name" -d "this is task description" # add task description message
    taski add "task name" -s inprogress # set task status 
    ```

3. **Update Task**
    ```sh
    taski update <task number> 
    ```
    This will open preferred editor in your terminal (default to vi) to be able to edit task data.
    Close editor to save new changes

4. **Delete Task**
    ```sh
    taski delete <task number>

    # or delete multiple tasks together
    taski delete <task number 1> <task number 2> <task number 3>
    ```

## Roadmap
Features coming up 
- [x] Create/Update/Delete/List tasks
- [x] Assign task to Project
- [ ] REPL Support
- [ ] Tags Support
- [ ] Priorities Support
- [ ] Task due date support


Started from this project: https://roadmap.sh/projects/task-tracker
