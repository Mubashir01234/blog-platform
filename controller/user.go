package controller

import (
	"blog/models"
	"blog/models/errors"
	"encoding/json"
	"net/http"

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

}
