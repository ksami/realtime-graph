package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "log"
)

func main() {
    // Routers
    http.HandleFunc("/", routeIndex)
    http.HandleFunc("/details", routeDetails)

    // set ip and port
    bind := fmt.Sprintf("%s:%s", os.Getenv("OPENSHIFT_GO_IP"), os.Getenv("OPENSHIFT_GO_PORT"))
    fmt.Printf("listening on %s...", bind)
    err := http.ListenAndServe(bind, nil)
    if err != nil {
        panic(err)
    }
}

func routeIndex(w http.ResponseWriter, r *http.Request) {
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