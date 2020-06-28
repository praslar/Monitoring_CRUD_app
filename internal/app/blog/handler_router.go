package blog

import (
	"net/http"

	"github.com/blogCRUDWebsocket/internal/pkg/http/router"
)

func (h *Handler) Routes() []router.Route {
	return []router.Route{
		{
			Path:    "/api/v1/blog",
			Handler: h.Create,
			Method:  http.MethodPost,
		},
		{
			Path:    "/api/v1/blog",
			Handler: h.Read,
			Method:  http.MethodGet,
		},
		{
			Path:    "/api/v1/blog/{blog_id:[a-z0-9-\\-]+}",
			Handler: h.Delete,
			Method:  http.MethodDelete,
		},
		{
			Path:    "/api/v1/blog/{blog_id:[a-z0-9-\\-]+}",
			Handler: h.Update,
			Method:  http.MethodPut,
		},
	}
}
