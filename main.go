package main

import (
	"fmt"
	"net/http"
	"simpleapp/controller"
)

func main() {

	controller.Master()
	// Then we create a new server and start listening for incoming requests with the http.ListenAndServe() function, passing in our servemux for it to match requests against as the second parameter.
	fmt.Println("Server serving...")
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":80", nil)

}
