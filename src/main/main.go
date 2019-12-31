package main

import (
	"log"
	"net/http"
	"encoding/json"	
	
	"io/ioutil"
	"time"
	"bytes"
	"fmt"
	"./report"
)


func main() {
	log.SetFlags(log.Lshortfile)
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		modtime := time.Now()

		content, err := ioutil.ReadFile("public/hello.pdf")
		if err != nil {
			log.Println("file" + " not read properly..")
		}

		// ServeContent uses the name for mime detection
		const name = "GPA-Report-Blinism.pdf"

		// tell the browser the returned content should be downloaded
		w.Header().Add("Content-Disposition", "Attachment")

		http.ServeContent(w, req, name, modtime, bytes.NewReader(content))
	})

	http.HandleFunc("/test", test)

	log.Fatal(http.ListenAndServe(":8081", nil))

}
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func test(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	switch r.Method {
		case "GET":
			fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
		case "POST":
			var rpd report.ReportData
			err := json.NewDecoder(r.Body).Decode(&rpd)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			report.GenReport(rpd)
		default:
			log.Println(w, "Sorry, only GET and POST methods are supported.")
	}
}
