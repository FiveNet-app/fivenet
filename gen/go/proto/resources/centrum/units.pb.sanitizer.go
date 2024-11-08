// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: resources/centrum/units.proto

package centrum

import (
	"github.com/fivenet-app/fivenet/pkg/html/htmlsanitizer"
)

func (m *Unit) Sanitize() error {

	m.Color = htmlsanitizer.StripTags(m.Color)

	if m.Description != nil {
		*m.Description = htmlsanitizer.Sanitize(*m.Description)
	}

	m.Initials = htmlsanitizer.Sanitize(m.Initials)

	m.Name = htmlsanitizer.Sanitize(m.Name)

	return nil
}

func (m *UnitStatus) Sanitize() error {

	if m.Code != nil {
		*m.Code = htmlsanitizer.Sanitize(*m.Code)
	}

	if m.Postal != nil {
		*m.Postal = htmlsanitizer.Sanitize(*m.Postal)
	}

	if m.Reason != nil {
		*m.Reason = htmlsanitizer.Sanitize(*m.Reason)
	}

	return nil
}
