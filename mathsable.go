package main

import (
  "net/http"
  "html/template"
  "log"
  "io"
  "time"
  "os"
)

const STATIC_URL string = "/www/static/"
const STATIC_ROOT string = "www/static/"

type Context struct {
  Title string `json:"title"`
  Static string `json:"static"`
}

var templates = template.Must(template.ParseGlob("www/templates/*"))

func handler(w http.ResponseWriter, r *http.Request) {
  context := Context{Title: "Hello"}
  render(w, "layout.html", context)
}

func render(w http.ResponseWriter, template string, context Context) {
  context.Static = STATIC_URL
  err := templates.ExecuteTemplate(w, template, context)
  if err != nil {
    log.Print("Failed to render template: ", err)
  }
}

func StaticHandler(w http.ResponseWriter, req *http.Request) {
  static_file := req.URL.Path[len(STATIC_URL):]
  if len(static_file) != 0 {
    f, err := http.Dir(STATIC_ROOT).Open(static_file)
    if err == nil {
      content := io.ReadSeeker(f)
      http.ServeContent(w, req, static_file, time.Now(), content)
      return
    }
  }
  http.NotFound(w, req)
}

func main() {
  http.HandleFunc("/", handler)
  http.HandleFunc(STATIC_URL, StaticHandler)
  http.ListenAndServe(":" + os.Getenv("PORT"), nil)
}
