package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/motaz/codeutils"
)

func main() {
	http.HandleFunc("/LogWhatsappMSG", logMSG)
	fmt.Println("Listening On: 9099")
	http.ListenAndServe(":9099", nil)

}

func logMSG(w http.ResponseWriter, r *http.Request) {

	call := queryParams(r)
	go writeLog(createKeyValuePairs(call), "calls")
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

func createKeyValuePairs(m map[string]string) string {
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
