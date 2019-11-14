package main

import (
	"github.com/astaxie/beego"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("hello"))
}

func main() {
	beego.Run()
	//mux := http.NewServeMux()
	//mux.HandleFunc("/", helloHandler)
	//_ = http.ListenAndServe(":12345", mux)
}
