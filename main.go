package main

import (
	"net/http"

	"github.com/danzhuxia/keywordspider/src/spider"
)

func main() {
	http.HandleFunc("/test", spider.GetResult)
	http.ListenAndServe("localhost:8080", nil)
}
