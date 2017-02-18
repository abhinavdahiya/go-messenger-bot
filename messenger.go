package mbotapi

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// This defines a bot
// Set Debug to true for debugging
type BotAPI struct {
	Token       string
	VerifyToken string
	AppSecret   string
	Debug       bool
	Client      *http.Client
}

// This helps create a BotAPI instance with token, verify_token
// By default Debug is set to false
func NewBotAPI(token string, vtoken string, secret string) *BotAPI {
	return &BotAPI{
		Token:       token,
		AppSecret:   secret,
		VerifyToken: vtoken,
		Debug:       false,
		Client:      &http.Client{},
	}
}

// This helps send request (send messages to users)
// It takes Request struct encoded into a buffer of json bytes
// The APIResponse contains the error from FB if any
// Should NOT be directly used, Use Send / SendFile
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

// This function helps send messages to users
// It takes Message / GenericTemplate / ButtonTemplate / ReceiptTemplate and
// sends it to the user
func (bot *BotAPI) Send(u User, c interface{}, notif string) (APIResponse, error) {
	var r Request

	n := RegularNotif
	if notif != "" {
		n = notif
	}

	switch c.(type) {
	case Request:
		return APIResponse{}, errors.New("Use MakeRequest to send Request!!")

	case Action:
		r = Request{
			Recipient: u,
			Action:    c.(Action),
			NotifType: n,
		}

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

	case ListTemplate:
		r = Request{
			Recipient: u,
			NotifType: n,
			Message: Message{
				Attachment: &Attachment{
					Type:    "template",
					Payload: c.(ListTemplate),
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

	//if r == (Request{}) {
	//	return APIResponse{}, errors.New("Unknown Error")
	//}
	payl, _ := json.Marshal(r)
	if bot.Debug {
		log.Printf("[INFO] Payload: %s", string(payl))
	}
	return bot.MakeRequest(bytes.NewBuffer(payl))
}

// This helps to send local images (currently) to users
// TODO: not tested yet!
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
	img := NewImageFromURL("")
	im, _ := json.Marshal(img)
	_ = writer.WriteField("message", string(im))

	err = writer.Close()
	if err != nil {
		return APIResponse{}, err
	}

	return bot.MakeRequest(body)
}

//This function verifies the message
func verifySignature(appSecret string, bytes []byte, expectedSignature string) bool {
	mac := hmac.New(sha1.New, []byte(appSecret))
	mac.Write(bytes)
	if fmt.Sprintf("%x", mac.Sum(nil)) != expectedSignature {
		return false
	}
	return true
}

// This function registers the handlers for
// - webhook verification
// - all callbacks made on the webhhoks
// It loops over all entries in the callback and
// pushes to the Callback channel
// This also return a *http.ServeMux which can be used to listenAndServe
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

			body, _ := ioutil.ReadAll(req.Body)
			req.Body = ioutil.NopCloser(bytes.NewReader(body))
			if bot.Debug {
				log.Printf("[INFO]%s", body)
			}

			if req.Header.Get("X-Hub-Signature") == "" || !verifySignature(bot.AppSecret, body, req.Header.Get("X-Hub-Signature")[5:]) {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
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
