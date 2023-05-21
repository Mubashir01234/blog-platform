package api

import (
	"blog/db"
)

type BlogAPI interface {
}
type BlogAPIImpl struct {
	db db.BlogDB
}

func NewBlogAPIImpl() *BlogAPIImpl {
	db := db.ConnectDB()
	return &BlogAPIImpl{
		db: db,
	}
}

var _ BlogAPI = &BlogAPIImpl{}
