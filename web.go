package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
    "log"
)

func routeIndex(w http.ResponseWriter, r *http.Request) {
    // parse arguments
    r.ParseForm()

    // print form information in server side
    fmt.Println(r.Form)
    fmt.Println("path", r.URL.Path)
    for k, v := range r.Form {
        fmt.Println("key:", k)
        arr := strings.Split(strings.Join(v, ""), ",")
        fmt.Println("val:", arr)
    }
    
    http.ServeFile(w, r, r.URL.Path[1:])
}

func routeDetails(w http.ResponseWriter, r *http.Request) {
    res, err := http.Get("https://dweet.io/get/latest/dweet/for/gmx-elf-ping")
    if err != nil {
        log.Fatal(err)
    }
    json, err := ioutil.ReadAll(res.Body)
    if err != nil {
        log.Fatal(err)
    }
    defer res.Body.Close()

    // send to client
    fmt.Fprintf(w, string(json[:]))
}

func main() {
    // Routers
    http.HandleFunc("/", routeIndex)
    http.HandleFunc("/details", routeDetails)

    // set port
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
