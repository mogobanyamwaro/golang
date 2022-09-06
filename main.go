package main
 import (
	 "fmt"
	 "log"
	 "net/http"
 )

func formHandler(w http.ResponseWriter, r *http.Request) {
if err:= r.ParseForm(); err !=nil{
	fmt.Fprintf(w, "ParseForm() err: %v", err)
	log.Printf("ParseForm() error: %v", err)
	return
}

if r.URL.Path != "/form"{
	http.Error(w, "404 not found.", http.StatusNotFound)
	return
}

if r.Method != "POST"{
	http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
	return
}
fmt.Fprintf(w, "Post request success!")
name := r.FormValue("name")
email := r.FormValue("email")
fmt.Fprintf(w, "Name: %s\n", name)
fmt.Fprintf(w, "Email: %s", email)
 }


func helloHandler(w http.ResponseWriter, r *http.Request) {
if r.URL.Path != "/hello"{
	http.Error(w, "404 not found.", http.StatusNotFound)
	return
}

if r.Method != "GET"{
	http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
	return
}
fmt.Fprintf(w, "Hello mogoba nyamwaro!")
 }


 func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server on port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
 }