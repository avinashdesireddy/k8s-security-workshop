package main

import (
  "html/template"
  "log"
  "net/http"
  "os"
)

type PageVariables struct {
  Color string
  Text string
}

func main() {
  http.HandleFunc("/", HomePage)
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomePage(w http.ResponseWriter, r *http.Request){
  HomePageVars := PageVariables{
    Color: os.Getenv("COLOR"),
    Text: "Hello...",
  }
  t, err := template.ParseFiles("homepage.html")
  if err != nil {
    log.Print("Template parsing error: ", err)
  }
  err = t.Execute(w, HomePageVars)
  if err != nil {
    log.Print("Template executing error: ", err)
  }
}
