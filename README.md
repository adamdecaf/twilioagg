## twilioagg

A twilio aggregator for multiple phone numbers that routes to a single (private) number.

### TODO / Future

- accept sms/calls/mms for multiple incoming numbers
  - options to forward along to an email or diff phone number
    - ability to respond from private number back out with correct public number
  - record incoming calls to fs, db, etc
- buy phone numbers from twilio
  - find easy to type numbers (based on some distance metric off the T9 keyboard...?)
  - cycle public numbers
