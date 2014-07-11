package recording

import (
	"github.com/natebrennand/twiliogo/common"
	"regexp"
)

type Recording struct {
	Sid         string          `json:"sid"`
	DateCreated common.JSONTime `json:"date_created"`
	DateUpdated common.JSONTime `json:"date_updated"`
	AccountSid  string          `json:"account_sid"`
	CallSid     string          `json:"call_sid"`
	Duration    string          `json:"duration"`
	APIVersion  string          `json:"api_version"`
	URI         string          `json:"uri"`
}

type RecordingList struct {
	common.ListResponseCore
	Recordings *[]Recording `json:"recordings"`
}

func validateRecSid(sid string) bool {
	match, _ := regexp.MatchString(`^RE[0-9a-z]{32}$`, string(sid))
	return match
}