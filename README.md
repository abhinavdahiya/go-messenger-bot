# Golang bindings for the Messenger Bot API

The scope of this project is just to provide a wrapper around the API
without any additional features. 

## Example

This is a very simple bot that just displays any gotten updates,
then replies it to it.

```go
package main

import (
	"log"
    "net/http"
	"github.com/abhinavdahiya/go-messenger-bot"
)

func main() {
	bot := mbotapi.NewBotAPI("ACCESS_TOKEN", "VERIFY_TOKEN")

	callbacks, mux := bot.SetWebhook("/webhook")
	go http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", mux)

    for callback := range callbacks {
        log.Printf("[%#v] %s", callback.Sender, callback.Message.Text)

        msg := mbotapi.NewMessage(callback.Message.Text)
        bot.Send(callback.Sender, msg, mbotapi.RegularNotif)
    }
}
```

Facebook messenger webhook needs a certificate certified by known CA,
Now that [Let's Encrypt](https://letsencrypt.org) has entered public beta,
you may wish to generate your free TLS certificate there.

## Inspiration

Messenger takes design cues from:

- [`go-telegram-bot-api/telegram-bot-api`](https://github.com/go-telegram-bot-api/telegram-bot-api)
