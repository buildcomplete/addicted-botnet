package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

var uMap map[string]int

func push(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Write([]byte("hello world"))
		return
	}
		
	add := r.URL.Path[1:]
	uMap[add]= uMap[add] + 1

	fmt.Println("U", r.Header["U"])

	w.Header().Set("Content-Type", "application/json")
	byt, _ := json.Marshal(uMap)
	w.Write(byt)
}


func main() {
	uMap = make(map[string]int)

	http.HandleFunc("/p/", push)

	http.ListenAndServe(":8080", nil)
	
}
