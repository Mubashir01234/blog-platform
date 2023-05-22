# Blogging Platform
A Blogging Platform is a web application that allows users to create and manage their blogs. It provides functionalities for user registration, authentication, and blog management. This project implements the backend functionality of the Blogging Platform using Go and MongoDB.

## Features
- User registration: Users can create an account by providing their username, email, and password. User information is stored in the MongoDB database.
- User authentication: Registered users can authenticate themselves using their email and password. JWT (JSON Web Token) is used for authentication and authorization.
- Blog management: Authenticated users can create, update, and delete their blogs. Each blog has a title, description, and associated user.

## Getting Started
1. Clone the repository:
```bash
$ git clone https://github.com/<your-username>/blog-platform.git
$ cd blog-platform
```
2. Install the dependencies:
```bash
$ go mod download
```
3. Configure the application:
Modify the `.env` file and add `Server Port`, `MongoDB` Connection string and `JWT secret`
4. Start the application:
```bash
$ go run main.go
```
The application running on http://localhost:8080.



### Directory Structure

```
+-- config
|   +-- config.go
+-- controllers
|   +-- blog.go
|   +-- module.go
|   +-- user.go
|   +-- utils.go
+-- db
|   +-- delete.go
|   +-- module.go
|   +-- read.go
|   +-- update.go
|   +-- write.go
+-- middleware.go
|   +-- jwt_generator.go
|   +-- jwt_verify.go
|   +-- response.go
+-- models
|   +-- errors
|   |   +-- errors.go
|   +-- models.go
|   +-- request.go
|   +-- response.go
+-- routes
|   +-- routes.go
+-- .env
+-- .gitignore
+-- go.mod
+-- go.sum
+-- main.go
+-- README.md
```
