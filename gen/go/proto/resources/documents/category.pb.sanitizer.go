// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: resources/documents/category.proto

package documents

import (
	"github.com/fivenet-app/fivenet/pkg/html/htmlsanitizer"
)

func (m *Category) Sanitize() error {

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

	return nil
}
