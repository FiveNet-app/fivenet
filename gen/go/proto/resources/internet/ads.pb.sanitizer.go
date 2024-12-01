// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: resources/internet/ads.proto

package internet

import (
	"github.com/fivenet-app/fivenet/pkg/html/htmlsanitizer"
)

func (m *Ad) Sanitize() error {

	m.Description = htmlsanitizer.StripTags(m.Description)

	m.Title = htmlsanitizer.StripTags(m.Title)

	return nil
}