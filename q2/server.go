package main

import (
	"log"
	"net/http"
	"treasure2018-pre/answers/koukyo1994/go/q2/get"
	"strings"
	"encoding/json"
	"fmt"
)

type List struct {
	url string
	title string
	description string
}

func main() {
	http.HandleFunc("/", getHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		http.Error(w, "url not specified", http.StatusBadRequest)
		return
	}

	urlList := strings.Split(url, ";+")
	for _, url := range urlList {
		fmt.Printf("%s\n", url)
	}
	var pageList  = []get.Page{}
	for _, urlStr := range urlList {
		p, err := get.Get(urlStr)
		if err != nil {
			http.Error(w, "request failed", http.StatusInternalServerError)
		} else {
			pageList = append(pageList, *p)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	var list = []List{}
	for i, page := range pageList {
		list = append(list, List{
			urlList[i],
			page.Title,
			page.Description,
		})
	}
	byte, err := json.Marshal(list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(byte)
}
