package mbotapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type BotAPI struct {
	Token       string
	VerifyToken string
	Debug       bool
	Client      *http.Client
}

func NewBotAPI(token string, vtoken string) *BotAPI {
	return &BotAPI{
		Token:       token,
		VerifyToken: vtoken,
		Debug:       false,
		Client:      &http.Client{},
	}
}

func (bot *BotAPI) MakeRequest(b *bytes.Buffer) (APIResponse, error) {
	uri := fmt.Sprintf(APIEndpoint, bot.Token)

	req, _ := http.NewRequest("POST", uri, b)
	req.Header.Set("Content-Type", "application/json")
	resp, err := bot.Client.Do(req)
	if err != nil {
		return APIResponse{}, err
	}
	defer resp.Body.Close()

	var rsp APIResponse
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&rsp)
	if err != nil {
		return APIResponse{}, nil
	}

	if resp.StatusCode != 200 {
		return rsp, errors.New(http.StatusText(resp.StatusCode))
	}
	return rsp, nil
}

func (bot *BotAPI) Send(u User, c interface{}, notif string) (APIResponse, error) {
	var r Request

	n := RegularNotif
	if notif != "" {
		n = notif
	}

	switch c.(type) {
	case Request:
		return APIResponse{}, errors.New("Use MakeRequest to send Request!!")
	case Message:
		r = Request{
			Recipient: u,
			Message:   c.(Message),
			NotifType: n,
		}

	case GenericTemplate:
		r = Request{
			Recipient: u,
			NotifType: n,
			Message: Message{
				Attachment: &Attachment{
					Type:    "template",
					Payload: c.(GenericTemplate),
				},
			},
		}

	case ButtonTemplate:
		r = Request{
			Recipient: u,
			NotifType: n,
			Message: Message{
				Attachment: &Attachment{
					Type:    "template",
					Payload: c.(ButtonTemplate),
				},
			},
		}

	case ReceiptTemplate:
		r = Request{
			Recipient: u,
			NotifType: n,
			Message: Message{
				Attachment: &Attachment{
					Type:    "template",
					Payload: c.(ReceiptTemplate),
				},
			},
		}

	default:
		return APIResponse{}, errors.New("Type is not supported")
	}

	if r == (Request{}) {
		return APIResponse{}, errors.New("Unknown Error")
	}
	payl, _ := json.Marshal(r)
	if bot.Debug {
		log.Printf("[INFO] Payload: %s", string(payl))
	}
	return bot.MakeRequest(bytes.NewBuffer(payl))
}

func (bot *BotAPI) SendFile(u User, path string) (APIResponse, error) {
	file, err := os.Open(path)
	if err != nil {
		return APIResponse{}, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("filedata", filepath.Base(path))
	if err != nil {
		return APIResponse{}, err
	}
	_, err = io.Copy(part, file)

	usr, _ := json.Marshal(u)
	_ = writer.WriteField("recipient", string(usr))
	img := NewImageMessage("")
	im, _ := json.Marshal(img)
	_ = writer.WriteField("message", string(im))

	err = writer.Close()
	if err != nil {
		return APIResponse{}, err
	}

	return bot.MakeRequest(body)
}

func (bot *BotAPI) SetWebhook(pattern string) (<-chan Callback, *http.ServeMux) {
	callbackChan := make(chan Callback, 100)

	mux := http.NewServeMux()
	mux.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET":
			if req.FormValue("hub.verify_token") == bot.VerifyToken {
				w.Write([]byte(req.FormValue("hub.challenge")))
				return
			}
			w.WriteHeader(http.StatusUnauthorized)
			return

		case "POST":
			defer req.Body.Close()

			var rsp Response
			decoder := json.NewDecoder(req.Body)
			decoder.Decode(&rsp)

			if rsp.Object == "page" {
				for _, e := range rsp.Entries {
					for _, c := range e.Messaging {
						callbackChan <- c
					}
				}
			}
			w.WriteHeader(http.StatusOK)
			return

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	})

	return callbackChan, mux

}
