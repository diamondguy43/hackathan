package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	util "github.com/diamondguy43/hackathan/util"
	"github.com/joho/godotenv"
)

const (
	PORT = 9900
)

type P_REQ struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Location  string `json:"location"`
}

func root(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// var res map[string]string
		// decoder := json.NewDecoder(r.Body)
		// err := decoder.Decode(&res)
		// util.Check(err)

		// fmt.Println(r.Body)

		// fmt.Println(res["sDate"])

		var gPost P_REQ
		err := json.NewDecoder(r.Body).Decode(&gPost)
		if err != nil {
			fmt.Println("asdlmnasdf")
			log.Fatal(err)
		}
		defer r.Body.Close()

		fmt.Println(r.Body)

		// json message decoder
		image := util.QRKeyGen(gPost.StartDate, gPost.EndDate, gPost.Location)
		w.Write(image)
		// http.Redirect(w, r, "localhost:9900/qr", http.StatusAccepted)
	default:
		tpl := template.Must(template.ParseGlob("*.html"))

		tpl.ExecuteTemplate(w, "layout", nil)
	}
}

func search(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		http.Redirect(w, r, "/qr", http.StatusSeeOther)
	default:
		tpl := template.Must(template.ParseGlob("*.html"))

		tpl.ExecuteTemplate(w, "layout", nil)
	}
}

func qr(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":

		var gPost P_REQ
		err := json.NewDecoder(r.Body).Decode(&gPost)
		if err != nil {
			fmt.Println("bruh")
			log.Fatal(err)
		}
		defer r.Body.Close()

		fmt.Println(r.Body)

		// json message decoder
		image := util.QRKeyGen(gPost.StartDate, gPost.EndDate, gPost.Location)
		w.Write(image)
	default:
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	util.GetTransaction(os.Getenv("walletAddr"))

	port := strconv.Itoa(PORT)

	http.HandleFunc("/", root)
	http.HandleFunc("/search", search)
	http.HandleFunc("/qr", qr)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// running on server
	fmt.Printf("Started web server at:  http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
