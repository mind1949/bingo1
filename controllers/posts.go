package controllers

import (
	"github.com/mind1949/bingo1/models"
	"github.com/mind1949/bingo1/render"
	"io"
	"log"
)

const (
	TemplateShow = "views/posts/show.html"
	TemplateIndex = "views/posts/index.html"
)

func ShowPost(w io.Writer, urlpath string) {
	post, err := models.FindPost(urlpath)
	if err != nil {
		log.Fatal(err)
	}

	render.Execute(w, TemplateShow, &post)
}

func Posts(w io.Writer, filepath string) {
	posts, err := models.FindAllPosts()
	if err != nil {
		log.Fatal(err)
	}

	render.Execute(w, TemplateIndex, &posts)
}
