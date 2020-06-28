package monitor

import (
	"net/http"

	"github.com/blogCRUDWebsocket/internal/pkg/http/router"
)

func (h *Handler) Routes() []router.Route {
	return []router.Route{
		{
			Path:    "/api/v1/monitor/room",
			Handler: h.ConnectMonitor,
			Method:  http.MethodGet,
		},
	}
}
