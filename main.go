package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type Context struct {
	Env map[string]string
}

func Index(w http.ResponseWriter, req *http.Request) {
	var env []string
	var context Context
	envs := make(map[string]string)

	env = os.Environ()

	for _, value := range env {
		variable := strings.Split(value, "=") // split by = sign
		envs[variable[0]] = variable[1]
	}

	fmt.Printf("%s, %s, %s\n", req.RemoteAddr, req.Method, req.URL)

	context.Env = envs
	render(w, "index.html", context)
}

func render(w http.ResponseWriter, tmpl string, context Context) {
	tmpl = fmt.Sprintf("templates/%s", tmpl)
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, context)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

func main() {
	http.HandleFunc("/", Index)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
