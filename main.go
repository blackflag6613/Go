package main

import (
    "fmt"
    "net/http"
)

func main(){
    http.HandleFunc("/", Hello)
    http.ListenAndServe(":9090", nil)
}

func Hello(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello world %s!", r.URL.Path[1:])
}