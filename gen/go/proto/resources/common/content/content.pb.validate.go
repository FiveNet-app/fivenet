// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: resources/common/content/content.proto

package content

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Content with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Content) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Content with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ContentMultiError, or nil if none found.
func (m *Content) ValidateAll() error {
	return m.validate(true)
}

func (m *Content) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.Version != nil {

		if utf8.RuneCountInString(m.GetVersion()) > 24 {
			err := ContentValidationError{
				field:  "Version",
				reason: "value length must be at most 24 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.Content != nil {

		if all {
			switch v := interface{}(m.GetContent()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ContentValidationError{
						field:  "Content",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ContentValidationError{
						field:  "Content",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetContent()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ContentValidationError{
					field:  "Content",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.RawContent != nil {

		if len(m.GetRawContent()) > 2000000 {
			err := ContentValidationError{
				field:  "RawContent",
				reason: "value length must be at most 2000000 bytes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return ContentMultiError(errors)
	}

	return nil
}

// ContentMultiError is an error wrapping multiple validation errors returned
// by Content.ValidateAll() if the designated constraints aren't met.
type ContentMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ContentMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ContentMultiError) AllErrors() []error { return m }

// ContentValidationError is the validation error returned by Content.Validate
// if the designated constraints aren't met.
type ContentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ContentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ContentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ContentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ContentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ContentValidationError) ErrorName() string { return "ContentValidationError" }

// Error satisfies the builtin error interface
func (e ContentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sContent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ContentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ContentValidationError{}

// Validate checks the field values on JSONNode with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *JSONNode) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on JSONNode with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in JSONNodeMultiError, or nil
// if none found.
func (m *JSONNode) ValidateAll() error {
	return m.validate(true)
}

func (m *JSONNode) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	// no validation rules for Id

	// no validation rules for Tag

	// no validation rules for Attributes

	// no validation rules for Class

	// no validation rules for Text

	for idx, item := range m.GetChildren() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, JSONNodeValidationError{
						field:  fmt.Sprintf("Children[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, JSONNodeValidationError{
						field:  fmt.Sprintf("Children[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return JSONNodeValidationError{
					field:  fmt.Sprintf("Children[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return JSONNodeMultiError(errors)
	}

	return nil
}

// JSONNodeMultiError is an error wrapping multiple validation errors returned
// by JSONNode.ValidateAll() if the designated constraints aren't met.
type JSONNodeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m JSONNodeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m JSONNodeMultiError) AllErrors() []error { return m }

// JSONNodeValidationError is the validation error returned by
// JSONNode.Validate if the designated constraints aren't met.
type JSONNodeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JSONNodeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JSONNodeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JSONNodeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JSONNodeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JSONNodeValidationError) ErrorName() string { return "JSONNodeValidationError" }

// Error satisfies the builtin error interface
func (e JSONNodeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJSONNode.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JSONNodeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JSONNodeValidationError{}
