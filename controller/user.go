package controller

import (
	"blog/middleware"
	"blog/models"
	"blog/models/errors"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Register is used to register a new user of any role like Admin, Author and Reader.
func (c *Controller) Register(rw http.ResponseWriter, r *http.Request) {
	var body models.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&body) // Decode the request body into a RegisterRequest struct
	if err != nil {
		errors.ServerErrResponse(err.Error(), rw) // Return a server error response if there is an error decoding the request body
		return
	}

	if err := roleValidator(body.Role); err != nil {
		errors.ErrorResponse(err.Error(), rw) // Return an error response if the role is not exists
		return
	}

	var user models.User                              // Create a new User struct
	if err := copier.Copy(&user, &body); err != nil { // Copy the values from the RegisterRequest to the User struct
		errors.ServerErrResponse(err.Error(), rw) // Return a server error response if there is an error copying the values
		return
	}

	_, err = c.db.CheckEmailExistsDB(r, "users", user.Email) // Check if the email already exists in the database
	if err != nil {
		errors.ErrorResponse(err.Error(), rw) // Return an error response if the email already exists
		return
	}

	_, err = c.db.CheckUsernameExistsDB(r, "users", user.Username) // Check if the username already exists in the database
	if err != nil {
		errors.ErrorResponse(err.Error(), rw) // Return an error response if the username already exists
		return
	}

	passwordHash, err := hashPassword(user.Password) // Hash the user's password
	if err != nil {
		errors.ServerErrResponse(err.Error(), rw) // Return a server error response if there is an error hashing the password
		return
	}

	user.Password = passwordHash
	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = time.Now().UTC()
	result, err := c.db.RegisterDB(r, "users", user) // Register the user in the database
	if err != nil {
		errors.ServerErrResponse(err.Error(), rw) // Return a server error response if there is an error registering the user
		return
	}

	res, err := json.Marshal(result.InsertedID) // Convert the inserted ID to JSON format
	if err != nil {
		errors.ServerErrResponse(err.Error(), rw) // Return a server error response if there is an error converting to JSON
		return
	}

	models.SuccessResponse(`inserted at `+strings.Replace(string(res), `"`, ``, 2), rw) // Return a success response with the inserted ID
}

// This Login function is used to login the user with email and password and in response is gives JWT token.
func (c *Controller) Login(rw http.ResponseWriter, r *http.Request) {
	var body models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&body) // Decode the request body into a LoginRequest struct
	if err != nil {
		errors.ServerErrResponse(err.Error(), rw) // Return a server error response if there is an error decoding the request body
		return
	}

	existingUser, err := c.db.GetUserByEmailDB(r, "users", body.Email) // Retrieve the user by email from the database
	if err != nil {
		errors.ErrorResponse(err.Error(), rw) // Return an error response if there is an error retrieving the user
		return
	}

	isPasswordMatch := checkPasswordHash(body.Password, existingUser.Password) // Check if the provided password matches the stored password
	if !isPasswordMatch {
		errors.ErrorResponse("password doesn't match", rw) // Return an error response if the passwords don't match
		return
	}

	token, err := middleware.GenerateJWT(*existingUser) // Generate a JWT token for the user
	if err != nil {
		errors.ErrorResponse("failed to generate token", rw) // Return an error response if there is an error generating the token
		return
	}

	models.SuccessResponse(*token, rw) // Return a success response with the generated token
}

// UpdateProfile will update the user profile and user have only access to update own profile
func (c *Controller) UpdateProfile(rw http.ResponseWriter, r *http.Request) {
	props, _ := r.Context().Value("props").(jwt.MapClaims) // Get the user properties from the request context

	userId, err := primitive.ObjectIDFromHex(props["user_id"].(string)) // Extract the user ID from the user properties
	if err != nil {
		errors.ServerErrResponse(err.Error(), rw) // Return a server error response if there is an error parsing the user ID
		return
	}

	var body models.UpdateProfileRequest
	err = json.NewDecoder(r.Body).Decode(&body) // Decode the request body into an UpdateProfileRequest struct
	if err != nil {
		errors.ServerErrResponse(err.Error(), rw) // Return a server error response if there is an error decoding the request body
		return
	}

	_, err = c.db.CheckUsernameExistsDB(r, "users", body.Username) // Check if the username already exists in the database
	if err != nil {
		errors.ErrorResponse(err.Error(), rw) // Return an error response if the username already exists
		return
	}

	user, err := c.db.GetUserByIdDB(r, "users", userId) // Retrieve the user by ID from the database
	if err != nil {
		errors.ErrorResponse(err.Error(), rw) // Return an error response if there is an error retrieving the user
		return
	}

	// Update the user properties if they are provided in the request body
	if len(body.Username) != 0 {
		user.Username = body.Username
	}
	if len(body.FullName) != 0 {
		user.FullName = body.FullName
	}
	if len(body.Role) != 0 {
		user.Role = body.Role
	}
	if len(body.Bio) != 0 {
		user.Bio = body.Bio
	}

	user.UpdatedAt = time.Now().UTC() // Update the "UpdatedAt" field with the current time

	err = c.db.UpdateUserDB(r, "users", *user) // Update the user in the database
	if err != nil {
		errors.ErrorResponse(err.Error(), rw) // Return an error response if there is an error updating the user
		return
	}

	token, err := middleware.GenerateJWT(*user) // Generate a JWT token for the updated user
	if err != nil {
		errors.ErrorResponse("failed to generate token", rw) // Return an error response if there is an error generating the token
		return
	}

	models.SuccessResponse(*token, rw) // Return a success response with the generated token
}

// GetProfile will show your own profile information with the help of login token
func (c *Controller) GetProfile(rw http.ResponseWriter, r *http.Request) {
	props, _ := r.Context().Value("props").(jwt.MapClaims) // Get the user properties from the request context

	userId, err := primitive.ObjectIDFromHex(props["user_id"].(string)) // Extract the user ID from the user properties
	if err != nil {
		errors.ServerErrResponse(err.Error(), rw) // Return a server error response if there is an error parsing the user ID
		return
	}

	user, err := c.db.GetUserByIdDB(r, "users", userId) // Retrieve the user by ID from the database
	if err != nil {
		errors.ErrorResponse(err.Error(), rw) // Return an error response if there is an error retrieving the user
		return
	}

	var userResp models.GetProfileResp
	if err := copier.Copy(&userResp, &user); err != nil {
		errors.ServerErrResponse(err.Error(), rw) // Return a server error response if there is an error copying the user data to the response struct
		return
	}

	models.SuccessRespond(userResp, rw) // Return a success response with the user profile data
}

// DeleteProfile will delete the own profile with the help of token
func (c *Controller) DeleteProfile(rw http.ResponseWriter, r *http.Request) {
	props, _ := r.Context().Value("props").(jwt.MapClaims) // Get the user properties from the request context

	userId, err := primitive.ObjectIDFromHex(props["user_id"].(string)) // Extract the user ID from the user properties
	if err != nil {
		errors.ServerErrResponse(err.Error(), rw) // Return a server error response if there is an error parsing the user ID
		return
	}

	err = c.db.DeleteProfileDB(r, "users", userId) // Delete the user profile from the database
	if err != nil {
		errors.ErrorResponse(err.Error(), rw) // Return an error response if there is an error deleting the user profile
		return
	}

	models.SuccessRespond("user deleted successfully", rw) // Return a success response indicating that the user profile has been deleted
}
