package homecontroller

import (
  "net/http"
  "html/template"
  "log"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
  //membuat template dari html
  temp, err := template.ParseFiles("views/home/index.html")
  if err != nil {
    log.Fatal(err)
  }

  //mengekseskusi file template
  err = temp.Execute(w, nil)
  if err != nil {
    log.Fatal(err)
  }
}