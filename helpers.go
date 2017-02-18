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

//Create a sender action
// takes const TypingON/TypingOFF/MarkSeen
func NewAction(ac Action) Action {
	return ac
}

// Creates a Message with given text to be sent by bot
func NewMessage(text string) Message {
	return Message{
		Text: text,
	}
}

//Creates a quick reply
//Takes two parameters:
// - title(string)
// - postback_payload(string)
func NewQuickReply(title string, pst string) QR {
	return QR{
		Title:   title,
		Payload: pst,
		Type:    "text",
	}
}

// Creates a Message with image attachment to be
// sent by bot
func NewImageFromURL(url string) Message {
	return Message{
		Attachment: &Attachment{
			Type: "image",
			Payload: FilePayload{
				URL: url,
			},
		},
	}
}

// Creates a Message with audio attachment to be
// sent by bot
func NewAudioFromURL(url string) Message {
	return Message{
		Attachment: &Attachment{
			Type: "audio",
			Payload: FilePayload{
				URL: url,
			},
		},
	}
}

// Creates a Message with video attachment to be
// sent by bot
func NewVideoFromURL(url string) Message {
	return Message{
		Attachment: &Attachment{
			Type: "video",
			Payload: FilePayload{
				URL: url,
			},
		},
	}
}

// Creates a Message with file attachment to be
// sent by bot
func NewFileFromURL(url string) Message {
	return Message{
		Attachment: &Attachment{
			Type: "file",
			Payload: FilePayload{
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

// Creates an empty list template
func NewListTemplate() ListTemplate {
	return ListTemplate{
		GenericTemplate: GenericTemplate{
			TemplateBase: TemplateBase{
				Type: "list",
			},
			Elements: []Element{},
		},
		TopElementStyle: "large",
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

// Creates a new Button of type `phone_number`
func NewPhoneButton(title, pn string) Button {
	return Button{
		Type:    "phone_number",
		Title:   title,
		Payload: pn,
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
