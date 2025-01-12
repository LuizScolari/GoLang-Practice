# GoLang Practice

This repository contains projects and practices done with GoLang.

The intention is to practice and develop knowledge about the language. All practice will be in this repository, and more complex projects will be in separate repositories.

Every project, even if simple, will be described in this README.

## Folder: API Development - TodoList API

The **TodoList API** is a simple API created in GoLang using the Gin framework. It allows managing a to-do list.

### Features

- **GET /tasks**: Lists all tasks.
- **POST /tasks**: Adds a new task.

### How to Run the Project

1. Clone the repository:

   ```bash
   git clone https://github.com/luizscolari/GoLang-Practice.git

2. Navigate to the TodoList API folder:
    ```bash
    cd API-Development/TodoList-API
    ```
3. Run the server:
    ```bash
    go run main.go
    ```
4. To add a task (POST):
   ```bash
   curl http://localhost:5050/tasks \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": 1, "task": "Learn Gin Framework", "description": "Study the Gin web framework", "completed": false}'
   ```
