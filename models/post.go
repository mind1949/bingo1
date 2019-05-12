package models

import (
	"github.com/mind1949/bingo1/ofm"
	yaml "gopkg.in/yaml.v2"
	"fmt"
)

type Meta struct {
	Title, Date, Short string
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
		fmt.Printf("error yaml.unmarshal: %s\n", err)
	}
	post.Content = slice[1]

	return &post, nil
}

func FindAllPosts() ([]*Post, error) {
	postsInfo, err := ofm.FindAll()
	if err != nil {
		return nil, err
	}

	posts := []*Post{}
	for _, postInfo := range postsInfo {
		post := Post{Meta: &Meta{}}
		if err := yaml.Unmarshal(postInfo[0], post.Meta); err != nil {
			fmt.Printf("error yaml.unmarshal: %s\n", err)
			continue
		}
		posts = append(posts, &post)
	}

	return posts, nil
}

func getFilepath(urlpath string) string {
	return ("." + urlpath + ".md")
}
