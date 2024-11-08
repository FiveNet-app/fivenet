// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: resources/documents/documents.proto

package documents

import (
	"github.com/fivenet-app/fivenet/pkg/html/htmlsanitizer"
)

func (m *Document) Sanitize() error {

	m.Content = htmlsanitizer.Sanitize(m.Content)

	if m.Data != nil {
		*m.Data = htmlsanitizer.Sanitize(*m.Data)
	}

	m.State = htmlsanitizer.Sanitize(m.State)

	m.Title = htmlsanitizer.Sanitize(m.Title)

	return nil
}

func (m *DocumentShort) Sanitize() error {

	m.Content = htmlsanitizer.Sanitize(m.Content)

	m.State = htmlsanitizer.Sanitize(m.State)

	m.Title = htmlsanitizer.Sanitize(m.Title)

	return nil
}
