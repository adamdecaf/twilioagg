package proxy

import (
	"fmt"
	"log"
	"strings"
	"os"

	"github.com/adamdecaf/twilioagg/phone"
)

const (
	smsMaxLength = 160
)

var (
	privateSMSNumber = os.Getenv("TWILIOAGG_PRIVATE_NUMBER")
)

func init() {
	if privateSMSNumber == "" {
		log.Fatalf("error - no TWILIOAGG_PRIVATE_NUMBER specified")
	}
}

// HandleSMS works to proxy incoming and outgoing sms messages across multiple public
// numbers to a private number.
//
// Here's an example flow
//  [Incoming SMS from 123-456-7890]: "hello, world!"
//    -> Proxy to privateSMSNumber: "SMS from 123-456-7890 - hello, world!"
//  [Incoming SMS from privateSMSNumber]: "123-456-7890 what's up world?"
//    -> rewrite message and send "what's up world?" to 123-456-7890
func HandleSMS(sms phone.SMS) {
	// If we've gotten a reply from the private number, parse it and
	// send the response back out.
	if sms.From.Number == privateSMSNumber {
		to, body := parseToNumberFromBody(sms.Body)
		if to == "" || body == "" {
			log.Printf("invalid message from [private] to %s", sms.To.Number)
			return
		}

		log.Printf("reply from [private] to %s", to)

		// Send the actual sms from the number we replied to
		err := sendSMS(sms.To.Number, to, body)
		if err != nil {
			log.Println(err)
		}
		return
	}

	// Otherwise, proxy incoming sms over to our private number
	from := sms.From.Number
	to := sms.To.Number
	msg := []rune(fmt.Sprintf("SMS from %s - \n%s", sms.From, sms.Body))
	msgs := split(msg, smsMaxLength)
	log.Printf("incoming sms from %s to %s", from, to)
	for i := range msgs {
		log.Printf("sending part %d/%d", i+1, len(msgs))
		// Send the actual sms from a number we own
		err := sendSMS(to, privateSMSNumber, string(msgs[i]))
		if err != nil {
			log.Println(err)
		}
	}

	return
}

// Parse a body like:
// 123-456-7890 what's up world?
// into
// to: +11234567890
// body: what's up world?
func parseToNumberFromBody(b string) (to, body string) {
	parts := strings.SplitN(b, " ", 2)

	if len(parts) == 1 {
		return "", ""
	}

	to = parts[0]
	body = strings.Join(parts[1:], "")

	// TODO(adam): assumes +1 for country code
	to = strings.Replace(to, "-", "", -1)
	if !strings.HasPrefix(to, "+1") {
		to = "+1" + to
	}

	return
}

// split a []string into chunks based on a fixed limit
// from https://gist.github.com/xlab/6e204ef96b4433a697b3
func split(buf []rune, lim int) [][]rune {
	var chunk []rune
	chunks := make([][]rune, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:len(buf)])
	}
	return chunks
}
