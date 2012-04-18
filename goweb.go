package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
)

var tmpl *template.Template

func main() {
    http.HandleFunc("/", hello)
    http.Handle("/files", http.StripPrefix("/files", http.FileServer(http.Dir(""))))
    //http.HandleFunc("/world", world)
    http.HandleFunc("/world/", worldDir)
    http.HandleFunc("/reflect", reflectIt)

    go watchTemplates("templates", &tmpl)

    err := http.ListenAndServe(":8081", nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}

func hello(w http.ResponseWriter, req *http.Request) {
    if req.URL.Path != "/" {
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintln(w, "Fail")
        return
    }
    fmt.Fprintln(w, "hello, world!")
}

func world(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(w, "HELLO, WORLD!")
}

func worldDir(w http.ResponseWriter, req *http.Request) {
    splits := strings.Split(req.RequestURI, "/")
    fmt.Fprintln(w, splits)
    fmt.Fprintln(w, "HELLO, WORLD NUMBER " + splits[len(splits)-1] + "!")
}

func reflectIt(w http.ResponseWriter, req *http.Request) {
    m := &MyModel{1, true, 3.14}

    err := tmpl.ExecuteTemplate(w, "mymodel.htmlgo", m)
    if err != nil {
        fmt.Fprintf(w, "Oops: %v", err)
    }
}
