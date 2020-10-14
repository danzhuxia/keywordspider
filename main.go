package main

import (
	"log"
	"net/http"

	"github.com/danzhuxia/keywordspider/src/spider"
)

func main() {

	log.Println("服务器开启...")
	http.HandleFunc("/test", spider.GetResult)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
