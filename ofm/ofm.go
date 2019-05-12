package ofm

import (
	"bytes"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"fmt"
	"path/filepath"
)

const(
	postsDir = "posts"
)

func Find(slice [][]byte, filepath string) error {
	fileContent, err := ioutil.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("ioutil.ReadFile: %q", err)
	}
	slice[0], slice[1] = filter(fileContent)

	return nil
}

func FindAll() ([][][]byte, error) {
	postsInfo := [][][]byte{}
	fileinfos, err := ioutil.ReadDir(postsDir)
	if err != nil {
		return nil, err
	}
	for _, fileinfo := range fileinfos {
		if fileinfo.IsDir() {
			continue
		}

		postInfo := make([][]byte, 2)
		if err := Find(postInfo, filepath.Join(postsDir, fileinfo.Name())); err != nil {
			return nil, err
		}
		postsInfo = append(postsInfo, postInfo)
	}

	return postsInfo, nil
}

func filter(filecontent []byte) (yamlRaw []byte, htmlContent []byte) {
	yamlRaw, markdownContent := splitContent(filecontent)
	htmlContent = blackfriday.MarkdownCommon(markdownContent)
	return
}

func splitContent(filecontent []byte) ([]byte, []byte) {
	filecontent = bytes.TrimSpace(filecontent)
	sep := []byte("---\n")
	result := bytes.SplitN(filecontent, sep, 3)
	return result[1], result[2]
}

