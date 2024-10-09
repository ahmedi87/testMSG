package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/motaz/codeutils"
)

func main() {
	http.HandleFunc("/webhook", logMSG)
	fmt.Println("Listening On: 9099")
	http.ListenAndServe(":9099", nil)

}

func logMSG(w http.ResponseWriter, r *http.Request) {
	fmt.Println("My request to API LOG")
	if r.Method == http.MethodGet {
		req := queryParams(r)
		fmt.Println("My req", req)
		challenge := req["hub.challenge"]
		mod := req["hub.mod"]
		fmt.Println("My mod", mod, "my challenge", challenge)
		if mod == "subscribe" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(challenge))
		} else {
			fmt.Fprint(w, "success")
		}
		return
	} else {
		req, _, err := ParseRequest(r)
		if err != nil {
			fmt.Fprint(w, "error")
			return
		}
		fmt.Println("My request:", createKeyValuePairs(req))
		// go writeLog(createKeyValuePairs(req), "logs")
	}
	fmt.Fprint(w, "success")
}

func queryParams(r *http.Request) (res map[string]string) {
	res = make(map[string]string)
	for k := range r.URL.Query() {

		val := r.URL.Query()[k]
		res[k] = val[0]
	}
	return
}

func createKeyValuePairs(m map[string]interface{}) string {
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
