package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/webhook", logMSG)
	http.HandleFunc("/SendMSG", SendMessage)
	fmt.Println("Listening On: 9099")
	http.ListenAndServe(":9099", nil)

}

func logMSG(w http.ResponseWriter, r *http.Request) {
	fmt.Println("My request to API LOG")
	if r.Method == http.MethodGet {
		HandleGetRequest(r, w)
	} else {
		HandlePostRequest(r, w)
		// go writeLog(createKeyValuePairs(req), "logs")
	}
}
