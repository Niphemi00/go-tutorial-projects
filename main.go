package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "static/form.html")
		return
	}
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return 
		}
		fmt.Fprintf(w, "POST request successful\n")
		name := r.FormValue("name")
		address := r.FormValue("address")
		fmt.Fprintf(w, "Name: %s \n", name)
		fmt.Fprintf(w, "Address: %s \n", address)
	}


	// if r.Method != "POST" {
	// 	http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
	// 	return
	// }
	// fmt.Fprintf(w, "Welcome!!!")

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	
	if r.URL.Path != "/index" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "static/index.html")

}




func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/index", helloHandler)

	fmt.Printf("starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}