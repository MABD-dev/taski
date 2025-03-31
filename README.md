# Taski
A simple offline first, kanban styled, task manager for your terminal


## Features
- Add, Update, View and Delete Tasks
- List Tasks: show tasks in kanban style table.
  - search tasks based on keyword(s) and highlight in kanban table/board
- Data is saved locally on your machine in a json file

#### Task data
    id
    name
    description (optional)
    status (todo, inprogress, done)
    tags (optional)
    project (optional)


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
<br/>

- **Add Task**
    ```sh
    # create new task with name "update taski reamde file"
    taski add "Upload files" 

    # add task description message
    taski add "Upload files" -d "this is task description" 

    # set task status 
    taski add "Upload files" -s inprogress 

    # assign this task to 'chatgpty' project
    taski add "Upload files" -p chatgpty 

    # add a tag to this task, 
    # NOTE: to add/edit/delete tags check `update task` command
    taski add "Upload files" -t web 
    ```
    You can mix and match these flags while creating new task
<br/>

- **Update Task**
    ```sh
    taski update <task number> 
    ```
    Opens preferred editor in your terminal (default to vi) to edit task data.
    Close editor to save new changes
<br/>

- **View Task details**
    ```sh
    taski view 10
    # format: taski view <tasknumber>...
    ```
<br/>

 - **Bulk set project name to tasks**
    ```sh
    taski project "prject name" 1 2 3 4
    # format: taski project "<project name>" <taskNumber>...
    ```
<br/>

- **Bulk set status to tasks**
    ```sh
    taski status "todo" 1 2 3 4
    # format: taski status "<status>" <taskNumber>...
    ```
<br/>

- **Delete Task**
    ```sh
    taski delete 1 2 3 10
    # format: taski delete <tasknumber>...
    ```


## Roadmap
Features coming up 
- [x] Create/Update/Delete/List tasks
- [x] Assign Task to Project
- [x] Support Task Tags
- [ ] REPL Support
- [ ] Priorities Support
- [ ] Task due date support


Started from this project: https://roadmap.sh/projects/task-tracker
