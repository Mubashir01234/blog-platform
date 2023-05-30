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
+-- config                                  ──────────────> load up the configuration from env file
|   +-- config.go
|       +-- LoadConfig                      ──────────────> LoadConfig will load the configuration
+-- controllers                             ──────────────> define your controller layer in which multiple apis present
|   +-- blog.go                             ──────────────> defined all the blog apis of the project   
|       +-- CreateBlog                      ──────────────> CreateBlog will upload a new blog for Author and then Reader can review the blog
|       +-- UpdateBlog                      ──────────────> in UpdateBlog the author can update own blog
|       +-- GetBlogById                     ──────────────> in GetBlogById user can view the Author blog with the help of blog id
|       +-- DeleteBlog                      ──────────────> DeleteBlog is used to delete a blog author only delete own blog but admin have access to all the authors blog but Reader have no access to delete any blog
|       +-- GetUserAllBlogsByUsername       ──────────────> in GetUserAllBlogsByUsername user can get all the blog of specific user with the help of username
|   +-- module.go                           ──────────────> it connect the routes with apis and apis with database
|   +-- user.go                             ──────────────> defined all the user apis of the project
|       +-- Register                        ──────────────> register is used to register a new user of any role like admin, author and reader
|       +-- Login                           ──────────────> this Login function is used to login the user with email and password and in response is gives jwt token
|       +-- UpdateProfile                   ──────────────> UpdateProfile will update the user profile and user have only access to update own profile
|       +-- GetProfile                      ──────────────> GetProfile will show your own profile information with the help of login token
|       +-- DeleteProfile                   ──────────────> DeleteProfile will delete the own profile with the help of token
|   +-- utils.go                            ──────────────> define basic function that is used by whole projects like, hashing, time conversion etc
|       +-- hashPassword                    ──────────────> hashPassword generates a bcrypt hash of the given password
|       +-- checkPasswordHash               ──────────────> checkPasswordHash compares a password with a bcrypt hash and returns true if they match
|       +-- roleValidator                   ──────────────> roleValidator is used to validate the user role and only accept the specified role
+-- db                                      ──────────────> defined your dataservices layer
|   +-- delete.go                           ──────────────> defined all the delete functions that delete data from database like delete user and delete blog etc 
|       +-- DeleteProfileDB                 ──────────────> DeleteProfileDB deletes a user profile from the specified collection by id
|       +-- DeleteBlogDB                    ──────────────> DeleteBlogDB deletes a blog from the specified collection by id
|   +-- module.go                           ──────────────> defined the connection of database of your project
|   +-- read.go                             ──────────────> defined all the read functions that read the user and blog data from database
|       +-- CheckEmailExistsDB              ──────────────> CheckEmailExistsDB checks if an email already exists in the specified collection
|       +-- CheckUsernameExistsDB           ──────────────> CheckUsernameExistsDB checks if a username already exists in the specified collection
|       +-- GetUserByEmailDB                ──────────────> GetUserByEmailDB retrieves a user by their email from the specified collection
|       +-- GetUserByIdDB                   ──────────────> GetUserByIdDB retrivers the user by their user id from the specified collection
|       +-- GetBlogByIdDB                   ──────────────> GetBlogByIdDB retrivers the blog by their blog id from the specified collection
|       +-- GetBlogsByUsernameDB            ──────────────> GetBlogsByUsernameDB retrivers all the blog of the user by username from the specified collection
|   +-- update.go                           ──────────────> defined all the update functions that update the user and blog data from database
|       +-- UpdateUserDB                    ──────────────> UpdateUserDB updates a user document in the specified database collection
|       +-- UpdateBlogDB                    ──────────────> UpdateBlogDB updates a blog document in the specified database collection
|   +-- write.go                            ──────────────> defined all the insert functions that insert the user and blog data into database
|       +-- RegisterDB                      ──────────────> RegisterDB inserts a user document into the specified database collection
|       +-- CreateBlogDB                    ──────────────> CreateBlogDB inserts a blog document into the specified database collection
+-- middleware.go                           ──────────────> jwt function written over there, for verify and creating the tokens
|   +-- jwt_generator.go                    ──────────────> defined the function that generate the jwt token on login
|       +-- GenerateJWT                     ──────────────> GenerateJWT generates a JWT token based on the provided user information
|   +-- jwt_verify.go                       ──────────────> defined the verification function that verify the login token and extract the user details from token
|       +-- IsAuthorized                    ──────────────> IsAuthorized is a middleware function that checks if the request is authorized and verify the token
|   +-- response.go                         ──────────────> defined the responses of the middleware layer
+-- models                                  ──────────────> defined all the structs of the project
|   +-- errors                              ──────────────> defined the error responses of your project
|   |   +-- errors.go
|   +-- models.go                           ──────────────> defined the main struct of your project like user and blog structs
|   +-- request.go                          ──────────────> defined the request struct of your project
|   +-- response.go                         ──────────────> defined the response struct of your project
+-- routes
|   +-- routes.go                           ──────────────> it contains all the routes of your project
+-- .env                                    ──────────────> it contains all the environment variables of your project
+-- .gitignore                              ──────────────> it defines all the gitignore variables of your project that we not want to function on GitHub
+-- go.mod                                  ──────────────> load up all dependencies required in a project
+-- go.sum                                  ──────────────> load up cryptographic hashes of exact version of dependencies required in a project
+-- main.go                                 ──────────────> heart of the project, its all start from here
+-- README.md                               ──────────────> introduction of the project
```
