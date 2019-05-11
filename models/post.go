package models

import (
	"strings"
)

type Post struct {
	Title   string
	Content []byte
}

func NewPost(filepath string) (*Post, error) {
	title := getPostTitle(filepath)
	content := getPostContent(filepath)

	return &Post{
		Title:   title,
		Content: content,
	}, nil
}

func getPostTitle(filepath string) string {
	return strings.Split(strings.TrimLeft(filepath, "/"), "/")[1]
}

func getPostContent(filepath string) []byte {
	return []byte("<p>post content</p>")
}
