package ofm

import (
	"bytes"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"fmt"
)

func Find(slice [][]byte, filepath string) error {
	fileContent, err := ioutil.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("ioutil.ReadFile: %q", err)
	}
	slice[0], slice[1] = filter(fileContent)

	return nil
}

func filter(filecontent []byte) (yamlRaw []byte, htmlContent []byte) {
	yamlRaw, markdownContent := splitContent(filecontent)
	htmlContent = blackfriday.MarkdownCommon(markdownContent)
	return
}

func splitContent(filecontent []byte) ([]byte, []byte) {
	sep := []byte("---\n")
	result := bytes.Split(filecontent, sep)
	return result[1], result[2]
}

