# GoLang Practice

This repository contains projects and practices done with GoLang.

The intention is to practice and develop knowledge about the language. All practice will be in this repository, and more complex projects will be in separate repositories.

## API Development 

Contains examples and practices focused on building APIs with GoLang, demonstrating key concepts and techniques for API design and implementation.

### TodoList API

This GoLang program implements a simple RESTful API using the Gin framework. It defines a todoList struct to represent tasks with attributes like ID, task name, description, and completion status. The program sets up two endpoints:
-   GET /tasks - Returns a list of all tasks in JSON format.
-   POST /tasks - Allows adding new tasks by accepting a JSON payload and appending it to the task list.

### Dollar Quote API

This GoLang program defines a simple web server that fetches the latest USD to BRL (dollar to Brazilian real) exchange rates from an external API (economia.awesomeapi.com.br). It exposes an endpoint /quote/dollar which, when accessed, returns the current buy and sell rates for the dollar, along with the date of the quote in JSON format.

## Cryptography GoLang

### Crypto-test

This GoLang program demonstrates RSA encryption using OAEP (Optimal Asymmetric Encryption Padding). It generates an RSA private key, encrypts a secret message with the public key using SHA-256 as the hash function, and then prints the resulting ciphertext in hexadecimal format.
- Encrypts a secret message with RSA and OAEP using the public key.
- Uses SHA-256 for hashing in the encryption process.
