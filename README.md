# Taski 
A simple task manager for the terminal. Manage your to-dos with ease, right from the terminal


## Features
- Add Tasks: quickly add tasks with name, description and status.
- Update Tasks: update task name, description or status.
- Delete Tasks: with one command you can delete 1 or more tasks
- List Tasks: show tasks in a nice table showing each task information.
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
    taski list -s todo # list tasks with status = todo
    taski list -s todo -s inprogress # list tasks with both status = todo and in-progress
    # You can mix and match status types
    # all possible task status options are: todo, inprogress, done
    ```
2. **Add Task**
    ```sh
    taski add "update taski readme file" # this will create new task with name "update taski reamde file"
    taski add "task name" -d "this is task description" # add task description message
    taski add "task name" -s inprogress # set task status 
    ```

3. **Update Task**
    ```sh
    taski update <task number> -n "new name" -d "new description" -s <new status>
    # you can mix and match any of these flags. At least one flag must be set
    ```

4. **Delete Task**
    ```sh
    taski delete <task number>

    # or delete multiple tasks together
    taski delete <task number 1> <task number 2> <task number 3>
    ```

## Roadmap
Features coming up 
- [x] Create/Update/Delete/List tasks
- [ ] REPL Support
- [ ] Tags Support
- [ ] Priorities Support
- [ ] Task due date support
- [ ] Project/Context for each task


Started from this project: https://roadmap.sh/projects/task-tracker