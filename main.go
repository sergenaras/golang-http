package main

import (
	"os"
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "ok\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func shutdown(w http.ResponseWriter, req *http.Request){
	fmt.Println(w,"Bye!")

	os.Exit(3)
}

func main() {

    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)
    http.HandleFunc("/shutdown", shutdown)

    http.ListenAndServe(":8090", nil)
}
