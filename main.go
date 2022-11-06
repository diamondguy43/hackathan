package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	util "github.com/diamondguy43/hackathan/util"
)

const (
	PORT = 8080
)

type P_REQ struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Location  string `json:"location"`
}

func root(w http.ResponseWriter, r *http.Request) {

}

// func info(w http.ResponseWriter, r *http.Request) {

// }

// func mapper(w http.ResponseWriter, r *http.Request) {

// }

func qr(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var gPost P_REQ
		err := json.NewDecoder(r.Body).Decode(&gPost)
		if err != nil {
			log.Fatal(err)
		}
		defer r.Body.Close()

		// json message decoder
		image := util.QRKeyGen(gPost.StartDate, gPost.EndDate, gPost.Location)
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

	// running on server
	fmt.Printf("Started web server at:  http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
