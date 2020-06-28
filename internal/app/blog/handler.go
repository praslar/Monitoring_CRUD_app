package blog

import (
	"encoding/json"
	"net/http"

	"github.com/blogCRUDWebsocket/internal/pkg/http/response"
	"github.com/gorilla/mux"
)

type (
	service interface {
		Create(req *CreateBlog) (*Blog, error)
		Read() ([]Blog, error)
		Update(id string, req *CreateBlog) (*Blog, error)
		Delete(id string) error
	}

	Handler struct {
		srv service
	}
)

func NewHandler(srv service) *Handler {
	return &Handler{
		srv: srv,
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {

	var req CreateBlog

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	info, err := h.srv.Create(&req)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, info)

}
func (h *Handler) Read(w http.ResponseWriter, r *http.Request) {

	infos, err := h.srv.Read()
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, infos)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["blog_id"]

	if id == "" {
		response.JSON(w, http.StatusBadRequest, "invalid id")
		return
	}

	var req CreateBlog

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	info, err := h.srv.Update(id, &req)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, info)

}
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["blog_id"]

	if id == "" {
		response.JSON(w, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.srv.Delete(id); err != nil {
		response.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, id)
}
