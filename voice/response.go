package voice

import (
	"github.com/natebrennand/twiliogo/common"
	"regexp"
)

type Call struct {
	common.ResponseCore
	Price          common.JSONPrice `json:"price"`
	ParentCallSid  string
	PhoneNumberSid string
	StartTime      common.JSONTime `json:"start_time"`
	EndTime        common.JSONTime `json:"end_time"`
	Duration       string          `json:"duration"`
	AnsweredBy     string          `json:"answered_by"`
	ForwardedFrom  string          `json:"fowarded_from"`
	CallerName     string          `json:"caller_name"`
}

type CallList struct {
	common.ListResponseCore
	Calls *[]Call `json:"calls"`
}

func validateCallSid(sid string) bool {
	match, _ := regexp.MatchString(`^CA[0-9a-z]{32}$`, string(sid))
	return match
}

func validateRecSid(sid string) bool {
	match, _ := regexp.MatchString(`^RE[0-9a-z]{32}$`, string(sid))
	return match
}