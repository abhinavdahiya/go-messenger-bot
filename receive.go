package mbotapi

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
	Sender    User          `json:"sender"`
	Recipient Page          `json:"recipient"`
	Timestamp int64         `json:"timestamp"`
	Optin     InputOptin    `json:"optin"`
	Message   InputMessage  `json:"message,omitempty"`
	Postback  InputPostback `json:"postback,omitempty"`
	Delivery  InputDelivery `json:"delivery,omitempty"`
}

type User struct {
	ID          int64  `json:"id,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

type Page struct {
	ID int64 `json:"id"`
}

type InputOptin struct {
	Ref string `json:"ref"`
}

type InputMessage struct {
	MID         string            `json:"mid"`
	Seq         int64             `json:"seq"`
	Text        string            `json:"text"`
	Attachments []InputAttachment `json:"attachments,omitempty"`
}

type InputAttachment struct {
	Type    string             `json:"type"`
	Payload InputAttachPayload `json:"payload"`
}

type InputAttachPayload struct {
	URL string `json:"url"`
}

type InputDelivery struct {
	MIDs      []string `json:"mids"`
	Watermark int64    `json:"watermark"`
	Seq       int64    `json:"seq"`
}

type InputPostback struct {
	Payload string `json:"payload"`
}
