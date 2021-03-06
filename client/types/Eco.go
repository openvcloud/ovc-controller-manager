// DO NOT EDIT THIS FILE. This file will be overwritten when re-running go-raml.
package types

import (
	"encoding/json"
	"gopkg.in/validator.v2"
)

// Error conditions, is a custom error format that allow to investigate
// an exceptions. It is used to serialize the error happening during task execution
type Eco struct {
	_limit             int             `json:"_limit" validate:"nonzero"`
	_traceback         string          `json:"_traceback" validate:"nonzero"`
	Appname            string          `json:"appname" validate:"nonzero"`
	Category           string          `json:"category" validate:"nonzero"`
	Closetime          int             `json:"closetime" validate:"nonzero"`
	Code               string          `json:"code" validate:"nonzero"`
	Data               json.RawMessage `json:"data,omitempty"`
	Epoch              int             `json:"epoch" validate:"nonzero"`
	Errormessage       string          `json:"errormessage" validate:"nonzero"`
	ErrormessagePub    string          `json:"errormessagePub" validate:"nonzero"`
	Exceptionclassname string          `json:"exceptionclassname" validate:"nonzero"`
	Funcfilename       string          `json:"funcfilename" validate:"nonzero"`
	Funclinenr         int             `json:"funclinenr" validate:"nonzero"`
	Funcname           string          `json:"funcname" validate:"nonzero"`
	Guid               string          `json:"guid" validate:"nonzero"`
	Jid                int             `json:"jid" validate:"nonzero"`
	Lasttime           int             `json:"lasttime" validate:"nonzero"`
	Level              int             `json:"level" validate:"nonzero"`
	Masterjid          int             `json:"masterjid" validate:"nonzero"`
	Occurrences        int             `json:"occurrences" validate:"nonzero"`
	Pid                int             `json:"pid" validate:"nonzero"`
	State              string          `json:"state" validate:"nonzero"`
	Tags               string          `json:"tags" validate:"nonzero"`
	Type               string          `json:"type" validate:"nonzero"`
	Uniquekey          string          `json:"uniquekey" validate:"nonzero"`
}

func (s Eco) Validate() error {

	return validator.Validate(s)
}
