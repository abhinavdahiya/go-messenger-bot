package mbotapi

// Payload received by the webhook
// The Object field is always set to `page`
// Contains bacthed entries
type Response struct {
	Object  string  `json:"object"`
	Entries []Entry `json:"entry"`
}

// This defines an Entry in the payload by webhook
type Entry struct {
	PageID    string     `json:"id"`
	Time      int64      `json:"time"`
	Messaging []Callback `json:"messaging"`
}

// This represents the content of the message sent by
// the user
// Various kinds of callbacks from user are -
// OptinCallback
// MessageCallback
// PostbackCallback
// DeliveryCallback
//
//TODO: Create a way to identify the type of callback
type Callback struct {
	Sender    User          `json:"sender"`
	Recipient Page          `json:"recipient"`
	Timestamp int64         `json:"timestamp"`
	Optin     InputOptin    `json:"optin"`
	Message   InputMessage  `json:"message,omitempty"`
	Postback  InputPostback `json:"postback,omitempty"`
	Delivery  InputDelivery `json:"delivery,omitempty"`
}

func (c Callback) IsMessage() bool {
	return !(c.Message.Text == "" && len(c.Message.Attachments) == 0)
}

func (c Callback) IsOptin() bool {
	return !(c.Optin == (InputOptin{}))
}

func (c Callback) IsPostback() bool {
	return !(c.Postback == (InputPostback{}))
}

func (c Callback) IsDelivery() bool {
	return !(len(c.Delivery.MIDs) == 0)
}

// This defines an user
// One of the fields will be set to identify the user
type User struct {
	ID          int64  `json:"id,omitempty,string"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

type Page struct {
	ID int64 `json:"id,string"`
}

// Ref contains the `data-ref` set for message optin for the bot
type InputOptin struct {
	Ref string `json:"ref"`
}

// This represents a Message from user
// If text message only Text field exists
// If media message Attachments fields contains an array of attachmensts sent
type InputMessage struct {
	MID         string            `json:"mid"`
	Seq         int64             `json:"seq"`
	Text        string            `json:"text"`
	Attachments []InputAttachment `json:"attachments,omitempty"`
	QuickReply  InputQRPayload    `json:"quick_reply,omitempty"`
}

//Represents a quick reply payload
type InputQRPayload struct {
	Payload string `json:"payload"`
}

// Represents an attachement
// The types are image/audio/video/location
type InputAttachment struct {
	Type    string             `json:"type"`
	Payload InputAttachPayload `json:"payload"`
}

type InputAttachPayload struct {
	URL    string      `json:"url,omitempty"`
	Coords InputCoords `json:"coordinates,omitempty"`
}

type InputCoords struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

// This contains delivery reports for batch
// of messages(mids)
type InputDelivery struct {
	MIDs      []string `json:"mids"`
	Watermark int64    `json:"watermark"`
	Seq       int64    `json:"seq"`
}

// Represents a postback sent by clicking on Postback Button
type InputPostback struct {
	Payload string `json:"payload"`
}
