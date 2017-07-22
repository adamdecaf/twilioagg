## twilioagg

A twilio aggregator for multiple phone numbers that routes to a single (private) number.

## Usage

The following environment variables are required for the `twilioagg` binary.

- `TWILIOAGG_PRIVATE_NUMBER`
- `TWILIO_ACCOUNT_SID`
- `TWILIO_AUTH_TOKEN`

### Features

- Accept sms/calls/mms for multiple incoming numbers
  - Forward along to an email or diff phone number
  - Respond from private number back out with correct public number

### TODO / Future

- record incoming calls to fs, db, etc
- buy phone numbers from twilio
  - find easy to type numbers (based on some distance metric off the T9 keyboard...?)
  - cycle public numbers
