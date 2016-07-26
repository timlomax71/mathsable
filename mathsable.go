package main

import (
  "net/http"
  "html/template"
)

func handler(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("./tmpl/layout.html", "./tmpl/head.html", "./tmpl/login.html", "./tmpl/naviagtion.html")
  t.ExecuteTemplate(w, "layout", "")
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8081", nil)
}
