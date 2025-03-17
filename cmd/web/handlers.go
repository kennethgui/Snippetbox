package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Server", "Go") 
 
  files := []string{
    "./ui/html/base.tmpl",
    "./ui/html/partials/nav.tmpl",
    "./ui/html/pages/home.tmpl",
  }

  //template.ParseFiles() function to read the templare file into a
  //template set. If there's an error, we log the detailed error message, use the
  //http.Error() function to send an internal server error response to the user
  //and then return from the handler so no subsequent code is executed.
  ts, err := template.ParseFiles(files...)
  if err != nil {
    log.Print(err.Error())
    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    return
  }
  
  //Use the ExcecuteTemplate() method to write the content of the "base
  err = ts.ExecuteTemplate(w, "base", nil)
  if err != nil {
    log.Print(err.Error())
    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
  }
}

func snippetView(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.PathValue("id"))
    if err != nil || id < 1 {
        http.NotFound(w, r)
        return
    }

    fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Display a form for creating a new snippet..."))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("Save a new snippet..."))
}
