package controllers

import (
	"github.com/mind1949/bingo1/models"
	"github.com/mind1949/bingo1/render"
	"io"
	"log"
)

func ShowPost(w io.Writer, filepath string) {
	post, err := models.NewPost(filepath)
	if err != nil {
		log.Fatal(err)
	}

	render.Execute(w, "views/posts/show.html", &post)
}
