package phone

// AccountSid - [ACa3a93978e12e4db98f70fd57f800e7e6]
// ApiVersion - [2010-04-01]
// Body - [Test]
// From - [+15158050656]
// FromCity - []
// FromCountry - [US]
// FromState - [IA]
// FromZip - []
// MessageSid - [SMbf763e9221358fe15bc5dc11268f936e]
// NumMedia - [0]
// SmsMessageSid - [SMbf763e9221358fe15bc5dc11268f936e]
// SmsSid - [SMbf763e9221358fe15bc5dc11268f936e]
// SmsStatus - [received]
// To - [+15412334632]
// ToCity - [PRINEVILLE]
// ToCountry - [US]
// ToState - [OR]
// ToZip - [97754]

type SMS struct {
	Id        string // MessageSid, SmsSid, SmsMessageSid
	Body      string
	From      Subject
	To        Subject
	MediaUrls []string
}
