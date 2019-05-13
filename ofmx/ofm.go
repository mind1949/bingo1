package ofmx

import (
	"bytes"
	"fmt"
	"github.com/russross/blackfriday"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"
)

func getPostInfo(filepath string) (postInfo []byte, err error) {
	postInfo, err = ioutil.ReadFile(filepath)
	return
}

func ParsePostInfo(postInfo []byte) (yml []byte, html []byte) {
	yml, md := filter(postInfo)
	html = blackfriday.MarkdownCommon(md)
	return
}

func filter(postInfo []byte) (yml []byte, md []byte) {
	postInfo = bytes.TrimSpace(postInfo)
	sep := []byte("---\n")
	re := bytes.SplitN(postInfo, sep, 3)
	yml, md = re[1], re[2]
	return
}

func unmarshal(out interface{}, postInfo []byte) (err error) {
	yml, html := ParsePostInfo(postInfo)
	v := reflect.ValueOf(out).Elem()

	switch v.Kind() {
	case reflect.Struct:
		err = yaml.Unmarshal(yml, v.FieldByName("Meta").Interface())
		v.FieldByName("Content").Set(reflect.ValueOf(html))
	default:
		err = fmt.Errorf("cannot decode postInfo into %v", v.Type())
	}

	return
}

func Scan(filepath string, to interface{}) (err error) {
	postInfo, err := getPostInfo(filepath)
	if err != nil {
		err = fmt.Errorf("Scan(container, filepath): %s", err.Error())
		return
	}

	err = unmarshal(to, postInfo)
	return
}
