// DO NOT EDIT THIS FILE. This file will be overwritten when re-running go-raml.
package types

import (
	"encoding/json"
	"gopkg.in/validator.v2"
)

type ServiceCreate struct {
	Data     json.RawMessage `json:"data,omitempty"`
	Name     string          `json:"name,omitempty"`
	Template string          `json:"template" validate:"nonzero"`
}

func (s ServiceCreate) Validate() error {

	return validator.Validate(s)
}
