# GoLang Practice

Este repositório contém projetos e práticas feitas com GoLang.

A intenção é praticar e desenvolver conhecimento sobre a linguagem, toda prática estará nesses repositório, projetos mais complexos estarão em repositórios separados.

Todo Projeto, mesmo que simples, será descrito nesse README.

## Pasta: API Development - TodoList API

A **TodoList API** é uma API simples criada em GoLang utilizando o framework Gin. Ela permite gerenciar uma lista de tarefas.

### Funcionalidades

- **GET /tasks**: Lista todas as tarefas.
- **POST /tasks**: Adiciona uma nova tarefa.

### Como Rodar o Projeto

1. Clone o repositório:

   ```bash
   git clone https://github.com/luizscolari/GoLang-Practice.git

2. Navegue até a pasta TodoList API
    ```bash
    cd API-Development/TodoList-API
    ```
3. Execute o servidor:
    ```bash
    go run main.go
    ```
4. Para adicionar uma tarefa (POST):
   ```bash
   curl http://localhost:5050/tasks \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": 1, "task": "Learn Gin Framework", "description": "Study the Gin web framework", "completed": false}'
   ```
