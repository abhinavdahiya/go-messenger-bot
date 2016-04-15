package mbotapi

import "github.com/satori/go.uuid"

func NewUserFromID(id int64) User {
	return User{
		ID: id,
	}
}

func NewUserFromPhone(p string) User {
	return User{
		PhoneNumber: p,
	}
}

func NewMessage(text string) Message {
	return Message{
		Text: text,
	}
}

func NewImageMesssage(url string) Message {
	return Message{
		Attachment: Attachment{
			Type: "image",
			Payload: ImagePayload{
				URL: url,
			},
		},
	}
}

func NewGenericTemplate() GenericTemplate {
	return GenericTemplate{
		TemplateBase: TemplateBase{
			Type: "generic",
		},
		Elements: []Element{},
	}
}

func NewButtonTemplate(text string) ButtonTemplate {
	return ButtonTemplate{
		TemplateBase: TemplateBase{
			Type: "button",
		},
		Text:    text,
		Buttons: []Button{},
	}
}

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
