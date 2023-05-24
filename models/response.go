package models

import (
	"blog/models/errors"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Response is used to show the response with status code, message and Data related to request.
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SuccessArrRespond -> response formatter
func SuccessArrRespond(fields interface{}, writer http.ResponseWriter) {
	// var fields["status"] := "success"
	_, err := json.Marshal(fields)
	type data struct {
		Data       interface{} `json:"data"`
		Statuscode int         `json:"status"`
		Message    string      `json:"msg"`
	}
	temp := &data{Data: fields, Statuscode: 200, Message: "success"}
	if err != nil {
		errors.ServerErrResponse(err.Error(), writer)
	}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(temp)
}

// SuccessRespond -> response formatter
func SuccessRespond(fields interface{}, writer http.ResponseWriter) {
	_, err := json.Marshal(fields)
	type data struct {
		Data       interface{} `json:"data"`
		Statuscode int         `json:"status"`
		Message    string      `json:"msg"`
	}
	temp := &data{Data: fields, Statuscode: 200, Message: "success"}
	if err != nil {
		errors.ServerErrResponse(err.Error(), writer)
	}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(temp)
}

// SuccessResponse -> success formatter
func SuccessResponse(msg string, writer http.ResponseWriter) {
	type errdata struct {
		Statuscode int    `json:"status"`
		Message    string `json:"msg"`
	}
	temp := &errdata{Statuscode: 200, Message: msg}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(temp)
}

// GetProfileResp is used to get the user profile
type GetProfileResp struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Username  string             `json:"username" bson:"username,omitempty"`
	Role      string             `json:"role" bson:"role,omitempty"`
	Email     string             `json:"email" bson:"email,omitempty"`
	FullName  string             `json:"full_name" bson:"full_name,omitempty"`
	Bio       string             `json:"bio" bson:"bio,omitempty"`
	IsNew     bool               `json:"is_new" bson:"is_new,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
}

// GetBlogResp is used to get the blog information
type GetBlogResp struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UserId      string             `json:"user_id" bson:"user_id,omitempty"`
	Username    string             `json:"username" bson:"username,omitempty"`
	Title       string             `json:"title" bson:"title,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
}
