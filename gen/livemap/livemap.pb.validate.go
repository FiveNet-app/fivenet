// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: livemap/livemap.proto

package livemap

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

// Validate checks the field values on StreamRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StreamRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StreamRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StreamRequestMultiError, or
// nil if none found.
func (m *StreamRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StreamRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return StreamRequestMultiError(errors)
	}

	return nil
}

// StreamRequestMultiError is an error wrapping multiple validation errors
// returned by StreamRequest.ValidateAll() if the designated constraints
// aren't met.
type StreamRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StreamRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StreamRequestMultiError) AllErrors() []error { return m }

// StreamRequestValidationError is the validation error returned by
// StreamRequest.Validate if the designated constraints aren't met.
type StreamRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamRequestValidationError) ErrorName() string { return "StreamRequestValidationError" }

// Error satisfies the builtin error interface
func (e StreamRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamRequestValidationError{}

// Validate checks the field values on ServerStreamResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ServerStreamResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ServerStreamResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ServerStreamResponseMultiError, or nil if none found.
func (m *ServerStreamResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ServerStreamResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetUsers() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ServerStreamResponseValidationError{
						field:  fmt.Sprintf("Users[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ServerStreamResponseValidationError{
						field:  fmt.Sprintf("Users[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ServerStreamResponseValidationError{
					field:  fmt.Sprintf("Users[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetDispatches() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ServerStreamResponseValidationError{
						field:  fmt.Sprintf("Dispatches[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ServerStreamResponseValidationError{
						field:  fmt.Sprintf("Dispatches[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ServerStreamResponseValidationError{
					field:  fmt.Sprintf("Dispatches[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ServerStreamResponseMultiError(errors)
	}

	return nil
}

// ServerStreamResponseMultiError is an error wrapping multiple validation
// errors returned by ServerStreamResponse.ValidateAll() if the designated
// constraints aren't met.
type ServerStreamResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServerStreamResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServerStreamResponseMultiError) AllErrors() []error { return m }

// ServerStreamResponseValidationError is the validation error returned by
// ServerStreamResponse.Validate if the designated constraints aren't met.
type ServerStreamResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServerStreamResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServerStreamResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServerStreamResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServerStreamResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServerStreamResponseValidationError) ErrorName() string {
	return "ServerStreamResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ServerStreamResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServerStreamResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServerStreamResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServerStreamResponseValidationError{}

// Validate checks the field values on Marker with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Marker) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Marker with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in MarkerMultiError, or nil if none found.
func (m *Marker) ValidateAll() error {
	return m.validate(true)
}

func (m *Marker) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for X

	// no validation rules for Y

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Icon

	// no validation rules for Popup

	if len(errors) > 0 {
		return MarkerMultiError(errors)
	}

	return nil
}

// MarkerMultiError is an error wrapping multiple validation errors returned by
// Marker.ValidateAll() if the designated constraints aren't met.
type MarkerMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MarkerMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MarkerMultiError) AllErrors() []error { return m }

// MarkerValidationError is the validation error returned by Marker.Validate if
// the designated constraints aren't met.
type MarkerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MarkerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MarkerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MarkerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MarkerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MarkerValidationError) ErrorName() string { return "MarkerValidationError" }

// Error satisfies the builtin error interface
func (e MarkerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMarker.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MarkerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MarkerValidationError{}
