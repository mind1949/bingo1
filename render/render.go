package render

import (
	"html/template"
	"io"
)

func toTemplateHTML(content []byte) template.HTML {
	return template.HTML(string(content))
}

func Execute(w io.Writer, filename string, data interface{}) {
	filenames := []string{"views/layouts/application.html", filename}

	funcMap := template.FuncMap{"toHTML": toTemplateHTML}
	t := template.Must(template.ParseFiles(filenames[0]))
	tc := t.New("content").Funcs(funcMap)
	tc.ParseFiles(filenames[1])

	t.Execute(w, data)
}
