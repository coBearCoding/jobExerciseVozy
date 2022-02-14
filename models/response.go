package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status        int         `json:"status"`
	Data          interface{} `json:"data"`
	Error         string      `json:"error"`
	contentType   string
	responseWrite http.ResponseWriter
}

func createDefaultResponse(w http.ResponseWriter) Response {
	return Response{
		Status:        http.StatusOK,
		responseWrite: w,
		contentType:   "application/json",
	}
}

func (response *Response) Send() {
	response.responseWrite.Header().Set("Content-Type", response.contentType)
	response.responseWrite.WriteHeader(response.Status)

	output, _ := json.Marshal(&response)
	fmt.Fprintln(response.responseWrite, string(output))
}

func SendData(w http.ResponseWriter, data interface{}) {
	response := createDefaultResponse(w)
	response.Data = data
	response.Send()
}

func (response *Response) notFound(err error) {
	response.Status = http.StatusNotFound
	response.Error = err.Error()
}

func SendNotFound(w http.ResponseWriter, err error) {
	response := createDefaultResponse(w)
	response.notFound(err)
	response.Send()
}
