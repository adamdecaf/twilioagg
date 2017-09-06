package proxy

import (
	"log"
	"os"
	gotwilio "github.com/sfreiberg/gotwilio"
)

var (
	twilio *gotwilio.Twilio
)

func init() {
	sid := os.Getenv("TWILIO_ACCOUNT_SID")
	token := os.Getenv("TWILIO_AUTH_TOKEN")

	if sid == "" || token == "" {
		log.Fatalf("error - no sid or token specified, check TWILIO_* env vars")
	}
	twilio = gotwilio.NewTwilioClient(sid, token)
}

func sendSMS(from, to, body string, mediaUrls []string) error {
	var resp *gotwilio.SmsResponse
	var exc *gotwilio.Exception
	var err error

	if len(mediaUrls) > 0 {
		// TODO(adam): Why only 1 MediaUrl??
		resp, exc, err = twilio.SendMMS(from, to, body, mediaUrls[0], "", "")
		log.Printf("send mms from %s to %s - %s", from, to, resp.Status)
	} else {
		resp, exc, err = twilio.SendSMS(from, to, body, "", "")
		log.Printf("send sms from %s to %s - %s", from, to, resp.Status)
	}

	if exc != nil {
		log.Println(exc)
	}

	return err
}
