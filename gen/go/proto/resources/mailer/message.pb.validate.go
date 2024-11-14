// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: resources/mailer/message.proto

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

// Validate checks the field values on Message with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Message) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Message with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in MessageMultiError, or nil if none found.
func (m *Message) ValidateAll() error {
	return m.validate(true)
}

func (m *Message) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for ThreadId

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MessageValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MessageValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MessageValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if l := utf8.RuneCountInString(m.GetTitle()); l < 3 || l > 255 {
		err := MessageValidationError{
			field:  "Title",
			reason: "value length must be between 3 and 255 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetContent()); l < 3 || l > 8192 {
		err := MessageValidationError{
			field:  "Content",
			reason: "value length must be between 3 and 8192 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.SenderEmailId != nil {
		// no validation rules for SenderEmailId
	}

	if m.SenderEmail != nil {

		if all {
			switch v := interface{}(m.GetSenderEmail()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MessageValidationError{
						field:  "SenderEmail",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MessageValidationError{
						field:  "SenderEmail",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetSenderEmail()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MessageValidationError{
					field:  "SenderEmail",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.SenderUserId != nil {
		// no validation rules for SenderUserId
	}

	if m.SenderUser != nil {

		if all {
			switch v := interface{}(m.GetSenderUser()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MessageValidationError{
						field:  "SenderUser",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MessageValidationError{
						field:  "SenderUser",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetSenderUser()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MessageValidationError{
					field:  "SenderUser",
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
					errors = append(errors, MessageValidationError{
						field:  "UpdatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MessageValidationError{
						field:  "UpdatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MessageValidationError{
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
					errors = append(errors, MessageValidationError{
						field:  "DeletedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MessageValidationError{
						field:  "DeletedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetDeletedAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MessageValidationError{
					field:  "DeletedAt",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Data != nil {

		if all {
			switch v := interface{}(m.GetData()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MessageValidationError{
						field:  "Data",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MessageValidationError{
						field:  "Data",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MessageValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return MessageMultiError(errors)
	}

	return nil
}

// MessageMultiError is an error wrapping multiple validation errors returned
// by Message.ValidateAll() if the designated constraints aren't met.
type MessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MessageMultiError) AllErrors() []error { return m }

// MessageValidationError is the validation error returned by Message.Validate
// if the designated constraints aren't met.
type MessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MessageValidationError) ErrorName() string { return "MessageValidationError" }

// Error satisfies the builtin error interface
func (e MessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MessageValidationError{}

// Validate checks the field values on MessageData with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MessageData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MessageData with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MessageDataMultiError, or
// nil if none found.
func (m *MessageData) ValidateAll() error {
	return m.validate(true)
}

func (m *MessageData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetEntry()) > 3 {
		err := MessageDataValidationError{
			field:  "Entry",
			reason: "value must contain no more than 3 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetEntry() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MessageDataValidationError{
						field:  fmt.Sprintf("Entry[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MessageDataValidationError{
						field:  fmt.Sprintf("Entry[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MessageDataValidationError{
					field:  fmt.Sprintf("Entry[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return MessageDataMultiError(errors)
	}

	return nil
}

// MessageDataMultiError is an error wrapping multiple validation errors
// returned by MessageData.ValidateAll() if the designated constraints aren't met.
type MessageDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MessageDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MessageDataMultiError) AllErrors() []error { return m }

// MessageDataValidationError is the validation error returned by
// MessageData.Validate if the designated constraints aren't met.
type MessageDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MessageDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MessageDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MessageDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MessageDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MessageDataValidationError) ErrorName() string { return "MessageDataValidationError" }

// Error satisfies the builtin error interface
func (e MessageDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMessageData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MessageDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MessageDataValidationError{}

// Validate checks the field values on MessageDataEntry with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *MessageDataEntry) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MessageDataEntry with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MessageDataEntryMultiError, or nil if none found.
func (m *MessageDataEntry) ValidateAll() error {
	return m.validate(true)
}

func (m *MessageDataEntry) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return MessageDataEntryMultiError(errors)
	}

	return nil
}

// MessageDataEntryMultiError is an error wrapping multiple validation errors
// returned by MessageDataEntry.ValidateAll() if the designated constraints
// aren't met.
type MessageDataEntryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MessageDataEntryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MessageDataEntryMultiError) AllErrors() []error { return m }

// MessageDataEntryValidationError is the validation error returned by
// MessageDataEntry.Validate if the designated constraints aren't met.
type MessageDataEntryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MessageDataEntryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MessageDataEntryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MessageDataEntryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MessageDataEntryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MessageDataEntryValidationError) ErrorName() string { return "MessageDataEntryValidationError" }

// Error satisfies the builtin error interface
func (e MessageDataEntryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMessageDataEntry.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MessageDataEntryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MessageDataEntryValidationError{}
