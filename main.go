package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	util "github.com/diamondguy43/hackathan/util"
)

const (
	PORT = 8080
)

func root(w http.ResponseWriter, r *http.Request) {

}

// func info(w http.ResponseWriter, r *http.Request) {

// }

// func mapper(w http.ResponseWriter, r *http.Request) {

// }

func qr(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// json message decoder
		image := util.QRKeyGen("yes", "yes", "yes")
		w.Write(image)
	default:
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}

func main() {
	port := strconv.Itoa(PORT)

	http.HandleFunc("/", root)
	http.HandleFunc("/qr", qr)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	fmt.Printf("Started web server at:  http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
