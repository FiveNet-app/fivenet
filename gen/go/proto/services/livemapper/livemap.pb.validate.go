// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: services/livemapper/livemap.proto

package livemapper

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

// Validate checks the field values on StreamResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StreamResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StreamResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StreamResponseMultiError,
// or nil if none found.
func (m *StreamResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *StreamResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetJobsUsers() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, StreamResponseValidationError{
						field:  fmt.Sprintf("JobsUsers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, StreamResponseValidationError{
						field:  fmt.Sprintf("JobsUsers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StreamResponseValidationError{
					field:  fmt.Sprintf("JobsUsers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetUsers() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, StreamResponseValidationError{
						field:  fmt.Sprintf("Users[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, StreamResponseValidationError{
						field:  fmt.Sprintf("Users[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StreamResponseValidationError{
					field:  fmt.Sprintf("Users[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetJobsMarkers() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, StreamResponseValidationError{
						field:  fmt.Sprintf("JobsMarkers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, StreamResponseValidationError{
						field:  fmt.Sprintf("JobsMarkers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StreamResponseValidationError{
					field:  fmt.Sprintf("JobsMarkers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetMarkers() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, StreamResponseValidationError{
						field:  fmt.Sprintf("Markers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, StreamResponseValidationError{
						field:  fmt.Sprintf("Markers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StreamResponseValidationError{
					field:  fmt.Sprintf("Markers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return StreamResponseMultiError(errors)
	}

	return nil
}

// StreamResponseMultiError is an error wrapping multiple validation errors
// returned by StreamResponse.ValidateAll() if the designated constraints
// aren't met.
type StreamResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StreamResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StreamResponseMultiError) AllErrors() []error { return m }

// StreamResponseValidationError is the validation error returned by
// StreamResponse.Validate if the designated constraints aren't met.
type StreamResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamResponseValidationError) ErrorName() string { return "StreamResponseValidationError" }

// Error satisfies the builtin error interface
func (e StreamResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamResponseValidationError{}

// Validate checks the field values on CreateOrUpdateMarkerRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateOrUpdateMarkerRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateOrUpdateMarkerRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateOrUpdateMarkerRequestMultiError, or nil if none found.
func (m *CreateOrUpdateMarkerRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateOrUpdateMarkerRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetMarker() == nil {
		err := CreateOrUpdateMarkerRequestValidationError{
			field:  "Marker",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetMarker()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateOrUpdateMarkerRequestValidationError{
					field:  "Marker",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateOrUpdateMarkerRequestValidationError{
					field:  "Marker",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMarker()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateOrUpdateMarkerRequestValidationError{
				field:  "Marker",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateOrUpdateMarkerRequestMultiError(errors)
	}

	return nil
}

// CreateOrUpdateMarkerRequestMultiError is an error wrapping multiple
// validation errors returned by CreateOrUpdateMarkerRequest.ValidateAll() if
// the designated constraints aren't met.
type CreateOrUpdateMarkerRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateOrUpdateMarkerRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateOrUpdateMarkerRequestMultiError) AllErrors() []error { return m }

// CreateOrUpdateMarkerRequestValidationError is the validation error returned
// by CreateOrUpdateMarkerRequest.Validate if the designated constraints
// aren't met.
type CreateOrUpdateMarkerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOrUpdateMarkerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOrUpdateMarkerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOrUpdateMarkerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOrUpdateMarkerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOrUpdateMarkerRequestValidationError) ErrorName() string {
	return "CreateOrUpdateMarkerRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateOrUpdateMarkerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOrUpdateMarkerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOrUpdateMarkerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOrUpdateMarkerRequestValidationError{}

// Validate checks the field values on CreateOrUpdateMarkerResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateOrUpdateMarkerResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateOrUpdateMarkerResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateOrUpdateMarkerResponseMultiError, or nil if none found.
func (m *CreateOrUpdateMarkerResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateOrUpdateMarkerResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetMarker()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateOrUpdateMarkerResponseValidationError{
					field:  "Marker",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateOrUpdateMarkerResponseValidationError{
					field:  "Marker",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMarker()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateOrUpdateMarkerResponseValidationError{
				field:  "Marker",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateOrUpdateMarkerResponseMultiError(errors)
	}

	return nil
}

// CreateOrUpdateMarkerResponseMultiError is an error wrapping multiple
// validation errors returned by CreateOrUpdateMarkerResponse.ValidateAll() if
// the designated constraints aren't met.
type CreateOrUpdateMarkerResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateOrUpdateMarkerResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateOrUpdateMarkerResponseMultiError) AllErrors() []error { return m }

// CreateOrUpdateMarkerResponseValidationError is the validation error returned
// by CreateOrUpdateMarkerResponse.Validate if the designated constraints
// aren't met.
type CreateOrUpdateMarkerResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOrUpdateMarkerResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOrUpdateMarkerResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOrUpdateMarkerResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOrUpdateMarkerResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOrUpdateMarkerResponseValidationError) ErrorName() string {
	return "CreateOrUpdateMarkerResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateOrUpdateMarkerResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOrUpdateMarkerResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOrUpdateMarkerResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOrUpdateMarkerResponseValidationError{}

// Validate checks the field values on DeleteMarkerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteMarkerRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteMarkerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteMarkerRequestMultiError, or nil if none found.
func (m *DeleteMarkerRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteMarkerRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeleteMarkerRequestMultiError(errors)
	}

	return nil
}

// DeleteMarkerRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteMarkerRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteMarkerRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteMarkerRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteMarkerRequestMultiError) AllErrors() []error { return m }

// DeleteMarkerRequestValidationError is the validation error returned by
// DeleteMarkerRequest.Validate if the designated constraints aren't met.
type DeleteMarkerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteMarkerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteMarkerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteMarkerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteMarkerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteMarkerRequestValidationError) ErrorName() string {
	return "DeleteMarkerRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteMarkerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteMarkerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteMarkerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteMarkerRequestValidationError{}

// Validate checks the field values on DeleteMarkerResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteMarkerResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteMarkerResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteMarkerResponseMultiError, or nil if none found.
func (m *DeleteMarkerResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteMarkerResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteMarkerResponseMultiError(errors)
	}

	return nil
}

// DeleteMarkerResponseMultiError is an error wrapping multiple validation
// errors returned by DeleteMarkerResponse.ValidateAll() if the designated
// constraints aren't met.
type DeleteMarkerResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteMarkerResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteMarkerResponseMultiError) AllErrors() []error { return m }

// DeleteMarkerResponseValidationError is the validation error returned by
// DeleteMarkerResponse.Validate if the designated constraints aren't met.
type DeleteMarkerResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteMarkerResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteMarkerResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteMarkerResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteMarkerResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteMarkerResponseValidationError) ErrorName() string {
	return "DeleteMarkerResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteMarkerResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteMarkerResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteMarkerResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteMarkerResponseValidationError{}
