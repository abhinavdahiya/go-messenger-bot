package mbotapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	SettingsEndpoint = "https://graph.facebook.com/v2.6/me/thread_settings?access_token=%s"
)

type ThreadSetting struct {
	Type     string      `json:"setting_type"`
	State    string      `json:"thread_state,omitempty"`
	Action   interface{} `json:"call_to_actions,omitempty"`
	Greeting *Greeting   `json:"greeting,omitempty"`
}

type Greeting struct {
	Text string `json:"text"`
}

type GStarted struct {
	Payload string `json:"payload"`
}

func (bot *BotAPI) SetSettings(b *bytes.Buffer) error {
	uri := fmt.Sprintf(SettingsEndpoint, bot.Token)

	req, _ := http.NewRequest("POST", uri, b)
	req.Header.Set("Content-Type", "application/json")
	resp, err := bot.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	rsp, _ := ioutil.ReadAll(resp.Body)
	log.Printf("[SETTINGS] %s", rsp)
	if resp.StatusCode != 200 {
		return errors.New(http.StatusText(resp.StatusCode))
	}
	return nil
}

func (bot *BotAPI) SetGreeting(text string) error {
	g := ThreadSetting{
		Type: "greeting",
		Greeting: &Greeting{
			Text: text,
		},
	}

	log.Printf("[SETTINGS] %#v", g)
	payl, err := json.Marshal(g)
	if err != nil {
		return err
	}
	log.Printf("[SETTINGS] %s", payl)
	return bot.SetSettings(bytes.NewBuffer(payl))
}

func (bot *BotAPI) SetGStarted(text string) error {
	g := ThreadSetting{
		Type:  "call_to_actions",
		State: "new_thread",
		Action: []GStarted{
			{text},
		},
	}

	log.Printf("[SETTINGS] %#v", g)
	payl, err := json.Marshal(g)
	if err != nil {
		return err
	}
	log.Printf("[SETTINGS] %s", payl)
	return bot.SetSettings(bytes.NewBuffer(payl))
}

func (bot *BotAPI) SetMenu(bts []Button) error {
	g := ThreadSetting{
		Type:   "call_to_actions",
		State:  "existing_thread",
		Action: bts,
	}

	log.Printf("[SETTINGS] %#v", g)
	payl, err := json.Marshal(g)
	if err != nil {
		return err
	}
	log.Printf("[SETTINGS] %s", payl)
	return bot.SetSettings(bytes.NewBuffer(payl))
}
