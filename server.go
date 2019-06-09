package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

const templateStr = `
<html>
<head>
<title>Hi</title>
</head>
<body>
{{if .}}
Hello
    {{range .}}
            {{.}}
	{{end}}
	{{end}}
<form action="/" name=f method="POST">
<input maxLength=1024 size=70 name=s value="" title="Hi friend">
<input maxLength=1024 size=70 name=t value="" title="Hi friend">
<input type=submit value="Say hello" name=hi>
</form>
</body>
</html>
`

var templ = template.Must(template.New("hi").Parse(templateStr))

func Get(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	templ.Execute(w, "")
}

func Post(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")

	data := []string{
		r.PostFormValue("s"),
		r.PostFormValue("t"),
	}
	templ.Execute(w, data)
}

func main() {
	fmt.Println("Server started")
	router := httprouter.New()
	router.GET("/", Get)
	router.POST("/", Post)
	err := http.ListenAndServe("127.0.0.1:80", router)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
