package utils

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/apex/gateway"
	"github.com/joho/godotenv"
)

func SetDefaultHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
}

func CreateApiResponse(v interface{}) []byte {
	var response []byte
	if v != "" {
		jsonBody, err := json.Marshal(v)
		if err != nil {
			log.Fatalf("An error occurred in JSON marshal. Err: %s", err)
		}
		response = jsonBody
	}

	return response
}

func ParseRequestBody(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func ServeFunction(url string, handler func(http.ResponseWriter, *http.Request)) {
	port := flag.Int("port", -1, "specify a port")
	flag.Parse()
	listener := gateway.ListenAndServe
	addr := ""
	if *port != -1 {
		err := godotenv.Load()
		if err != nil {
			log.Print("Failed to load .env file")
		}
		addr = fmt.Sprintf(":%d", *port)
		listener = http.ListenAndServe
		http.Handle("/", http.FileServer(http.Dir("./public")))
	}
	http.HandleFunc(url, handler)

	log.Printf("Function `%s` running on %s...", url, addr)
	log.Fatal(listener(addr, nil))
}
