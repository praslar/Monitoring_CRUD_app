package blog

import (
	"fmt"
	"time"
)

type (
	inMemDB struct {
		db []Blog
	}
)

func NewDB(db []Blog) *inMemDB {
	return &inMemDB{
		db: db,
	}
}

func (r *inMemDB) Create(b *Blog) error {
	b.CreatedAt = time.Now()
	b.UpdateAt = b.CreatedAt
	r.db = append(r.db, *b)

	return nil
}
func (r *inMemDB) Read() ([]Blog, error) {
	if len(r.db) == 0 {
		return nil, fmt.Errorf("Empty blog")
	}
	return r.db, nil
}

func (r *inMemDB) Update(id string, req *CreateBlog) {
	fmt.Println("Updating: ", req)
	for i := 0; i < len(r.db); i++ {
		if r.db[i].ID == id {
			if req.Content != "" {
				r.db[i].Title = req.Title
			}
			if req.Title != "" {
				r.db[i].Content = req.Content
			}
			r.db[i].UpdateAt = time.Now()
		}
		break
	}
}
func (r *inMemDB) Delete(id string) error {
	for i := 0; i < len(r.db); i++ {
		if r.db[i].ID == id {
			// remove
			r.db[i] = r.db[len(r.db)-1]
			r.db = r.db[:len(r.db)-1]
		}
		break
	}
	return nil
}

func (r *inMemDB) FindByID(id string) (*Blog, error) {
	for _, blog := range r.db {
		if blog.ID == id {
			return &blog, nil
		}
	}
	return nil, fmt.Errorf("Not found")
}
