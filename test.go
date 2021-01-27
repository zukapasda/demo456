package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<a href=https://www.youtube.com/?gl=TW&hl=zh-tw >8787878</a>"))
	})
	fmt.Println("Aloha")
	log.Println("server on")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
