package controller

import (
	"blog/models"
	"blog/models/errors"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/jinzhu/copier"
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
	_, err = c.db.GetUserDataByEmailDB(r, "users", user.Email)
	if err != nil {
		errors.ErrorResponse(err.Error(), rw)
		return
	}

	_, err = c.db.GetUserDataByUsernameDB(r, "users", user.Username)
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

	res, _ := json.Marshal(result.InsertedID)
	if err != nil {
		errors.ServerErrResponse(err.Error(), rw)
		return
	}

	models.SuccessResponse(`inserted at `+strings.Replace(string(res), `"`, ``, 2), rw)
}
