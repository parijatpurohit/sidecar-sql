package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/parijatpurohit/sidecar-sql/vaccinepass/server/transport"
)

func ReadReqBody(req *http.Request, obj interface{}) error {
	reqBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(reqBytes, obj)
	if err != nil {
		return err
	}
	return nil
}

func CreateErrMessage(err error) *transport.ErrResponse {
	return &transport.ErrResponse{Message: err.Error()}
}

func createErrStr(err error) []byte {
	res, _ := json.Marshal(CreateErrMessage(err))
	return res
}

func WriteResponseToServer(w http.ResponseWriter, res interface{}) {
	resData, err := json.Marshal(res)
	if err != nil {
		writeResponse(w, createErrStr(err))
	}
	_, err = w.Write(resData)
	if err != nil {
		log.Println("error while writing response", err.Error())
	}
}

func writeResponse(w http.ResponseWriter, res []byte) {
	_, err := w.Write(res)
	if err != nil {
		log.Println("error while writing response", err.Error())
	}
}
