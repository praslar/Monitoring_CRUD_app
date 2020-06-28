package blog

import (
	"fmt"
	"time"

	"github.com/blogCRUDWebsocket/internal/app/monitor"
	"github.com/google/uuid"
)

type (
	database interface {
		Create(b *Blog) error
		Read() ([]Blog, error)
		Update(string, *CreateBlog)
		Delete(string) error

		FindByID(string) (*Blog, error)
	}

	MonitorRoom interface {
		Run()
		Write(msg monitor.ChangeInfo)
		Leave(client *monitor.Client)
		Join(client *monitor.Client)
	}

	Service struct {
		mr MonitorRoom
		db database
	}
)

func NewService(db database, mr MonitorRoom) *Service {
	return &Service{
		db: db,
		mr: mr,
	}
}

func (s *Service) Create(req *CreateBlog) (*Blog, error) {

	if req.Content == "" || req.Title == "" {
		return nil, fmt.Errorf("Empty value! %s", req.Content)
	}

	blog := &Blog{
		Content: req.Content,
		ID:      uuid.New().String(),
		Title:   req.Title,
	}

	if err := s.db.Create(blog); err != nil {
		return nil, err
	}
	msg := monitor.ChangeInfo{
		Type:   "Create new blog",
		Detail: fmt.Sprintf("ID: %s ; Title: %s; Contect: %s", blog.ID, blog.Title, blog.Content),
	}
	s.mr.Write(msg)
	return blog, nil
}
func (s *Service) Read() ([]Blog, error) {
	return s.db.Read()
}
func (s *Service) Update(id string, req *CreateBlog) (*Blog, error) {
	blog, err := s.db.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("Blog not found")
	}
	s.db.Update(id, req)
	blog = &Blog{
		Content:  req.Content,
		Title:    req.Title,
		ID:       blog.ID,
		UpdateAt: time.Now(),
	}
	msg := monitor.ChangeInfo{
		Type:   fmt.Sprintf("Update blog ID: %s", id),
		Detail: fmt.Sprintf("ID: %s ; Title: %s; Contect: %s", blog.ID, blog.Title, blog.Content),
	}
	s.mr.Write(msg)
	return blog, nil
}
func (s *Service) Delete(id string) error {
	_, err := s.db.FindByID(id)
	if err != nil {
		return fmt.Errorf("Blog not found")
	}
	msg := monitor.ChangeInfo{
		Type: fmt.Sprintf("Delete blog ID: %s", id),
	}
	s.mr.Write(msg)
	return s.db.Delete(id)
}
