package voice

const (
	postUrl              = "https://api.twilio.com/2010-04-01/Accounts/%s/Calls.json"
	updateUrl            = "https://api.twilio.com/2010-04-01/Accounts/%s/Calls/%s.json"
	getUrl               = "https://api.twilio.com/2010-04-01/Accounts/%s/Calls/%s.json"
	listUrl              = "https://api.twilio.com/2010-04-01/Accounts/%s/Calls.json"
	recordingUrl         = "https://api.twilio.com/2010-04-01/Accounts/%s/Recordings/%s.json"
	recordingListUrl     = "https://api.twilio.com/2010-04-01/Accounts/%s/Recordings.json"
	transcriptionUrl     = "https://api.twilio.com/2010-04-01/Accounts/%s/Transcriptions/%s.json"
	transcriptionListUrl = "https://api.twilio.com/2010-04-01/Accounts/%s/Transcriptions.json"
)

var errorCode = map[int]string{
	30001: "Queue Overflow",
	30002: "Account Suspended",
	30003: "Unreachable destination handset",
	30004: "Message blocked",
	30005: "Unknown destination handset",
	30006: "Landline or unreachable carrier",
	30007: "Carrier Violation",
	30008: "Unknown error",
}
