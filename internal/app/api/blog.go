package api

import (
	"github.com/blogCRUDWebsocket/internal/app/blog"
	"github.com/blogCRUDWebsocket/internal/app/monitor"
)

func newBlogService(room *monitor.Room) *blog.Service {
	var inMemDb []blog.Blog
	db := blog.NewDB(inMemDb)

	return blog.NewService(db, room)
}

func newBlogHandler(srv *blog.Service) *blog.Handler {
	return blog.NewHandler(srv)
}
