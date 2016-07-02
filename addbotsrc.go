package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/base64"
	"strings"
)

// global 'Database' of URL's
var uMap map[string]int

// number of times to return an url after each hit
var increment = 10

// number of urls to transmit in each retr
var sendcount = 20

func vers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("1"))
}

func retr(w http.ResponseWriter, r *http.Request) {
	if len(uMap) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	i := sendcount
	for k,v := range uMap {
		w.Write([]byte(k))
		w.Write([]byte("\n"))
		
		if (v == 1) {
			delete(uMap,k)
			fmt.Println(len(uMap))
		} else {
			uMap[k]=v-1
		}

		i--
		if i==0 {
			return
		}
	}
}

func push(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		
		// check if expected data is there...
		if len(r.Header["U"]) == 0 {
			return;
		}

		// U is a space sepperated list contaning the urls base64 encoded
		// extract add add to uMap

		data := r.Header["U"][0]
		for _, f := range strings.Fields(data) {
			if decoded, errD := base64.StdEncoding.DecodeString(f); errD == nil {
				decodedStr := string(decoded)
				uMap[decodedStr] = uMap[decodedStr] + increment
			} else {
				fmt.Print(errD)
			}
		}
		
		w.WriteHeader(http.StatusNoContent)
		fmt.Println(len(uMap))
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}


func main() {
	uMap = make(map[string]int)
	http.HandleFunc("/v", vers)
	http.HandleFunc("/p", push)
	http.HandleFunc("/r", retr)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
