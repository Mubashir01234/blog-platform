package controller

import (
	"blog/models"        // Importing the models package
	"blog/models/errors" // Importing the errors package for handling errors
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"                  // Importing the jwt package for JWT token handling
	"github.com/gorilla/mux"                     // Importing the gorilla/mux package for router functionality
	"github.com/jinzhu/copier"                   // Importing the copier package for copying values between structs
	"go.mongodb.org/mongo-driver/bson/primitive" // Importing the primitive package for MongoDB operations
)

func (c *Controller) CreateBlog(rw http.ResponseWriter, r *http.Request) {
	props, _ := r.Context().Value("props").(jwt.MapClaims) // Extracting the properties from the request context, assuming it contains JWT claims

	if props["role"].(string) != "Author" && props["role"].(string) != "Admin" {
		errors.AuthorizationResponse("you have not access to upload blog", rw) // Returning an authorization error response
		return
	}

	var body models.CreateBlogRequest
	err := json.NewDecoder(r.Body).Decode(&body) // Decoding the request body into the CreateBlogRequest struct
	if err != nil {
		errors.ServerErrResponse(err.Error(), rw) // Returning a server error response if there is an error decoding the request body
		return
	}

	var blog models.Blog                              // Creating a new Blog struct
	if err := copier.Copy(&blog, &body); err != nil { // Copying the values from the CreateBlogRequest struct to the Blog struct
		errors.ServerErrResponse(err.Error(), rw) // Returning a server error response if there is an error copying the values
		return
	}

	blog.UserId = props["user_id"].(string)    // Setting the UserId field of the Blog struct with the "user_id" claim from the JWT
	blog.Username = props["username"].(string) // Setting the Username field of the Blog struct with the "username" claim from the JWT
	blog.CreatedAt = time.Now().UTC()          // Setting the CreatedAt field of the Blog struct with the current UTC time
	blog.UpdatedAt = time.Now().UTC()          // Setting the UpdatedAt field of the Blog struct with the current UTC time

	result, err := c.db.CreateBlogDB(r, "blogs", blog) // Calling the CreateBlogDB function of the database to create a new blog entry
	if err != nil {
		errors.ServerErrResponse(err.Error(), rw) // Returning a server error response if there is an error creating the blog entry
		return
	}

	res, err := json.Marshal(result.InsertedID) // Marshaling the inserted ID into JSON format
	if err != nil {
		errors.ServerErrResponse(err.Error(), rw) // Returning a server error response if there is an error marshaling the result
		return
	}

	models.SuccessResponse(`inserted at `+strings.Replace(string(res), `"`, ``, 2), rw) // Returning a success response with the inserted ID
}

func (c *Controller) UpdateBlog(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)                                  // Getting the request parameters from the URL
	props, _ := r.Context().Value("props").(jwt.MapClaims) // Extracting the properties from the request context, assuming it contains JWT claims

	var body models.UpdateBlogRequest
	err := json.NewDecoder(r.Body).Decode(&body) // Decoding the request body into the UpdateBlogRequest struct
	if err != nil {
		errors.ServerErrResponse(err.Error(), rw) // Returning a server error response if there is an error decoding the request body
		return
	}
	blogId, err := primitive.ObjectIDFromHex(params["blog_id"]) // Converting the blog_id parameter from string to ObjectID
	if err != nil {
		errors.ServerErrResponse(err.Error(), rw) // Returning a server error response if there is an error converting the blog_id
		return
	}

	blog, err := c.db.GetBlogByIdDB(r, "blogs", blogId) // Retrieving the blog by ID from the database
	if err != nil {
		errors.ErrorResponse(err.Error(), rw) // Returning an error response if there is an error retrieving the blog
		return
	}

	if props["role"].(string) != "Admin" {
		if blog.UserId != props["user_id"].(string) { // Checking if the user is authorized to update the blog
			errors.AuthorizationResponse("you have not access to update this blog", rw) // Returning an authorization error response
			return
		}
	}

	if len(body.Title) != 0 { // Updating the blog title if it is provided in the request body
		blog.Title = body.Title
	}
	if len(body.Description) != 0 { // Updating the blog description if it is provided in the request body
		blog.Description = body.Description
	}

	blog.UpdatedAt = time.Now().UTC() // Updating the UpdatedAt field of the blog to the current UTC time

	err = c.db.UpdateBlogDB(r, "blogs", *blog) // Updating the blog in the database
	if err != nil {
		errors.ErrorResponse(err.Error(), rw) // Returning an error response if there is an error updating the blog
		return
	}

	models.SuccessResponse("blog is updated", rw) // Returning a success response
}

func (c *Controller) GetBlogById(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Getting the request parameters from the URL

	blogId, err := primitive.ObjectIDFromHex(params["blog_id"]) // Converting the blog_id parameter from string to ObjectID
	if err != nil {
		errors.ServerErrResponse(err.Error(), rw) // Returning a server error response if there is an error converting the blog_id
		return
	}

	blog, err := c.db.GetBlogByIdDB(r, "blogs", blogId) // Retrieving the blog by ID from the database
	if err != nil {
		errors.ErrorResponse(err.Error(), rw) // Returning an error response if there is an error retrieving the blog
		return
	}

	var blogResp models.GetBlogResp                       // Creating a new GetBlogResp struct
	if err := copier.Copy(&blogResp, &blog); err != nil { // Copying the values from the retrieved blog to the GetBlogResp struct
		errors.ServerErrResponse(err.Error(), rw) // Returning a server error response if there is an error copying the values
		return
	}
	models.SuccessRespond(blogResp, rw) // Returning a success response with the GetBlogResp struct
}

func (c *Controller) DeleteBlog(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Getting the request parameters from the URL

	props, _ := r.Context().Value("props").(jwt.MapClaims) // Extracting the properties from the request context, assuming it contains JWT claims

	blogId, err := primitive.ObjectIDFromHex(params["blog_id"]) // Converting the blog_id parameter from string to ObjectID
	if err != nil {
		errors.ServerErrResponse(err.Error(), rw) // Returning a server error response if there is an error converting the blog_id
		return
	}

	blog, err := c.db.GetBlogByIdDB(r, "blogs", blogId) // Retrieving the blog by ID from the database
	if err != nil {
		errors.ErrorResponse(err.Error(), rw) // Returning an error response if there is an error retrieving the blog
		return
	}

	if props["role"].(string) != "Admin" {
		if blog.UserId != props["user_id"].(string) { // Checking if the user is authorized to delete the blog
			errors.AuthorizationResponse("you have not access to delete this blog", rw) // Returning an authorization error response
			return
		}
	}

	err = c.db.DeleteBlogDB(r, "blogs", blogId) // Deleting the blog from the database
	if err != nil {
		errors.ErrorResponse(err.Error(), rw) // Returning an error response if there is an error deleting the blog
		return
	}
	models.SuccessRespond("blog deleted successfully", rw) // Returning a success response
}

func (c *Controller) GetUserAllBlogsByUsername(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Getting the request parameters from the URL

	blogs, err := c.db.GetBlogsByUsernameDB(r, "blogs", params["username"]) // Retrieving all blogs by username from the database
	if err != nil {
		errors.ErrorResponse(err.Error(), rw) // Returning an error response if there is an error retrieving the blogs
		return
	}

	models.SuccessArrRespond(blogs, rw) // Returning a success response with the array of blogs
}
