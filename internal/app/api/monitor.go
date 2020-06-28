package api

import "github.com/blogCRUDWebsocket/internal/app/monitor"

func newMonitorHandler(srv *monitor.Room) *monitor.Handler {
	return monitor.NewHandler(srv)
}

func newMonitorRoom() *monitor.Room {
	return monitor.NewRoom()
}
