package server

import (
	"fmt"
	"log"
	"net/http"
)

var requestMap = map[string]func(w http.ResponseWriter, req *http.Request){
	"GET /hello":            hello,
	"POST /user":            getUserDetails,
	"POST /user-vaccine":    getUserVaccineDetails,
	"POST /country-vaccine": getRequiredVaccineDetails,
}

func Serve() {
	http.HandleFunc("/", genericHandler)
	log.Println("Starting server on port 8090")
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
	log.Println("Started server on port 8090")
}

func genericHandler(w http.ResponseWriter, req *http.Request) {
	url := req.URL.Path
	method := req.Method
	fmt.Println(fmt.Sprintf("%s %s", method, url))
	requestMap[fmt.Sprintf("%s %s", method, url)](w, req)
}
