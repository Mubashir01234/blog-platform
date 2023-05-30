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
+-- config                  ──────────────> load up the configuration from env file
|   +-- config.go
+-- controllers             ──────────────> define your controller layer in which multiple apis present
|   +-- blog.go             ──────────────> defined all the blog apis of the project   
|   +-- module.go           ──────────────> it connect the routes with apis and apis with database
|   +-- user.go             ──────────────> defined all the user apis of the project
|   +-- utils.go            ──────────────> define basic function that is used by whole projects like, hashing, time conversion etc
+-- db                      ──────────────> defined your dataservices layer
|   +-- delete.go           ──────────────> defined all the delete functions that delete data from database like delete user and delete blog etc 
|   +-- module.go           ──────────────> defined the connection of database of your project
|   +-- read.go             ──────────────> defined all the read functions that read the user and blog data from database
|   +-- update.go           ──────────────> defined all the update functions that update the user and blog data from database
|   +-- write.go            ──────────────> defined all the insert functions that insert the user and blog data into database
+-- middleware.go           ──────────────> jwt function written over there, for verify and creating the tokens
|   +-- jwt_generator.go    ──────────────> defined the function that generate the jwt token on login
|   +-- jwt_verify.go       ──────────────> defined the verification function that verify the login token and extract the user details from token
|   +-- response.go         ──────────────> defined the responses of the middleware layer
+-- models                  ──────────────> defined all the structs of the project
|   +-- errors              ──────────────> defined the error responses of your project
|   |   +-- errors.go
|   +-- models.go           ──────────────> defined the main struct of your project like user and blog structs
|   +-- request.go          ──────────────> defined the request struct of your project
|   +-- response.go         ──────────────> defined the response struct of your project
+-- routes
|   +-- routes.go           ──────────────> it contains all the routes of your project
+-- .env                    ──────────────> it contains all the environment variables of your project
+-- .gitignore              ──────────────> it defines all the gitignore variables of your project that we not want to function on GitHub
+-- go.mod                  ──────────────> load up all dependencies required in a project
+-- go.sum                  ──────────────> load up cryptographic hashes of exact version of dependencies required in a project
+-- main.go                 ──────────────> heart of the project, its all start from here
+-- README.md               ──────────────> introduction of the project
```
