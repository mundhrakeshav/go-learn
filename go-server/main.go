package main

import (
	"fmt"
	"log"
	"net/http"
) 

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err!=nil {
		fmt.Println(w, "ParseForm() err", err)
		return
	}
	fmt.Fprintln(w, "Post req success");
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Println("TEST")
	fmt.Fprintln(w, name);
	fmt.Fprintln(w, address);
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Test1");
	if r.URL.Path != "/hello" {
		http.Error(w, "404", http.StatusNotFound)
		fmt.Println("Test2");
		return
	}
	if r.Method != "GET" {
		fmt.Println("Test2");
		http.Error(w, "Not supported", http.StatusBadGateway);
		return
	}
	fmt.Fprintf(w, "Hello")
}

func main() {
	fileServer := http.FileServer((http.Dir("./static")))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Server srarted at port 8080");
	if err := http.ListenAndServe(":8080", nil); err!=nil{
		log.Fatal(err)
	}
}