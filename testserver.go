package main

import (
//    "fmt"
    "net/http"
//    "net/http/httputil"
)

func main() {
    http.HandleFunc("/send", func(response http.ResponseWriter, request *http.Request) {
//        dump, _ := httputil.DumpRequest(request, true)
//	fmt.Printf("%q\n",dump)
	request.ParseForm()
        result := request.Method + " "
        if request.Method == "POST" {
            result += request.PostFormValue("value")
        } else if request.Method == "GET" {
            result += request.FormValue("value")
        }
//        fmt.Println(result)
//        fmt.Println(request.Header)
        response.Write([]byte(result))
    })
    http.ListenAndServe(":8080", nil)
}
