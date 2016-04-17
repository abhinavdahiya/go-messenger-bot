package mbotapi

import "github.com/satori/go.uuid"

// Creates a User from ID(int64)
func NewUserFromID(id int64) User {
	return User{
		ID: id,
	}
}

// Creates a User from phone no(string)
func NewUserFromPhone(p string) User {
	return User{
		PhoneNumber: p,
	}
}

// Creates a Message with given text to be sent by bot
func NewMessage(text string) Message {
	return Message{
		Text: text,
	}
}

// Creates a Message with image attachment to be
// sent by bot
// Currently supports url string only
// TODO: local file needs to be implmented
func NewImageMessage(url string) Message {
	return Message{
		Attachment: &Attachment{
			Type: "image",
			Payload: ImagePayload{
				URL: url,
			},
		},
	}
}

// Creates an empty generic template
func NewGenericTemplate() GenericTemplate {
	return GenericTemplate{
		TemplateBase: TemplateBase{
			Type: "generic",
		},
		Elements: []Element{},
	}
}

// Creates a new Element (info card) to be sent as
// part of the generic template
func NewElement(title string) Element {
	return Element{
		Title: title,
	}
}

// Creates a new Button of type `web_url`
func NewURLButton(title, url string) Button {
	return Button{
		Type:  "web_url",
		Title: title,
		URL:   url,
	}
}

// Creates a new Button of type `postback`
func NewPostbackButton(title, postback string) Button {
	return Button{
		Type:    "postback",
		Title:   title,
		Payload: postback,
	}
}

// Creates an empty Button Template
func NewButtonTemplate(text string) ButtonTemplate {
	return ButtonTemplate{
		TemplateBase: TemplateBase{
			Type: "button",
		},
		Text:    text,
		Buttons: []Button{},
	}
}

// Creates an empty Receipt Template
func NewReceiptTemplate(rname string) ReceiptTemplate {
	return ReceiptTemplate{
		TemplateBase: TemplateBase{
			Type: "receipt",
		},
		RecipientName: rname,
		ID:            uuid.NewV4().String(),
		Currency:      "USD",
		PaymentMethod: "",
		Items:         []OrderItem{},
		Summary:       OrderSummary{},
	}
}
