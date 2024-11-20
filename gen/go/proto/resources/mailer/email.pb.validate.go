// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: resources/mailer/email.proto

package mailer

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

// Validate checks the field values on Email with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Email) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Email with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in EmailMultiError, or nil if none found.
func (m *Email) ValidateAll() error {
	return m.validate(true)
}

func (m *Email) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EmailValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EmailValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EmailValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Disabled

	if l := utf8.RuneCountInString(m.GetEmail()); l < 6 || l > 80 {
		err := EmailValidationError{
			field:  "Email",
			reason: "value length must be between 6 and 80 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Internal

	if all {
		switch v := interface{}(m.GetAccess()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EmailValidationError{
					field:  "Access",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EmailValidationError{
					field:  "Access",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAccess()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EmailValidationError{
				field:  "Access",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.UpdatedAt != nil {

		if all {
			switch v := interface{}(m.GetUpdatedAt()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EmailValidationError{
						field:  "UpdatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EmailValidationError{
						field:  "UpdatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EmailValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.DeletedAt != nil {

		if all {
			switch v := interface{}(m.GetDeletedAt()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EmailValidationError{
						field:  "DeletedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EmailValidationError{
						field:  "DeletedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetDeletedAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EmailValidationError{
					field:  "DeletedAt",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Job != nil {

		if utf8.RuneCountInString(m.GetJob()) > 40 {
			err := EmailValidationError{
				field:  "Job",
				reason: "value length must be at most 40 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.UserId != nil {

		if m.GetUserId() <= 0 {
			err := EmailValidationError{
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
					errors = append(errors, EmailValidationError{
						field:  "User",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EmailValidationError{
						field:  "User",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EmailValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.EmailChanged != nil {

		if all {
			switch v := interface{}(m.GetEmailChanged()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EmailValidationError{
						field:  "EmailChanged",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EmailValidationError{
						field:  "EmailChanged",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetEmailChanged()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EmailValidationError{
					field:  "EmailChanged",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Label != nil {

		if utf8.RuneCountInString(m.GetLabel()) > 128 {
			err := EmailValidationError{
				field:  "Label",
				reason: "value length must be at most 128 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.Settings != nil {

		if all {
			switch v := interface{}(m.GetSettings()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EmailValidationError{
						field:  "Settings",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EmailValidationError{
						field:  "Settings",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetSettings()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EmailValidationError{
					field:  "Settings",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return EmailMultiError(errors)
	}

	return nil
}

// EmailMultiError is an error wrapping multiple validation errors returned by
// Email.ValidateAll() if the designated constraints aren't met.
type EmailMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EmailMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EmailMultiError) AllErrors() []error { return m }

// EmailValidationError is the validation error returned by Email.Validate if
// the designated constraints aren't met.
type EmailValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmailValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmailValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmailValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmailValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmailValidationError) ErrorName() string { return "EmailValidationError" }

// Error satisfies the builtin error interface
func (e EmailValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmail.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmailValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmailValidationError{}
