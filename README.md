# Taski 
A simple task manager for the terminal. Manage your to-dos with ease, right from the terminal


## Features
- Add, Update, View and Delete Tasks
- List Tasks: show tasks in kanban style table.
  - search tasks based on keyword(s) and highlight in kanban table/board
- Bulk set project name to multiple tasks
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
- **Listing Tasks**
    ```sh
    taski list # get list of all tasks 
    taski list -s "bug" # show all tasks and highlight 'bug' keyword
    ```
- **Add Task**
    ```sh
    taski add "update taski readme file" # this will create new task with name "update taski reamde file"
    taski add "task name" -d "this is task description" # add task description message
    taski add "task name" -s inprogress # set task status 
    ```

- **Update Task**
    ```sh
    taski update <task number> 
    ```
    This will open preferred editor in your terminal (default to vi) to be able to edit task data.
    Close editor to save new changes

- **View Task details**
    ```sh
    taski view 10
    # format: taski view <tasknumber>...
    ```

 - **Bulk set project name to tasks**
    ```sh
    taski project "prject name" 1 2 3 4
    # format: taski project "<project name>" <taskNumber>...
    ```

- **Bulk set status to tasks**
    ```sh
    taski status "todo" 1 2 3 4
    # format: taski status "<status>" <taskNumber>...
    ```


- **Delete Task**
    ```sh
    taski delete 1 2 3 10
    # format: taski delete <tasknumber>...
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
