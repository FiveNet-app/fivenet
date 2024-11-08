// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: resources/qualifications/exam.proto

package qualifications

import (
	"github.com/fivenet-app/fivenet/pkg/html/htmlsanitizer"
)

func (m *ExamQuestion) Sanitize() error {

	if m.Description != nil {
		*m.Description = htmlsanitizer.StripTags(*m.Description)
	}

	m.Title = htmlsanitizer.StripTags(m.Title)

	return nil
}

func (m *ExamResponseSingleChoice) Sanitize() error {

	m.Choice = htmlsanitizer.StripTags(m.Choice)

	return nil
}

func (m *ExamResponseText) Sanitize() error {

	m.Text = htmlsanitizer.StripTags(m.Text)

	return nil
}
