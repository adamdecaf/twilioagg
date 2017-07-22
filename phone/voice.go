package phone

// AccountSid - [ACa3a93978e12e4db98f70fd57f800e7e6]
// CallSid - [CAb7028ff79e88cf1cc33c08b5ed85dfab]
// CallStatus - [ringing]
// Called - [+15412334632]
// CalledCity - [PRINEVILLE]
// CalledCountry - [US]
// CalledState - [OR]
// CalledZip - [97754]
// Caller - [+15158050656]
// CallerCity - []
// CallerCountry - [US]
// CallerName - [URBANDALE  IA]
// CallerState - [IA]
// CallerZip - []
// Direction - [inbound]
// From - [+15158050656]
// FromCity - []
// FromCountry - [US]
// FromState - [IA]
// FromZip - []
// To - [+15412334632]
// ToCity - [PRINEVILLE]
// ToCountry - [US]
// ToState - [OR]
// ToZip - [97754]

type Voice struct {
	Id string
	Name string
	From Subject
	To Subject
}
