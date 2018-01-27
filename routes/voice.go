package routes

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/adamdecaf/twilioagg/phone"
	"github.com/adamdecaf/twilioagg/proxy"
)

func IncomingVoice(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("error parsing form - %s", err)
		fmt.Fprintf(w, "error")
		return
	}

	voice := parseVoice(r.PostForm)
	log.Printf("incoming voice http webhook from %s\n", voice.From.Number)
	proxy.HandleVoice(voice)
}

func parseVoice(v url.Values) phone.Voice {
	voice := phone.Voice{}
	voice.Id = v.Get("CallSid")
	voice.Name = v.Get("CallerName")

	voice.From = phone.Subject{}
	voice.From.Number = v.Get("From")
	voice.From.City = v.Get("FromCity")
	voice.From.Country = v.Get("FromCountry")
	voice.From.State = v.Get("FromState")
	voice.From.Zip = v.Get("FromState")

	voice.To = phone.Subject{}
	voice.To.Number = v.Get("To")
	voice.To.City = v.Get("ToCity")
	voice.To.Country = v.Get("ToCountry")
	voice.To.State = v.Get("ToState")
	voice.To.Zip = v.Get("ToState")

	return voice
}
