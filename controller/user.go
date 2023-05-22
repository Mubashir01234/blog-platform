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

func (c *Controller) Register(rw http.ResponseWriter, r *http.Request) {
	var body models.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		errors.ServerErrResponse(err.Error(), rw)
		return
	}

	var user models.User
	if err := copier.Copy(&user, &body); err != nil {
		errors.ServerErrResponse(err.Error(), rw)
		return
	}
	_, err = c.db.CheckEmailExistsDB(r, "users", user.Email)
	if err != nil {
		errors.ErrorResponse(err.Error(), rw)
		return
	}

	_, err = c.db.CheckUsernameExistsDB(r, "users", user.Username)
	if err != nil {
		errors.ErrorResponse(err.Error(), rw)
		return
	}

	passwordHash, err := hashPassword(user.Password)
	if err != nil {
		errors.ServerErrResponse(err.Error(), rw)
		return
	}

	user.Password = passwordHash
	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = time.Now().UTC()
	result, err := c.db.RegisterDB(r, "users", user)
	if err != nil {
		errors.ServerErrResponse(err.Error(), rw)
		return
	}

	res, err := json.Marshal(result.InsertedID)
	if err != nil {
		errors.ServerErrResponse(err.Error(), rw)
		return
	}

	models.SuccessResponse(`inserted at `+strings.Replace(string(res), `"`, ``, 2), rw)
}

func (c *Controller) Login(rw http.ResponseWriter, r *http.Request) {
	var body models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		errors.ServerErrResponse(err.Error(), rw)
		return
	}

	existingUser, err := c.db.GetUserByEmailDB(r, "users", body.Email)
	if err != nil {
		errors.ErrorResponse(err.Error(), rw)
		return
	}

	isPasswordMatch := checkPasswordHash(body.Password, existingUser.Password)
	if !isPasswordMatch {
		errors.ErrorResponse("password doesn't match", rw)
		return
	}

	token, err := middleware.GenerateJWT(*existingUser)
	if err != nil {
		errors.ErrorResponse("failed to generate token", rw)
		return
	}

	models.SuccessResponse(*token, rw)
}

func (c *Controller) UpdateProfile(rw http.ResponseWriter, r *http.Request) {
	props, _ := r.Context().Value("props").(jwt.MapClaims)

	userId, err := primitive.ObjectIDFromHex(props["user_id"].(string))
	if err != nil {
		errors.ServerErrResponse(err.Error(), rw)
		return
	}

	var body models.UpdateProfileRequest
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		errors.ServerErrResponse(err.Error(), rw)
		return
	}

	_, err = c.db.CheckUsernameExistsDB(r, "users", body.Username)
	if err != nil {
		errors.ErrorResponse(err.Error(), rw)
		return
	}

	user, err := c.db.GetUserByIdDB(r, "users", userId)
	if err != nil {
		errors.ErrorResponse(err.Error(), rw)
		return
	}

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

	user.UpdatedAt = time.Now().UTC()

	err = c.db.UpdateUserDB(r, "users", *user)
	if err != nil {
		errors.ErrorResponse(err.Error(), rw)
		return
	}

	token, err := middleware.GenerateJWT(*user)
	if err != nil {
		errors.ErrorResponse("failed to generate token", rw)
		return
	}

	models.SuccessResponse(*token, rw)
}

func (c *Controller) GetProfile(rw http.ResponseWriter, r *http.Request) {
	props, _ := r.Context().Value("props").(jwt.MapClaims)

	userId, err := primitive.ObjectIDFromHex(props["user_id"].(string))
	if err != nil {
		errors.ServerErrResponse(err.Error(), rw)
		return
	}

	user, err := c.db.GetUserByIdDB(r, "users", userId)
	if err != nil {
		errors.ErrorResponse(err.Error(), rw)
		return
	}

	var userResp models.GetProfileResp
	if err := copier.Copy(&userResp, &user); err != nil {
		errors.ServerErrResponse(err.Error(), rw)
		return
	}

	models.SuccessRespond(userResp, rw)
}

func (c *Controller) DeleteProfile(rw http.ResponseWriter, r *http.Request) {
	props, _ := r.Context().Value("props").(jwt.MapClaims)

	userId, err := primitive.ObjectIDFromHex(props["user_id"].(string))
	if err != nil {
		errors.ServerErrResponse(err.Error(), rw)
		return
	}

	err = c.db.DeleteProfileDB(r, "users", userId)
	if err != nil {
		errors.ErrorResponse(err.Error(), rw)
		return
	}
	models.SuccessRespond("user deleted successfully", rw)
}
