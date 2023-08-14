// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: resources/dispatch/dispatches.proto

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

	if utf8.RuneCountInString(m.GetJob()) > 20 {
		err := DispatchValidationError{
			field:  "Job",
			reason: "value length must be at most 20 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetMessage()) > 255 {
		err := DispatchValidationError{
			field:  "Message",
			reason: "value length must be at most 255 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for X

	// no validation rules for Y

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

		if all {
			switch v := interface{}(m.GetStatus()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DispatchValidationError{
						field:  "Status",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DispatchValidationError{
						field:  "Status",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetStatus()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DispatchValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Description != nil {

		if utf8.RuneCountInString(m.GetDescription()) > 1024 {
			err := DispatchValidationError{
				field:  "Description",
				reason: "value length must be at most 1024 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.Attributes != nil {

		if all {
			switch v := interface{}(m.GetAttributes()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DispatchValidationError{
						field:  "Attributes",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DispatchValidationError{
						field:  "Attributes",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetAttributes()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DispatchValidationError{
					field:  "Attributes",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Anon != nil {
		// no validation rules for Anon
	}

	if m.UserId != nil {
		// no validation rules for UserId
	}

	if m.User != nil {

		if all {
			switch v := interface{}(m.GetUser()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DispatchValidationError{
						field:  "User",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DispatchValidationError{
						field:  "User",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DispatchValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

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

// Validate checks the field values on Attributes with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Attributes) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Attributes with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AttributesMultiError, or
// nil if none found.
func (m *Attributes) ValidateAll() error {
	return m.validate(true)
}

func (m *Attributes) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return AttributesMultiError(errors)
	}

	return nil
}

// AttributesMultiError is an error wrapping multiple validation errors
// returned by Attributes.ValidateAll() if the designated constraints aren't met.
type AttributesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AttributesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AttributesMultiError) AllErrors() []error { return m }

// AttributesValidationError is the validation error returned by
// Attributes.Validate if the designated constraints aren't met.
type AttributesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttributesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttributesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AttributesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttributesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttributesValidationError) ErrorName() string { return "AttributesValidationError" }

// Error satisfies the builtin error interface
func (e AttributesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttributes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttributesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttributesValidationError{}

// Validate checks the field values on DispatchAssignments with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DispatchAssignments) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DispatchAssignments with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DispatchAssignmentsMultiError, or nil if none found.
func (m *DispatchAssignments) ValidateAll() error {
	return m.validate(true)
}

func (m *DispatchAssignments) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for DispatchId

	// no validation rules for Job

	for idx, item := range m.GetUnits() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DispatchAssignmentsValidationError{
						field:  fmt.Sprintf("Units[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DispatchAssignmentsValidationError{
						field:  fmt.Sprintf("Units[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DispatchAssignmentsValidationError{
					field:  fmt.Sprintf("Units[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return DispatchAssignmentsMultiError(errors)
	}

	return nil
}

// DispatchAssignmentsMultiError is an error wrapping multiple validation
// errors returned by DispatchAssignments.ValidateAll() if the designated
// constraints aren't met.
type DispatchAssignmentsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DispatchAssignmentsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DispatchAssignmentsMultiError) AllErrors() []error { return m }

// DispatchAssignmentsValidationError is the validation error returned by
// DispatchAssignments.Validate if the designated constraints aren't met.
type DispatchAssignmentsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DispatchAssignmentsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DispatchAssignmentsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DispatchAssignmentsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DispatchAssignmentsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DispatchAssignmentsValidationError) ErrorName() string {
	return "DispatchAssignmentsValidationError"
}

// Error satisfies the builtin error interface
func (e DispatchAssignmentsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDispatchAssignments.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DispatchAssignmentsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DispatchAssignmentsValidationError{}

// Validate checks the field values on DispatchAssignment with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DispatchAssignment) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DispatchAssignment with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DispatchAssignmentMultiError, or nil if none found.
func (m *DispatchAssignment) ValidateAll() error {
	return m.validate(true)
}

func (m *DispatchAssignment) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for DispatchId

	// no validation rules for UnitId

	if m.Unit != nil {

		if all {
			switch v := interface{}(m.GetUnit()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DispatchAssignmentValidationError{
						field:  "Unit",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DispatchAssignmentValidationError{
						field:  "Unit",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetUnit()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DispatchAssignmentValidationError{
					field:  "Unit",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.CreatedAt != nil {

		if all {
			switch v := interface{}(m.GetCreatedAt()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DispatchAssignmentValidationError{
						field:  "CreatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DispatchAssignmentValidationError{
						field:  "CreatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DispatchAssignmentValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.ExpiresAt != nil {

		if all {
			switch v := interface{}(m.GetExpiresAt()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DispatchAssignmentValidationError{
						field:  "ExpiresAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DispatchAssignmentValidationError{
						field:  "ExpiresAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetExpiresAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DispatchAssignmentValidationError{
					field:  "ExpiresAt",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return DispatchAssignmentMultiError(errors)
	}

	return nil
}

// DispatchAssignmentMultiError is an error wrapping multiple validation errors
// returned by DispatchAssignment.ValidateAll() if the designated constraints
// aren't met.
type DispatchAssignmentMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DispatchAssignmentMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DispatchAssignmentMultiError) AllErrors() []error { return m }

// DispatchAssignmentValidationError is the validation error returned by
// DispatchAssignment.Validate if the designated constraints aren't met.
type DispatchAssignmentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DispatchAssignmentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DispatchAssignmentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DispatchAssignmentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DispatchAssignmentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DispatchAssignmentValidationError) ErrorName() string {
	return "DispatchAssignmentValidationError"
}

// Error satisfies the builtin error interface
func (e DispatchAssignmentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDispatchAssignment.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DispatchAssignmentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DispatchAssignmentValidationError{}

// Validate checks the field values on DispatchStatus with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DispatchStatus) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DispatchStatus with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DispatchStatusMultiError,
// or nil if none found.
func (m *DispatchStatus) ValidateAll() error {
	return m.validate(true)
}

func (m *DispatchStatus) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for DispatchId

	if _, ok := DISPATCH_STATUS_name[int32(m.GetStatus())]; !ok {
		err := DispatchStatusValidationError{
			field:  "Status",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.CreatedAt != nil {

		if all {
			switch v := interface{}(m.GetCreatedAt()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DispatchStatusValidationError{
						field:  "CreatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DispatchStatusValidationError{
						field:  "CreatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DispatchStatusValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.UnitId != nil {
		// no validation rules for UnitId
	}

	if m.Unit != nil {

		if all {
			switch v := interface{}(m.GetUnit()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DispatchStatusValidationError{
						field:  "Unit",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DispatchStatusValidationError{
						field:  "Unit",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetUnit()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DispatchStatusValidationError{
					field:  "Unit",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Reason != nil {

		if utf8.RuneCountInString(m.GetReason()) > 255 {
			err := DispatchStatusValidationError{
				field:  "Reason",
				reason: "value length must be at most 255 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.Code != nil {

		if utf8.RuneCountInString(m.GetCode()) > 20 {
			err := DispatchStatusValidationError{
				field:  "Code",
				reason: "value length must be at most 20 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.UserId != nil {

		if m.GetUserId() <= 0 {
			err := DispatchStatusValidationError{
				field:  "UserId",
				reason: "value must be greater than 0",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.User != nil {

		if all {
			switch v := interface{}(m.GetUser()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DispatchStatusValidationError{
						field:  "User",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DispatchStatusValidationError{
						field:  "User",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DispatchStatusValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.X != nil {
		// no validation rules for X
	}

	if m.Y != nil {
		// no validation rules for Y
	}

	if len(errors) > 0 {
		return DispatchStatusMultiError(errors)
	}

	return nil
}

// DispatchStatusMultiError is an error wrapping multiple validation errors
// returned by DispatchStatus.ValidateAll() if the designated constraints
// aren't met.
type DispatchStatusMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DispatchStatusMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DispatchStatusMultiError) AllErrors() []error { return m }

// DispatchStatusValidationError is the validation error returned by
// DispatchStatus.Validate if the designated constraints aren't met.
type DispatchStatusValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DispatchStatusValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DispatchStatusValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DispatchStatusValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DispatchStatusValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DispatchStatusValidationError) ErrorName() string { return "DispatchStatusValidationError" }

// Error satisfies the builtin error interface
func (e DispatchStatusValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDispatchStatus.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DispatchStatusValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DispatchStatusValidationError{}
