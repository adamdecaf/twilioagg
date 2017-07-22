package routes

import (
	"fmt"
	"net/http"
	"net/url"
	"log"

	"github.com/adamdecaf/twilioagg/phone"
	"github.com/adamdecaf/twilioagg/proxy"
)

func IncomingSMS(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("error parsing form - %s", err)
		fmt.Fprintf(w, "error")
		return
	}

	sms := parseSMS(r.PostForm)
	proxy.HandleSMS(sms)
}

func parseSMS(v url.Values) phone.SMS {
	sms := phone.SMS{}
	sms.Id = v.Get("MessageSid")
	sms.Body = v.Get("Body")

	sms.From = phone.Subject{}
	sms.From.Number = v.Get("From")
	sms.From.City = v.Get("FromCity")
	sms.From.Country = v.Get("FromCountry")
	sms.From.State = v.Get("FromState")
	sms.From.Zip = v.Get("FromZip")

	sms.To = phone.Subject{}
	sms.To.Number = v.Get("To")
	sms.To.City = v.Get("ToCity")
	sms.To.Country = v.Get("ToCountry")
	sms.To.State = v.Get("ToState")
	sms.To.Zip = v.Get("ToZip")

	return sms
}
