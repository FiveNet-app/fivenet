// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: resources/users/users.proto

package users

import (
	"github.com/fivenet-app/fivenet/pkg/htmlsanitizer"
)

func (m *UserActivity) Sanitize() error {

	m.Key = htmlsanitizer.Sanitize(m.Key)

	m.Reason = htmlsanitizer.Sanitize(m.Reason)

	return nil
}
