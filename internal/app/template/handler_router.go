package template

import (
	"net/http"

	"github.com/blogCRUDWebsocket/internal/pkg/http/router"
)

func (h *Handler) Routes() []router.Route {
	return []router.Route{
		{
			Path:    "/api/v1/monitor",
			Handler: h.ServeTemplate,
			Method:  http.MethodGet,
		},
	}
}
