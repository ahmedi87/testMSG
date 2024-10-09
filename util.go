package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/motaz/codeutils"
)

func queryParams(r *http.Request) (res map[string]interface{}) {
	res = make(map[string]interface{})
	for k := range r.URL.Query() {

		val := r.URL.Query()[k]
		res[k] = val[0]
	}
	return
}

func createKeyValuePairs(m map[string]interface{}) string {
	fmt.Println("MY createPairs", m)
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s:\"%s\", ", key, value)
	}
	return b.String()
}

func writeLog(event string, name string) {

	if name == "" {
		fmt.Println("Error :", event)
		name = "log/Error"
	}
	codeutils.WriteToLog(event, name)
}

// ParseRequest ...
func ParseRequest(r *http.Request) (request map[string]interface{}, body []byte, err error) {
	request = make(map[string]interface{})
	body, err = io.ReadAll(r.Body)

	if err != nil {
		codeutils.WriteToLog("Error in processRequest reading request: "+err.Error(), "")
		return
	}
	err = json.Unmarshal(body, &request)
	if err != nil {
		codeutils.WriteToLog("Error in processRequest unmarshal: "+err.Error(), "")
		return
	}
	return
}

func SendResponse(w http.ResponseWriter, success bool, msg string, errorcode int, result map[string]interface{}) {
	var resp Response
	resp.Success, resp.ErrorCode = success, errorcode
	resp.Message, resp.Result = msg, result
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	output, _ := json.Marshal(resp)
	w.Write(output)
}

func CallURLPost(url string, req interface{}) (response MessageResponse) {
	jsonReq, _ := json.Marshal(req)
	client := &http.Client{Timeout: 60 * time.Second}
	request, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Set("Authorization", "Bearer EAAO4Th5gYfEBOwbHZAhvLGmU8WZAExayMGQgZAUZCBiBXiAjTVMvW4EwYbCmYgc9h1MpAEPRXT6Mqc8qZBoIG1P771qTFUS4Hk3GZAVptCK7czgzjNrIpzKTs8FnZCl6xv3L2AlsyajeEppvrJ8o4MccHFg5Yyba7WVv2GpCJckFvD3Fcdg4rImxazrjMUZBa1pOG07X5vPSnZBCowPMw9YCCHvZCduIMZD")
	request.Close = true
	resp, err := client.Do(request)
	if err != nil {
		writeLog("Error in callURLPost "+err.Error(), "")
		response.Resp.Message = err.Error()
		response.Resp.ErrorCode = 500
		return
	}
	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &response)
	return
}
