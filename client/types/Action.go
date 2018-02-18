// DO NOT EDIT THIS FILE. This file will be overwritten when re-running go-raml.
package types

import (
	"gopkg.in/validator.v2"
)

type Action struct {
	Arguments []string `json:"arguments,omitempty"`
	Name      string   `json:"name" validate:"nonzero"`
}

func (s Action) Validate() error {

	return validator.Validate(s)
}
