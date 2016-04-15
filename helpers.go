package mbotapi

import "github.com/satori/go.uuid"

func NewUserFromID(id int64) User {
	return User{
		ID: id,
	}
}

func NewUserFromPhone(p string) User {
	return User{
		PhoneNo: p,
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
		Type:     "generic",
		Elements: []Element{},
	}
}

func NewButtonTemplate(title string) ButtonTemplate {
	return ButtonTemplate{
		Type:    "button",
		Title:   title,
		Buttons: []Button{},
	}
}

func NewReceiptTemplate(rname string) ReceiptTemplate {
	return ReceiptTemplate{
		Type:          "receipt",
		RecipientName: rname,
		ID:            uuid.NewV4().String(),
		Currency:      "USD",
		PaymentMethod: "",
		Items:         []OrderItem{},
		Summary:       OrderSummary{},
	}
}
