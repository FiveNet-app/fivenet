// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: services/centrum/centrum.proto

package centrum

import (
	"github.com/fivenet-app/fivenet/pkg/html/htmlsanitizer"
)

func (m *TakeDispatchRequest) Sanitize() error {

	if m.Reason != nil {
		*m.Reason = htmlsanitizer.Sanitize(*m.Reason)
	}

	return nil
}

func (m *UpdateDispatchStatusRequest) Sanitize() error {

	if m.Code != nil {
		*m.Code = htmlsanitizer.Sanitize(*m.Code)
	}

	if m.Reason != nil {
		*m.Reason = htmlsanitizer.Sanitize(*m.Reason)
	}

	return nil
}

func (m *UpdateUnitStatusRequest) Sanitize() error {

	if m.Code != nil {
		*m.Code = htmlsanitizer.Sanitize(*m.Code)
	}

	if m.Reason != nil {
		*m.Reason = htmlsanitizer.Sanitize(*m.Reason)
	}

	return nil
}
