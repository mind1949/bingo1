package controllers

import (
	"fmt"
	"io"
	"github.com/mind1949/bingo1/models"
	"log"
)

func ShowPost(w io.Writer, filepath string) {
	post, err := models.NewPost(filepath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, post)
}

