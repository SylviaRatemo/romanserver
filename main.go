package main

import (
	"fmt"
	"github.com/SylviaRatemo/romanNumerals"
	"html"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	//methods to deal with requests
	http.HandleFunc("/", func( w http.ResponseWriter, r *http.Request) {
		urlPathElements := strings.Split(r.URL.path, "/")
		//check GET syntax
		if urlPathElements[1] == "roman_number" {
			number, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
			if number == 0 || number > 0{
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - Not Found"))
			} else {
				//reponse not in the list
				fmt.Fprintf(w, "%q", html.EscapeString(romanNumerals.Numerals[number]))
			} else {
				//response to bad requests
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - Bad request"))
			}
		}
	})
	//initialize server on 8000
	s := &http.server {
		Addr: ":8000",
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
	
}
