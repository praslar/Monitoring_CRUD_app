package blog_test

import (
	"fmt"
	"testing"

	"github.com/blogCRUDWebsocket/internal/app/blog"
	"github.com/blogCRUDWebsocket/internal/app/monitor"
	"github.com/golang/mock/gomock"
)

func TestCreate(t *testing.T) {
	Mockdatabase := NewMockdatabase(gomock.NewController(t))
	MockMonitorRoom := NewMockMonitorRoom(gomock.NewController(t))
	service := blog.NewService(Mockdatabase, MockMonitorRoom)
	testCases := []struct {
		desc   string
		expect func()
		input  *blog.CreateBlog
		output blog.Blog
		err    error
	}{
		{
			desc:   "Empty Value",
			expect: func() {},
			input: &blog.CreateBlog{
				Content: "",
				Title:   "",
			},
			err:    fmt.Errorf("Empty value! %s", ""),
			output: blog.Blog{},
		},
		{
			desc: "Created",
			expect: func() {
				Mockdatabase.Create(&blog.Blog{Content: "test", Title: "test", ID: "test"})
				MockMonitorRoom.Write(monitor.ChangeInfo{Detail: "test", Type: "test"})
			},
			input: &blog.CreateBlog{
				Content: "test",
				Title:   "test",
			},
			err: nil,
			output: blog.Blog{
				Content: "test",
				ID:      "test",
				Title:   "test",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.expect()
			blog, _ := service.Create(tC.input)
			if blog.Content != tC.output.Content || blog.ID != tC.output.ID || blog.Title != tC.output.Title {
				t.Errorf("Got but want: %s : %s", blog, tC.output)
			}
		})
	}
}
