package main

import(
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w,"404 not found",http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "GET method is not supported", http.StatusNotFound)
		return
	}

	if err:= r.ParseForm(); err !=nil {
		fmt.Fprintf(w,"Parse form error: %v",err)
		return
	}
	fmt.Fprintf(w, "POST request successfull\n")
	name:= r.FormValue("name")
	address:= r.FormValue("address")
	fmt.Fprintf(w,"Name: %s \n",name)
	fmt.Fprintf(w,"Address: %s \n", address)
}


func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w,"404 not found",http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Post method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprint(w, "hello")
}

func main() {
	fileServer:= http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	
	fmt.Println("Starting server at port 8080")

	if err:= http.ListenAndServe(":8080",nil); err!=nil {
		log.Fatal(err)
	}
}