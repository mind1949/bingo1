package models

import (
	"github.com/mind1949/bingo1/ofm"
	yaml "gopkg.in/yaml.v2"
	"fmt"
)

type Meta struct {
	Title, Date string
	Tags  []string
}

type Post struct {
	Content []byte
	Meta    *Meta
}

func FindPost(urlpath string) (*Post, error) {
	// TODO(更改为传入一个post实例，而不是slice)
	filepath := getFilepath(urlpath)
	slice := make([][]byte, 2)
	if err := ofm.Find(slice, filepath); err != nil {
		fmt.Printf("error ofmt.Find: %s", err)
	}

	post := Post{Meta: &Meta{}}
	if err := yaml.Unmarshal(slice[0], post.Meta); err != nil {
		fmt.Printf("error yaml.Unmarshal: %s\n", err)
	}
	post.Content = slice[1]

	return &post, nil
}

func getFilepath(urlpath string) string {
	return ("." + urlpath + ".md")
}
