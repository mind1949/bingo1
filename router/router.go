// bingo 路由器
package router

import (
	"fmt"
	"github.com/mind1949/bingo1/controllers"
	"net/http"
	"strings"
)

type Router struct{}

func New() *Router {
	return &Router{}
}

func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	// home page
	if path == "/" {
		fmt.Fprint(w, "hello bingo")
	}
	// posts index
	if path == "/posts" || path == "/posts/" {
		controllers.Posts(w, path)
	}
	// show post
	if strings.HasPrefix(path, "/posts/") && path != "/posts/" {
		controllers.ShowPost(w, path)
	}
	// tags index
	if path == "/tags" || path == "/tags/" {
		fmt.Fprint(w, "tags index")
	}
}
