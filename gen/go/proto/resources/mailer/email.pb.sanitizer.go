// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: resources/mailer/email.proto

package mailer

import (
	"github.com/fivenet-app/fivenet/pkg/html/htmlsanitizer"
)

func (m *Email) Sanitize() error {

	m.Email = htmlsanitizer.StripTags(m.Email)

	if m.Label != nil {
		*m.Label = htmlsanitizer.StripTags(*m.Label)
	}

	return nil
}

func (m *EmailShort) Sanitize() error {

	m.Email = htmlsanitizer.StripTags(m.Email)

	if m.Label != nil {
		*m.Label = htmlsanitizer.StripTags(*m.Label)
	}

	return nil
}
