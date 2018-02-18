// DO NOT EDIT THIS FILE. This file will be overwritten when re-running go-raml.
package client

import (
	"net/http"
)

const (
	defaultBaseURI = ""
)

type ZeroRobot struct {
	client     http.Client
	AuthHeader string // Authorization header, will be sent on each request if not empty
	BaseURI    string
	common     service // Reuse a single struct instead of allocating one for each service on the heap.

	Blueprints *BlueprintsService
	Services   *ServicesService
	Templates  *TemplatesService
}

type service struct {
	client *ZeroRobot
}

func NewZeroRobot() *ZeroRobot {
	c := &ZeroRobot{
		BaseURI: defaultBaseURI,
		client:  http.Client{},
	}
	c.common.client = c

	c.Blueprints = (*BlueprintsService)(&c.common)
	c.Services = (*ServicesService)(&c.common)
	c.Templates = (*TemplatesService)(&c.common)

	return c
}