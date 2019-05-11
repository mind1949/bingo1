// bingo 路由器
package router

import (
	"fmt"
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
		fmt.Fprint(w, "posts index")
	}
	// show post
	if strings.HasPrefix(path, "/posts/") && path != "/posts/" {
		fmt.Fprintf(w, "post: %s", "postname")
	}
	// tags index
	if path == "/tags" || path == "/tags/" {
		fmt.Fprint(w, "tags index")
	}
}
