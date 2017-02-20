package mbotapi

import "errors"

const (
	RegularNotif = "REGULAR"
	SilentNotif  = "SILENT_PUSH"
	NoNotif      = "NO_PUSH"
	TypingON     = "typing_on"
	TypingOFF    = "typing_off"
	MarkSeen     = "mark_seen"
)

var (
	ErrTitleTooLong         = errors.New("Template Title exceeds the 25 character limit")
	ErrSubtitleTooLong      = errors.New("Template Subtitle exceeds the 80 character limit")
	ErrButtonsLimitExceeded = errors.New("Max 3 buttons allowed on GenericTemplate")
	ErrBubblesLimitExceeded = errors.New("Max 10 bubbles allowed on GenericTemplate")
)

const (
	APIEndpoint = "https://graph.facebook.com/v2.6/me/messages?access_token=%s"
)

type Request struct {
	Recipient User    `json:"recipient"`
	Message   Message `json:"message,omitempty"`
	Action    Action  `json:"sender_action,omitempty"`
	NotifType string  `json:"notification_type"`
}

type Action string

type Message struct {
	Text       string      `json:"text,omitempty"`
	Attachment *Attachment `json:"attachment,omitempty"`
	QuickReply []QR        `json:"quick_replies,omitempty"`
}

func (m *Message) AddQR(q ...QR) {
	m.QuickReply = append(m.QuickReply, q...)
}

type QR struct {
	Type    string `json:"content_type"`
	Title   string `json:"title"`
	Payload string `json:"payload"`
}

type Attachment struct {
	Type    string            `json:"type"`
	Payload AttachmentPayload `json:"payload"`
}

type AttachmentPayload interface{}

// Used as payload for type image/audio/video/file
type FilePayload struct {
	URL string `json:"url"`
}

type TemplateBase struct {
	Type string `json:"template_type"`
}

//Used as payload for type template(generic)
type GenericTemplate struct {
	TemplateBase
	Elements []Element `json:"elements"`
}

func (g GenericTemplate) Validate() error {
	if len(g.Elements) > 10 {
		return ErrBubblesLimitExceeded
	}
	for _, e := range g.Elements {
		if len(e.Title) > 45 {
			return ErrTitleTooLong
		}

		if len(e.Subtitle) > 80 {
			return ErrSubtitleTooLong
		}

		if len(e.Buttons) > 3 {
			return ErrButtonsLimitExceeded
		}
	}
	return nil
}

func (g *GenericTemplate) AddElement(e ...Element) {
	g.Elements = append(g.Elements, e...)
}

type ListTemplate struct {
	GenericTemplate
	TopElementStyle string `json:"top_element_style"` // compact or large(default)
}

type Element struct {
	Title    string   `json:"title"`
	URL      string   `json:"item_url,omitempty"`
	ImageURL string   `json:"image_url,omitempty"`
	Subtitle string   `json:"subtitle,omitempty"`
	Buttons  []Button `json:"buttons,omitempty"`
}

func (e *Element) AddButton(b ...Button) {
	e.Buttons = append(e.Buttons, b...)
}

type Button struct {
	Type    string `json:"type"`
	Title   string `json:"title,omitempty"`
	URL     string `json:"url,omitempty"`
	Payload string `json:"payload,omitempty"`
}

//Used as payload for type template(button)
type ButtonTemplate struct {
	TemplateBase
	Text    string   `json:"text,omitempty"`
	Buttons []Button `json:"buttons,omitempty"`
}

func (b *ButtonTemplate) AddButton(bt ...Button) {
	b.Buttons = append(b.Buttons, bt...)
}

//Used as payload for type template(receipt)
type ReceiptTemplate struct {
	TemplateBase
	RecipientName string            `json:"recipient_name"`
	ID            string            `json:"order_number"`
	Currency      string            `json:"currency"`
	PaymentMethod string            `json:"payment_method"`
	Timestamp     int64             `json:"timestamp,omitempty"`
	URL           string            `json:"order_url,omitempty"`
	Items         []OrderItem       `json:"elements"`
	Address       *OrderAddress     `json:"address,omitempty"`
	Summary       OrderSummary      `json:"summary"`
	Adjustments   []OrderAdjustment `json:"adjustments,omitempty"`
}

type OrderItem struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
	Price    int    `json:"price,omitempty"`
	Currency string `json:"currency,omitempty"`
	ImageURL string `json:"image_url,omiempty"`
}

type OrderAddress struct {
	Street1    string `json:"street_1"`
	Street2    string `json:"street_2,omitempty"`
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
	State      string `json:"state"`
	Country    string `json:"country"`
}

type OrderSummary struct {
	TotalCost    int `json:"total_cost,omitempty"`
	Subtotal     int `json:"subtotal,omitempty"`
	ShippingCost int `json:"shipping_cost,omitempty"`
	TotalTax     int `json:"total_tax,omitempty"`
}

type OrderAdjustment struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type APIResponse struct {
	RID   int64         `json:"recipient_id,string"`
	MID   string        `json:"message_id"`
	Error ErrorResponse `json:"error"`
}

type ErrorResponse struct {
	Message    string `json:"message"`
	Type       string `json:"type"`
	Code       int    `json:"code"`
	ErrorData  string `json:"error_data"`
	FBstraceID string `json:"fbstrace_id"`
}
