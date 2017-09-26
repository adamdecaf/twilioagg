package proxy

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/adamdecaf/twilioagg/phone"
)

const (
	smsMaxLength = 160
)

var (
	privateSMSNumber = os.Getenv("TWILIOAGG_PRIVATE_NUMBER")

	// ttl cache for voice calls
	// Some callers seem to be calling rapidly in a short burst,
	// which is probably due to the quick response/cancel from twilio.
	voiceTTLCache = sync.Map{}
	voiceTTLThreshold = 30 * time.Second
	voiceTTLTimestampFormat = time.UnixDate
	voiceTTLCacheCleanInterval = 10 * time.Second
)

func init() {
	if privateSMSNumber == "" {
		log.Fatalf("error - no TWILIOAGG_PRIVATE_NUMBER specified")
	}

	initCacheCleaning()
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
		if to == "" || (body == "" && len(sms.MediaUrls) == 0) {
			log.Printf("invalid message from [private] to %s", sms.To.Number)
			return
		}

		log.Printf("reply from [private] to %s", to)

		// Send the actual sms from the number we replied to
		err := sendSMS(sms.To.Number, to, body, sms.MediaUrls)
		if err != nil {
			log.Println(err)
		}
		return
	}

	// Otherwise, proxy incoming sms over to our private number
	from := sms.From.Number
	to := sms.To.Number

	var msg []rune
	var tpe string
	if len(sms.MediaUrls) > 0 {
		tpe = "MMS"
	} else {
		tpe = "SMS"
	}

	if strings.TrimSpace(sms.Body) == "" {
		msg = []rune(fmt.Sprintf("%s from %s\n", tpe, sms.From))
	} else {
		msg = []rune(fmt.Sprintf("%s from %s - \n%s", tpe, sms.From, sms.Body))
	}

	msgs := split(msg, smsMaxLength)

	log.Printf("incoming sms from %s to %s", from, to)

	for i := range msgs {
		log.Printf("sending part %d/%d", i+1, len(msgs))

		// Only add the MMS attachment onto the first message
		murls := make([]string, 0)
		if i == 0 {
			murls = sms.MediaUrls
		} else {
			murls = nil
		}

		// Send the actual sms from a number we own
		err := sendSMS(to, privateSMSNumber, string(msgs[i]), murls)
		if err != nil {
			log.Println(err)
		}
	}

	return
}

func getTTL() string {
	return time.Now().UTC().Add(voiceTTLThreshold).Format(voiceTTLTimestampFormat)
}

func ttlToOld(ttl string) bool {
	t, err := time.Parse(voiceTTLTimestampFormat, ttl)
	if err != nil {
		log.Printf("error parsing ttl from cache - %s\n", err)
		return false
	}
	return time.Now().After(t)
}

func initCacheCleaning() {
	clean := func() {
		for _ = range time.Tick(voiceTTLCacheCleanInterval) {
			voiceTTLCache.Range(func (k, v interface{}) bool {
				ttl, ok := v.(string)
				if ok && ttlToOld(ttl) {
					voiceTTLCache.Delete(k)
				}
				return true
			})
		}
	}
	go clean()
}

// HandleVoice currently sends an sms to the private number telling
// which number is calling and what public number was called
func HandleVoice(voice phone.Voice) {
	from := voice.From.Number

	// Check if the number is in the voice ttl cache
	when := getTTL()
	v, exists := voiceTTLCache.LoadOrStore(from, when)
	ttl, ok := v.(string)
	if !ok {
		log.Printf("error casting ttl as string")
		return
	}

	// check the ttl, if it's too old (or no ttl existed) then send the sms
	if !exists || ttlToOld(ttl) {
		// send sms
		to := privateSMSNumber
		details :=  fmt.Sprintf("Name: %s\nNumber: %s\nAddress: %s", voice.Name, voice.From.Number, voice.From.String())
		body := fmt.Sprintf("Incoming voice from %s, details:\n %s", voice.From.Number, details)
		err := sendSMS(voice.To.Number, to, body, nil)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Printf("not sending sms for voice req from %s\n", from)
	}
}

// Parse a body like:
// 123-456-7890 what's up world?
// into
// to: +11234567890
// body: what's up world?
func parseToNumberFromBody(b string) (to, body string) {
	parts := strings.SplitN(b, " ", 2)

	// return the first part in case of MMS messages
	// where there's no additional message to be sent
	if len(parts) == 1 {
		return parts[0], ""
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
