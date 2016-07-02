package mbotapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

const (
	SettingsEndpoint = "https://graph.facebook.com/v2.6/me/thread_settings?access_token=%s"
)

type ThreadSetting struct {
	Type     string     `json:"setting_type"`
	State    string     `json:"thread_statei,omitempty"`
	GStarted []GStarted `json:"call_to_action,omitempty"`
	Menu     []Button   `json:"call_to_action,omitempty"`
	Greeting Greeting   `json:"greeting,omitempty"`
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

	var rsp struct {
		Result string `json:"result"`
	}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&rsp)
	if err != nil {
		return err
	}

	log.Printf("[SETTINGS] %#v", rsp)
	if resp.StatusCode != 200 {
		return errors.New(http.StatusText(resp.StatusCode))
	}
	return nil
}

func (bot *BotAPI) SetGreeting(text string) error {
	g := ThreadSetting{
		Type: "greeting",
		Greeting: Greeting{
			Text: text,
		},
	}

	payl, _ := json.Marshal(g)
	return bot.SetSettings(bytes.NewBuffer(payl))
}

func (bot *BotAPI) SetGStarted(text string) error {
	g := ThreadSetting{
		Type:  "call_to_actions",
		State: "new_thread",
		GStarted: []GStarted{
			{text},
		},
	}

	payl, _ := json.Marshal(g)
	return bot.SetSettings(bytes.NewBuffer(payl))
}

func (bot *BotAPI) SetMenu(bts []Button) error {
	g := ThreadSetting{
		Type:  "call_to_actions",
		State: "existing_thread",
		Menu:  bts,
	}

	payl, _ := json.Marshal(g)
	return bot.SetSettings(bytes.NewBuffer(payl))
}
