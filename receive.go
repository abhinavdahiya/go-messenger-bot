package messengerapi

type Response struct {
	Object  string  `json:"object"`
	Entries []Entry `json:"entry"`
}

type Entry struct {
	PageID    int64      `json:"id"`
	Time      int64      `json:"time"`
	Messaging []Callback `json:"messaging"`
}

type Callback struct {
	Sender    User     `json:"sender"`
	Recipient Page     `json:"recipient"`
	Timestamp int64    `json:"timestamp"`
	Optin     Optin    `json:"optin"`
	Message   Message  `json:"message,omitempty"`
	Postback  Postback `json:"postback,omitempty"`
	Delivery  Delivery `json:"delivery,omitempty"`
}

type User struct {
	ID          int64  `json:"id,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

type Page struct {
	ID int64 `json:"id"`
}

type Optin struct {
	Ref string `json:"ref"`
}

type Message struct {
	MID         string       `json:"mid"`
	Seq         int64        `json:"seq"`
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

type Attachment struct {
	Type    string        `json:"type"`
	Payload AttachPayload `json:"payload"`
}

type AttachPayload struct {
	URL string `json:"url"`
}

type Delivery struct {
	MIDs      []string `json:"mids"`
	Watermark int64    `json:"watermark"`
	Seq       int64    `json:"seq"`
}

type Postback struct {
	Payload string `json:"payload"`
}
