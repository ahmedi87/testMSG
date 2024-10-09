package main

type NotificationReq struct {
	Object string `json:"object"`
	Entry  []struct {
		Changes []Change `json:"changes"`
	} `json:"entry"`
}

type Change struct {
	Field string `json:"field"`
	Value struct {
		MessagingProduct string `json:"messaging_product"`
		Metadata         struct {
			DisplayPhoneNumber string `json:"display_phone_number"`
			PhoneNumberID      string `json:"phone_number_id"`
		} `json:"metadata"`
		Contacts []struct {
			Profile struct {
				Name string `json:"name"`
			} `json:"profile"`
			WaID string `json:"wa_id"`
		} `json:"contacts"`
		Messages []struct {
			From      string `json:"from"`
			ID        string `json:"id"`
			Timestamp string `json:"timestamp"`
			Type      string `json:"type"`
			Text      struct {
				Body string `json:"body"`
			} `json:"text"`
		} `json:"messages"`
	} `json:"value"`
}

type TextMessageRequest struct {
	MessagingProduct string `json:"messaging_product"`
	RecipientType    string `json:"recipient_type"`
	To               string `json:"to"`
	Type             string `json:"type"`
	Text             struct {
		PreviewURL bool   `json:"preview_url"`
		Body       string `json:"body"`
	} `json:"text"`
}

type MessageResponse struct {
	MessagingProduct string `json:"messaging_product"`
	Contacts         []struct {
		Input string `json:"input"`
		WaID  string `json:"wa_id"`
	} `json:"contacts"`
	Messages []struct {
		ID            string `json:"id"`
		MessageStatus string `json:"message_status"`
	} `json:"messages"`
	Resp Response
}

type ReplyRequest struct {
	MessagingProduct string `json:"messaging_product"`
	RecipientType    string `json:"recipient_type"`
	To               string `json:"to"`
	Context          struct {
		MessageID string `json:"message_id"`
	} `json:"context"`
	Type string `json:"type"`
	Text struct {
		Body string `json:"body"`
	} `json:"text"`
}

type Response struct {
	Success   bool                   `json:"success"`
	ErrorCode int                    `json:"errorcode"`
	Message   string                 `json:"message"`
	Result    map[string]interface{} `json:"result"`
}

type SendMessageRequest struct {
	Mode int    `json:"mode"`
	To   string `json:"to"`
	ID   string `json:"id"`
	Text string `json:"text"`
}
