// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: resources/dispatch/dispatch.proto

package dispatch

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

// Validate checks the field values on Unit with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Unit) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Unit with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in UnitMultiError, or nil if none found.
func (m *Unit) ValidateAll() error {
	return m.validate(true)
}

func (m *Unit) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if utf8.RuneCountInString(m.GetName()) > 20 {
		err := UnitValidationError{
			field:  "Name",
			reason: "value length must be at most 20 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetInitials()) > 4 {
		err := UnitValidationError{
			field:  "Initials",
			reason: "value length must be at most 4 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetAssigned() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UnitValidationError{
						field:  fmt.Sprintf("Assigned[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UnitValidationError{
						field:  fmt.Sprintf("Assigned[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UnitValidationError{
					field:  fmt.Sprintf("Assigned[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Description != nil {

		if utf8.RuneCountInString(m.GetDescription()) > 255 {
			err := UnitValidationError{
				field:  "Description",
				reason: "value length must be at most 255 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.Status != nil {
		// no validation rules for Status
	}

	if m.Reason != nil {
		// no validation rules for Reason
	}

	if len(errors) > 0 {
		return UnitMultiError(errors)
	}

	return nil
}

// UnitMultiError is an error wrapping multiple validation errors returned by
// Unit.ValidateAll() if the designated constraints aren't met.
type UnitMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UnitMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UnitMultiError) AllErrors() []error { return m }

// UnitValidationError is the validation error returned by Unit.Validate if the
// designated constraints aren't met.
type UnitValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UnitValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UnitValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UnitValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UnitValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UnitValidationError) ErrorName() string { return "UnitValidationError" }

// Error satisfies the builtin error interface
func (e UnitValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUnit.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UnitValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UnitValidationError{}

// Validate checks the field values on Dispatch with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Dispatch) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Dispatch with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DispatchMultiError, or nil
// if none found.
func (m *Dispatch) ValidateAll() error {
	return m.validate(true)
}

func (m *Dispatch) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Title

	for idx, item := range m.GetUnits() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DispatchValidationError{
						field:  fmt.Sprintf("Units[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DispatchValidationError{
						field:  fmt.Sprintf("Units[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DispatchValidationError{
					field:  fmt.Sprintf("Units[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Attributes

	if m.CreatedAt != nil {

		if all {
			switch v := interface{}(m.GetCreatedAt()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DispatchValidationError{
						field:  "CreatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DispatchValidationError{
						field:  "CreatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DispatchValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.UpdatedAt != nil {

		if all {
			switch v := interface{}(m.GetUpdatedAt()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DispatchValidationError{
						field:  "UpdatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DispatchValidationError{
						field:  "UpdatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DispatchValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Status != nil {
		// no validation rules for Status
	}

	if m.Description != nil {
		// no validation rules for Description
	}

	if len(errors) > 0 {
		return DispatchMultiError(errors)
	}

	return nil
}

// DispatchMultiError is an error wrapping multiple validation errors returned
// by Dispatch.ValidateAll() if the designated constraints aren't met.
type DispatchMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DispatchMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DispatchMultiError) AllErrors() []error { return m }

// DispatchValidationError is the validation error returned by
// Dispatch.Validate if the designated constraints aren't met.
type DispatchValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DispatchValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DispatchValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DispatchValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DispatchValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DispatchValidationError) ErrorName() string { return "DispatchValidationError" }

// Error satisfies the builtin error interface
func (e DispatchValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDispatch.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DispatchValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DispatchValidationError{}

// Validate checks the field values on DispatchLog with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DispatchLog) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DispatchLog with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DispatchLogMultiError, or
// nil if none found.
func (m *DispatchLog) ValidateAll() error {
	return m.validate(true)
}

func (m *DispatchLog) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for DispatchId

	if len(errors) > 0 {
		return DispatchLogMultiError(errors)
	}

	return nil
}

// DispatchLogMultiError is an error wrapping multiple validation errors
// returned by DispatchLog.ValidateAll() if the designated constraints aren't met.
type DispatchLogMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DispatchLogMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DispatchLogMultiError) AllErrors() []error { return m }

// DispatchLogValidationError is the validation error returned by
// DispatchLog.Validate if the designated constraints aren't met.
type DispatchLogValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DispatchLogValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DispatchLogValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DispatchLogValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DispatchLogValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DispatchLogValidationError) ErrorName() string { return "DispatchLogValidationError" }

// Error satisfies the builtin error interface
func (e DispatchLogValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDispatchLog.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DispatchLogValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DispatchLogValidationError{}
