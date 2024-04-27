package main

import (
	"log"
	"net/http"
	"html/template"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		t:=template.Must(template.ParseFiles("index.html"));
		t.Execute(w,nil);
	}

	http.HandleFunc("/", helloHandler);

	log.Fatal(http.ListenAndServe(":8000", nil));
}