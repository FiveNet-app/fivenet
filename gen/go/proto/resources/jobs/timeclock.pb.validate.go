// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: resources/jobs/timeclock.proto

package jobs

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

// Validate checks the field values on PlayerTime with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PlayerTime) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PlayerTime with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PlayerTimeMultiError, or
// nil if none found.
func (m *PlayerTime) ValidateAll() error {
	return m.validate(true)
}

func (m *PlayerTime) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if utf8.RuneCountInString(m.GetJob()) > 20 {
		err := PlayerTimeValidationError{
			field:  "Job",
			reason: "value length must be at most 20 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetDate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PlayerTimeValidationError{
					field:  "Date",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PlayerTimeValidationError{
					field:  "Date",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PlayerTimeValidationError{
				field:  "Date",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for UserId

	// no validation rules for Minutes

	if m.User != nil {

		if all {
			switch v := interface{}(m.GetUser()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PlayerTimeValidationError{
						field:  "User",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PlayerTimeValidationError{
						field:  "User",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PlayerTimeValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return PlayerTimeMultiError(errors)
	}

	return nil
}

// PlayerTimeMultiError is an error wrapping multiple validation errors
// returned by PlayerTime.ValidateAll() if the designated constraints aren't met.
type PlayerTimeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PlayerTimeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PlayerTimeMultiError) AllErrors() []error { return m }

// PlayerTimeValidationError is the validation error returned by
// PlayerTime.Validate if the designated constraints aren't met.
type PlayerTimeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlayerTimeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlayerTimeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlayerTimeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlayerTimeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlayerTimeValidationError) ErrorName() string { return "PlayerTimeValidationError" }

// Error satisfies the builtin error interface
func (e PlayerTimeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlayerTime.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlayerTimeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlayerTimeValidationError{}

// Validate checks the field values on TimeclockEntry with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TimeclockEntry) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TimeclockEntry with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TimeclockEntryMultiError,
// or nil if none found.
func (m *TimeclockEntry) ValidateAll() error {
	return m.validate(true)
}

func (m *TimeclockEntry) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return TimeclockEntryMultiError(errors)
	}

	return nil
}

// TimeclockEntryMultiError is an error wrapping multiple validation errors
// returned by TimeclockEntry.ValidateAll() if the designated constraints
// aren't met.
type TimeclockEntryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TimeclockEntryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TimeclockEntryMultiError) AllErrors() []error { return m }

// TimeclockEntryValidationError is the validation error returned by
// TimeclockEntry.Validate if the designated constraints aren't met.
type TimeclockEntryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TimeclockEntryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TimeclockEntryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TimeclockEntryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TimeclockEntryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TimeclockEntryValidationError) ErrorName() string { return "TimeclockEntryValidationError" }

// Error satisfies the builtin error interface
func (e TimeclockEntryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTimeclockEntry.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TimeclockEntryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TimeclockEntryValidationError{}
