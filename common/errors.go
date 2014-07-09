package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func DecodeError(err error, body []byte) error {
	return errors.New(fmt.Sprintf("Error while decoding json => %s, recieved msg => %s", err.Error(), string(body)))
}

func CheckResponse(resp http.Response, err error) error {
	if err != nil {
		return errors.New(fmt.Sprintf("Error sending req => %s", err.Error()))
	}
	if resp.StatusCode != 201 {
		return NewTwilioError(resp)
	}
	return nil
}

type Error struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	MoreInfo string `json:"more_info"`
	Status   string `json:"status"`
}

func (e Error) Error() string {
	return fmt.Sprintf("Twilio Error %d => %s, more info @ %s", e.Code, e.Message, e.MoreInfo)
}

func NewTwilioError(resp http.Response) error {
	var twilioErr Error
	var buf bytes.Buffer
	_, err := buf.ReadFrom(resp.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("Twilio error encountered, failure while reading body => %s", err.Error()))
	}

	err = json.Unmarshal(buf.Bytes(), &twilioErr)
	if err != nil {
		return errors.New(fmt.Sprintf("Twilio error encountered, failure while parsing => %s", err.Error()))
	}
	return twilioErr
}
