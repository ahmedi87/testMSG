package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	msgURL = "https://graph.facebook.com/v20.0/466488716541712/messages"
)

func HandleGetRequest(r *http.Request, w http.ResponseWriter) {
	req := queryParams(r)
	writeLog(createKeyValuePairs(req), "Get")
	challenge := fmt.Sprint(req["hub.challenge"])
	mod := req["hub.mode"]
	fmt.Println("My mod", mod, "my challenge", challenge)
	if mod == "subscribe" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(challenge))
	} else {
		fmt.Fprint(w, "success")
	}
}

func HandlePostRequest(r *http.Request, w http.ResponseWriter) {
	req, _, err := ParseRequest(r)
	if err != nil {
		SendResponse(w, false, err.Error(), 500, nil)
		return
	}
	writeLog(createKeyValuePairs(req), "POST")
	var notifyReq NotificationReq
	data, _ := json.Marshal(req)
	json.Unmarshal(data, &notifyReq)
	fmt.Println("My notification", notifyReq)
	obj := notifyReq.Entry.Changes[0].Value.Messages[0]
	message := map[string]interface{}{"sender": obj.From,
		"message_id": obj.ID, "text": obj.Text.Body}
	writeLog(createKeyValuePairs(message), "messages")
	fmt.Println("My struct:", notifyReq, message)
	SendResponse(w, true, "message received!!", 0, req)
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	req, _, err := ParseRequest(r)
	if err != nil {
		SendResponse(w, false, err.Error(), 500, nil)
		return
	}
	writeLog(createKeyValuePairs(req), "POST")
	var sendReq SendMessageRequest
	data, _ := json.Marshal(req)
	json.Unmarshal(data, &sendReq)
	var request interface{}
	if sendReq.Mode == 1 {
		request = TextMessageRequest{
			MessagingProduct: "whatsapp",
			RecipientType:    "individual",
			To:               sendReq.To,
			Type:             "text",
			Text: struct {
				PreviewURL bool   "json:\"preview_url\""
				Body       string "json:\"body\""
			}{
				PreviewURL: false,
				Body:       sendReq.Text,
			},
		}
	} else if sendReq.Mode == 2 {
		request = ReplyRequest{
			MessagingProduct: "whatsapp",
			RecipientType:    "individual",
			To:               sendReq.To,
			Type:             "text",
			Context: struct {
				MessageID string "json:\"message_id\""
			}{
				MessageID: sendReq.ID,
			},
			Text: struct {
				Body string "json:\"body\""
			}{
				Body: sendReq.Text,
			},
		}
	} else {

	}
	resp := CallURLPost(msgURL, request)
	fmt.Println("My response:", resp)
	SendResponse(w, resp.Resp.ErrorCode == 0, "sent", resp.Resp.ErrorCode, nil)
}
