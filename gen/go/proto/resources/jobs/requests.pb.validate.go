// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: resources/jobs/requests.proto

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

// Validate checks the field values on RequestEntry with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RequestEntry) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RequestEntry with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RequestEntryMultiError, or
// nil if none found.
func (m *RequestEntry) ValidateAll() error {
	return m.validate(true)
}

func (m *RequestEntry) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if utf8.RuneCountInString(m.GetJob()) > 20 {
		err := RequestEntryValidationError{
			field:  "Job",
			reason: "value length must be at most 20 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := REQUEST_TYPE_name[int32(m.GetType())]; !ok {
		err := RequestEntryValidationError{
			field:  "Type",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetMessage()); l < 3 || l > 2048 {
		err := RequestEntryValidationError{
			field:  "Message",
			reason: "value length must be between 3 and 2048 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetCreatorId() <= 0 {
		err := RequestEntryValidationError{
			field:  "CreatorId",
			reason: "value must be greater than 0",
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
					errors = append(errors, RequestEntryValidationError{
						field:  "CreatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RequestEntryValidationError{
						field:  "CreatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RequestEntryValidationError{
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
					errors = append(errors, RequestEntryValidationError{
						field:  "UpdatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RequestEntryValidationError{
						field:  "UpdatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RequestEntryValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.BeginsAt != nil {

		if all {
			switch v := interface{}(m.GetBeginsAt()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RequestEntryValidationError{
						field:  "BeginsAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RequestEntryValidationError{
						field:  "BeginsAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetBeginsAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RequestEntryValidationError{
					field:  "BeginsAt",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.EndsAt != nil {

		if all {
			switch v := interface{}(m.GetEndsAt()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RequestEntryValidationError{
						field:  "EndsAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RequestEntryValidationError{
						field:  "EndsAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetEndsAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RequestEntryValidationError{
					field:  "EndsAt",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Status != nil {
		// no validation rules for Status
	}

	if m.Creator != nil {

		if all {
			switch v := interface{}(m.GetCreator()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RequestEntryValidationError{
						field:  "Creator",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RequestEntryValidationError{
						field:  "Creator",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCreator()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RequestEntryValidationError{
					field:  "Creator",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Approved != nil {
		// no validation rules for Approved
	}

	if m.ApproverUserId != nil {
		// no validation rules for ApproverUserId
	}

	if m.ApproverUser != nil {

		if all {
			switch v := interface{}(m.GetApproverUser()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RequestEntryValidationError{
						field:  "ApproverUser",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RequestEntryValidationError{
						field:  "ApproverUser",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetApproverUser()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RequestEntryValidationError{
					field:  "ApproverUser",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return RequestEntryMultiError(errors)
	}

	return nil
}

// RequestEntryMultiError is an error wrapping multiple validation errors
// returned by RequestEntry.ValidateAll() if the designated constraints aren't met.
type RequestEntryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RequestEntryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RequestEntryMultiError) AllErrors() []error { return m }

// RequestEntryValidationError is the validation error returned by
// RequestEntry.Validate if the designated constraints aren't met.
type RequestEntryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RequestEntryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RequestEntryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RequestEntryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RequestEntryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RequestEntryValidationError) ErrorName() string { return "RequestEntryValidationError" }

// Error satisfies the builtin error interface
func (e RequestEntryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRequestEntry.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RequestEntryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RequestEntryValidationError{}
