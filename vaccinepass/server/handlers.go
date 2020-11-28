package server

import (
	"fmt"
	"net/http"

	"github.com/parijatpurohit/sidecar-sql/vaccinepass/server/transport"

	"github.com/parijatpurohit/sidecar-sql/vaccinepass/logic"
	"github.com/parijatpurohit/sidecar-sql/vaccinepass/utils/server"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func getUserDetails(w http.ResponseWriter, req *http.Request) {
	reqData := &transport.GetUserDetailsRequest{}

	err := server.ReadReqBody(req, reqData)
	if err != nil {
		server.WriteResponseToServer(w, server.CreateErrMessage(err))
		return
	}

	res, err := logic.GetUserDetails(reqData)
	if err != nil {
		server.WriteResponseToServer(w, server.CreateErrMessage(err))
		return
	}

	server.WriteResponseToServer(w, res)
}

func getUserVaccineDetails(w http.ResponseWriter, req *http.Request) {
	reqData := &transport.GetUserVaccineDetailsRequest{}

	err := server.ReadReqBody(req, reqData)
	if err != nil {
		server.WriteResponseToServer(w, server.CreateErrMessage(err))
		return
	}

	res, err := logic.GetUserVaccineDetails(reqData)
	if err != nil {
		server.WriteResponseToServer(w, server.CreateErrMessage(err))
		return
	}

	server.WriteResponseToServer(w, res)
}

func getRequiredVaccineDetails(w http.ResponseWriter, req *http.Request) {
	reqData := &transport.GetRequiredVaccineRequest{}
	err := server.ReadReqBody(req, reqData)
	if err != nil {
		server.WriteResponseToServer(w, server.CreateErrMessage(err))
		return
	}

	res, err := logic.GetRequiredVaccines(reqData)
	if err != nil {
		server.WriteResponseToServer(w, server.CreateErrMessage(err))
		return
	}

	server.WriteResponseToServer(w, res)
}
