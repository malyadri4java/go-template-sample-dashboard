package utils

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

var templates *template.Template

func LoadTemplates(pattern string) {
	//templates = template.Must(template.ParseGlob(pattern))
}

func ExecuteTemplate(w http.ResponseWriter, fileName string, data interface{}) {
	lp := path.Join("templates", "layout.html")
	fp := path.Join("templates/view", fileName)
	templates, err := templates.ParseFiles(lp, fp)
	err = templates.ExecuteTemplate(w, "layout", data)
	logFatal(err)
}

func ExecuteTemplateWitoutlayout(w http.ResponseWriter, fileName string, data interface{}) {
	fp := path.Join("templates/view", fileName)
	templates, err := templates.ParseFiles(fp)
	err = templates.ExecuteTemplate(w, fileName, data)
	logFatal(err)
}

func logFatal(err error) {
	if err != nil {
		log.Println(err)
	}
}
