## twilioagg

A twilio aggregator for multiple phone numbers that routes to a single (private) number.

## Usage

The following environment variables are required for the `twilioagg` binary.

- `TWILIOAGG_PRIVATE_NUMBER`
- `TWILIO_ACCOUNT_SID`
- `TWILIO_AUTH_TOKEN`

After the binary is running you'd need to configure the number's webhooks in twilio to send requests towards your instance. The default would be:

- SMS: `your-server.com:8080/sms`
- Voice: `your-server.com:8080/voice`

### Sending from your private number

Your private number will get SMS messages like this:

```
SMS from <E.164 number> from <city>, <state> <zip> <country> - <message>
```

You can send replies back out by sending a text message to a public number (from your private number) in the format:

```
<number> <message>
```

An example (note, you can currently leave off the country code, it defaults to `+1`):

```
5555555555 hi mom!
```

The recipient would get the following message:

```
hi mom!
```

### Features

- Accept sms/calls/mms for multiple incoming numbers
  - Forward along to an email or diff phone number
  - Respond from private number back out with correct public number

### TODO / Future

- blocking incoming numbers
- alert the private number on incoming calls
- record incoming calls to fs, db, etc
- buy phone numbers from twilio
  - find easy to type numbers (based on some distance metric off the T9 keyboard...?)
  - cycle public numbers
