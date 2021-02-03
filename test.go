package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JustinSDK/goexample"
)

type Person struct {
	name string
	age  uint
}

var data = "58444"

func httpc() {
	http.HandleFunc("/root", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write([]byte(data))
		//	w.Write([]byte())
	})

}

func main() {

	httpc()
	fmt.Println("wake server")
	log.Println("server on55")
	log.Fatal(http.ListenAndServe(":4000", nil))
	goexample.Hi()
	goexample.Hello()
}
