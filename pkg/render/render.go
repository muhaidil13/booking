package render

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderFile(w http.ResponseWriter, form string) {
	tmp, _ := template.ParseFiles("./template/" + form)
	err := tmp.Execute(w, nil)
	if err != nil {
		fmt.Println("kesalahan pada", err)
		return
	}
}
