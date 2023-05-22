package controller

import (
	"blog/models"
	"blog/models/errors"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *Controller) CreateBlog(rw http.ResponseWriter, r *http.Request) {
	props, _ := r.Context().Value("props").(jwt.MapClaims)

	var body models.CreateBlogRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		errors.ServerErrResponse(err.Error(), rw)
		return
	}

	var blog models.Blog
	if err := copier.Copy(&blog, &body); err != nil {
		errors.ServerErrResponse(err.Error(), rw)
		return
	}

	blog.UserId = props["user_id"].(string)
	blog.Username = props["username"].(string)
	blog.CreatedAt = time.Now().UTC()
	blog.UpdatedAt = time.Now().UTC()
	result, err := c.db.CreateBlogDB(r, "blogs", blog)
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

func (c *Controller) UpdateBlog(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	props, _ := r.Context().Value("props").(jwt.MapClaims)

	var body models.UpdateBlogRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		errors.ServerErrResponse(err.Error(), rw)
		return
	}
	blogId, err := primitive.ObjectIDFromHex(params["blog_id"])
	if err != nil {
		errors.ServerErrResponse(err.Error(), rw)
		return
	}

	blog, err := c.db.GetBlogByIdDB(r, "blogs", blogId)
	if err != nil {
		errors.ErrorResponse(err.Error(), rw)
		return
	}

	if blog.UserId != props["user_id"].(string) {
		errors.AuthorizationResponse("you have not access to update this blog", rw)
		return
	}

	if len(body.Title) != 0 {
		blog.Title = body.Title
	}
	if len(body.Description) != 0 {
		blog.Description = body.Description
	}

	blog.UpdatedAt = time.Now().UTC()

	err = c.db.UpdateBlogDB(r, "blogs", *blog)
	if err != nil {
		errors.ErrorResponse(err.Error(), rw)
		return
	}

	models.SuccessResponse("blog is updated", rw)
}

func (c *Controller) GetBlogById(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	blogId, err := primitive.ObjectIDFromHex(params["blog_id"])
	if err != nil {
		errors.ServerErrResponse(err.Error(), rw)
		return
	}

	blog, err := c.db.GetBlogByIdDB(r, "blogs", blogId)
	if err != nil {
		errors.ErrorResponse(err.Error(), rw)
		return
	}

	var blogResp models.GetBlogResp
	if err := copier.Copy(&blogResp, &blog); err != nil {
		errors.ServerErrResponse(err.Error(), rw)
		return
	}
	models.SuccessRespond(blogResp, rw)
}
