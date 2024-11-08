// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: resources/livemap/livemap.proto

package livemap

import (
	"github.com/fivenet-app/fivenet/pkg/html/htmlsanitizer"
)

func (m *IconMarker) Sanitize() error {

	m.Icon = htmlsanitizer.StripTags(m.Icon)

	return nil
}

func (m *MarkerInfo) Sanitize() error {

	if m.Color != nil {
		*m.Color = htmlsanitizer.StripTags(*m.Color)
	}

	if m.Description != nil {
		*m.Description = htmlsanitizer.Sanitize(*m.Description)
	}

	if m.Icon != nil {
		*m.Icon = htmlsanitizer.StripTags(*m.Icon)
	}

	m.Name = htmlsanitizer.Sanitize(m.Name)

	if m.Postal != nil {
		*m.Postal = htmlsanitizer.Sanitize(*m.Postal)
	}

	return nil
}
