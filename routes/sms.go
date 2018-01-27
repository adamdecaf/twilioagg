package routes

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

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
	log.Printf("incoming sms http webhook from %s\n", sms.From.Number)
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

	// Pull out each "MediaUrl%d" param
	n, err := strconv.Atoi(v.Get("NumMedia"))
	if err == nil && n > 0 {
		urls := make([]string, 0)
		for i := 0; i < n; i++ {
			u := v.Get(fmt.Sprintf("MediaUrl%d", i))
			urls = append(urls, u)
		}
		sms.MediaUrls = urls
	}

	return sms
}
