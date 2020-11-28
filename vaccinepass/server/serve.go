package server

import (
	"net/http"
)

func Serve() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/user", getUserDetails)
	http.HandleFunc("/user-vaccine", getUserVaccineDetails)
	http.HandleFunc("/country-vaccine", getRequiredVaccineDetails)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
}
