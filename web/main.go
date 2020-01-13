package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	sfeir  = "sfeir"
	origin = "/"
)

var (
	tmpl  *template.Template
	sfURL string = ""
)

func init() {
	var err error
	tmpl = template.New("gopherTemplate")
	tmpl, err = tmpl.ParseFiles("template/gopher.tmpl.html")
	if err != nil {
		log.Printf("Error parsing template %v \n", err)
	}

	sfURL = os.Getenv("GARLAND_ACTION_URL")
	if sfURL == "" {
		log.Println("Error while reading garland action URL")
		return
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("entering root handler")

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err := tmpl.ExecuteTemplate(w, "gopher.tmpl.html", nil)
	if err != nil {
		log.Printf("Error executing template %v \n", err)
	}
}

func sendHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("entering send handler")

	if sfURL == "" {
		log.Printf("Unable to execute command, wrong configuration \n")
		http.Redirect(w, r, origin, http.StatusFound)
		return
	}

	word := r.FormValue("word")

	if word == "" {
		word = sfeir
	}

	word = strings.Replace(word, " ", "", -1)
	word = strings.ToLower(word)

	resp, err := http.Get(sfURL + word)
	if err != nil {
		log.Printf("Error executing command %s %v \n", word, err)
	} else {
		log.Printf("Successfully sent command %s \n", sfURL+word)
	}

	defer func() {
		if resp != nil && resp.Body != nil {
			err := resp.Body.Close()
			if err != nil {
				log.Printf("Error closing command body %v \n", err)
			}
		}
	}()

	http.Redirect(w, r, origin, http.StatusFound)
}

func main() {
	log.Println("servers starting...")

	// seeding random numbers
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/send", sendHandler)

	port := os.Getenv("PORT")
	if port == "" {
		log.Println("Error while reading default HTTP port, using 80080 as default")
		port = "8080"
	}

	log.Printf("launching HTTP on localhost:%s \n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Printf("error starting webserver %v", err)
	}

	log.Println("...servers stopped")
}
