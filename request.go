package gosoap

import (
	"fmt"
	"net/http"
)

// Request Soap Request
type Request struct {
	Method  string
	Params  SoapParams
	Headers http.Header
}

func NewRequest(m string, p SoapParams, h http.Header) *Request {
	return &Request{
		Method:  m,
		Params:  p,
		Headers: h,
	}
}

// RequestStruct soap request interface
type RequestStruct interface {
	SoapBuildRequest() *Request
}

// NewRequestByStruct create a new request using builder
func NewRequestByStruct(s RequestStruct) (*Request, error) {
	if s == nil {
		return nil, fmt.Errorf("'s' cannot be 'nil'")
	}

	return s.SoapBuildRequest(), nil
}
