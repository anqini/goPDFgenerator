package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"bytes"
	"fmt"
	"io/ioutil"
	"time"

	"./randStr"
	"./report"
)

// func enableCors(w *http.ResponseWriter) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
// }

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}

func main() {
	// init logger
	f, err := os.OpenFile("text.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	logger := log.New(f, "prefix", log.LstdFlags)
	// define the only route
	http.HandleFunc("/genReport", func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w, r)
		modtime := time.Now()
		// generate random string to store file
		filename := randStr.GenRandFilename()
		// deal with get method and post method
		// log.Println(r.Method)
		switch r.Method {
		case "OPTIONS": // preflight request
			return
		case "GET":
			fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
			logger.Println("GET: server page.")
			return
		case "POST":
			var rpd report.ReportData
			err := json.NewDecoder(r.Body).Decode(&rpd)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			err2 := report.GenReport(rpd, filename, logger)
			if err2 != nil {
				return
			}
			logger.Println("POST: generate pdf for " + rpd.Name + ". phone: " + rpd.Phone)
		default:
			logger.Println("Sorry, only GET and POST methods are supported.")
			return
		}
		content, err := ioutil.ReadFile("public/" + filename)
		if err != nil {
			logger.Println("file" + " not read properly..")
			return
		}

		// ServeContent uses the name for mime detection
		const name = "GPA-Report-Blinism.pdf"

		// tell the browser the returned content should be downloaded
		w.Header().Add("Content-Disposition", "Attachment")
		// w.Header().Set("Access-Control-Allow-Origin", "*")

		http.ServeContent(w, r, name, modtime, bytes.NewReader(content))
	})

	logger.Fatal(http.ListenAndServe(":3023", nil))

}
