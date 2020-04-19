package main

import (
	"crypto/tls"
	"html/template"
	"log"
	"net/http"
)

// Page - what data appears on the page
type Page struct {
	Cookie             string
	IPAddress          string
	TLSState           *tls.ConnectionState
	SuperCookie        string
	BrowserFingerprint string
}

func handleView(w http.ResponseWriter, r *http.Request) {
	v := Page{
		IPAddress: r.RemoteAddr,
		TLSState:  r.TLS,
	}
	t, _ := template.ParseFiles("templates/index.html")
	err := t.Execute(w, v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", handleView)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
