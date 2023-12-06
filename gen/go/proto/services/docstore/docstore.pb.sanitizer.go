// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: services/docstore/docstore.proto

package docstore

import (
	"github.com/galexrt/fivenet/pkg/htmlsanitizer"
)

func (m *CreateDocumentReqRequest) Sanitize() error {

	if m.Reason != nil {
		*m.Reason = htmlsanitizer.Sanitize(*m.Reason)
	}

	return nil
}

func (m *CreateDocumentRequest) Sanitize() error {

	m.Content = htmlsanitizer.Sanitize(m.Content)

	m.State = htmlsanitizer.Sanitize(m.State)

	m.Title = htmlsanitizer.StripTags(m.Title)

	return nil
}

func (m *UpdateDocumentReqRequest) Sanitize() error {

	if m.Reason != nil {
		*m.Reason = htmlsanitizer.Sanitize(*m.Reason)
	}

	return nil
}

func (m *UpdateDocumentRequest) Sanitize() error {

	m.Content = htmlsanitizer.Sanitize(m.Content)

	m.State = htmlsanitizer.Sanitize(m.State)

	m.Title = htmlsanitizer.StripTags(m.Title)

	return nil
}
