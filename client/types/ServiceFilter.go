// DO NOT EDIT THIS FILE. This file will be overwritten when re-running go-raml.
package types

import (
	"gopkg.in/validator.v2"
)

// type used in the query parameters of the ListService method
type ServiceFilter struct {
}

func (s ServiceFilter) Validate() error {

	return validator.Validate(s)
}
