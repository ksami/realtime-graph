package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
    "log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()  // parse arguments, you have to call this by yourself
    fmt.Println(r.Form)  // print form information in server side
    fmt.Println("path", r.URL.Path)
    for k, v := range r.Form {
        fmt.Println("key:", k)
        arr := strings.Split(strings.Join(v, ""), ",")
        fmt.Println("val:", arr)
    }
    //fmt.Fprintf(w, "Hello astaxie!") // send data to client side
    http.ServeFile(w, r, r.URL.Path[1:])
}

func giveDetails(w http.ResponseWriter, r *http.Request) {
    res, err := http.Get("https://dweet.io/get/latest/dweet/for/gmx-elf-ping")
    if err != nil {
        log.Fatal(err)
    }
    json, err := ioutil.ReadAll(res.Body)
    if err != nil {
        log.Fatal(err)
    }
    defer res.Body.Close()
    
    fmt.Println("path", r.URL.Path)
    fmt.Fprintf(w, string(json[:]))
}

func main() {
    http.HandleFunc("/", sayhelloName) // set router
    http.HandleFunc("/details", giveDetails)
    err := http.ListenAndServe(":9090", nil) // set listen port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
