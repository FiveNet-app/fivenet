// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: resources/calendar/calendar.proto

package calendar

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

// Validate checks the field values on Calendar with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Calendar) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Calendar with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CalendarMultiError, or nil
// if none found.
func (m *Calendar) ValidateAll() error {
	return m.validate(true)
}

func (m *Calendar) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if l := utf8.RuneCountInString(m.GetName()); l < 3 || l > 255 {
		err := CalendarValidationError{
			field:  "Name",
			reason: "value length must be between 3 and 255 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Public

	// no validation rules for Closed

	if utf8.RuneCountInString(m.GetColor()) > 12 {
		err := CalendarValidationError{
			field:  "Color",
			reason: "value length must be at most 12 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetCreatorJob()) > 20 {
		err := CalendarValidationError{
			field:  "CreatorJob",
			reason: "value length must be at most 20 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetAccess()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CalendarValidationError{
					field:  "Access",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CalendarValidationError{
					field:  "Access",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAccess()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CalendarValidationError{
				field:  "Access",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.CreatedAt != nil {

		if all {
			switch v := interface{}(m.GetCreatedAt()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CalendarValidationError{
						field:  "CreatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CalendarValidationError{
						field:  "CreatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CalendarValidationError{
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
					errors = append(errors, CalendarValidationError{
						field:  "UpdatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CalendarValidationError{
						field:  "UpdatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CalendarValidationError{
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
					errors = append(errors, CalendarValidationError{
						field:  "DeletedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CalendarValidationError{
						field:  "DeletedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetDeletedAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CalendarValidationError{
					field:  "DeletedAt",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Job != nil {

		if utf8.RuneCountInString(m.GetJob()) > 20 {
			err := CalendarValidationError{
				field:  "Job",
				reason: "value length must be at most 20 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.Description != nil {

		if utf8.RuneCountInString(m.GetDescription()) > 512 {
			err := CalendarValidationError{
				field:  "Description",
				reason: "value length must be at most 512 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.CreatorId != nil {
		// no validation rules for CreatorId
	}

	if m.Creator != nil {

		if all {
			switch v := interface{}(m.GetCreator()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CalendarValidationError{
						field:  "Creator",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CalendarValidationError{
						field:  "Creator",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCreator()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CalendarValidationError{
					field:  "Creator",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CalendarMultiError(errors)
	}

	return nil
}

// CalendarMultiError is an error wrapping multiple validation errors returned
// by Calendar.ValidateAll() if the designated constraints aren't met.
type CalendarMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CalendarMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CalendarMultiError) AllErrors() []error { return m }

// CalendarValidationError is the validation error returned by
// Calendar.Validate if the designated constraints aren't met.
type CalendarValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CalendarValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CalendarValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CalendarValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CalendarValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CalendarValidationError) ErrorName() string { return "CalendarValidationError" }

// Error satisfies the builtin error interface
func (e CalendarValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCalendar.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CalendarValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CalendarValidationError{}

// Validate checks the field values on CalendarShort with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CalendarShort) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CalendarShort with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CalendarShortMultiError, or
// nil if none found.
func (m *CalendarShort) ValidateAll() error {
	return m.validate(true)
}

func (m *CalendarShort) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if l := utf8.RuneCountInString(m.GetName()); l < 3 || l > 255 {
		err := CalendarShortValidationError{
			field:  "Name",
			reason: "value length must be between 3 and 255 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Public

	// no validation rules for Closed

	if utf8.RuneCountInString(m.GetColor()) > 12 {
		err := CalendarShortValidationError{
			field:  "Color",
			reason: "value length must be at most 12 runes",
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
					errors = append(errors, CalendarShortValidationError{
						field:  "CreatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CalendarShortValidationError{
						field:  "CreatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CalendarShortValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Description != nil {

		if utf8.RuneCountInString(m.GetDescription()) > 512 {
			err := CalendarShortValidationError{
				field:  "Description",
				reason: "value length must be at most 512 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return CalendarShortMultiError(errors)
	}

	return nil
}

// CalendarShortMultiError is an error wrapping multiple validation errors
// returned by CalendarShort.ValidateAll() if the designated constraints
// aren't met.
type CalendarShortMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CalendarShortMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CalendarShortMultiError) AllErrors() []error { return m }

// CalendarShortValidationError is the validation error returned by
// CalendarShort.Validate if the designated constraints aren't met.
type CalendarShortValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CalendarShortValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CalendarShortValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CalendarShortValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CalendarShortValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CalendarShortValidationError) ErrorName() string { return "CalendarShortValidationError" }

// Error satisfies the builtin error interface
func (e CalendarShortValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCalendarShort.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CalendarShortValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CalendarShortValidationError{}

// Validate checks the field values on CalendarEntry with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CalendarEntry) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CalendarEntry with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CalendarEntryMultiError, or
// nil if none found.
func (m *CalendarEntry) ValidateAll() error {
	return m.validate(true)
}

func (m *CalendarEntry) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for CalendarId

	if all {
		switch v := interface{}(m.GetStartTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CalendarEntryValidationError{
					field:  "StartTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CalendarEntryValidationError{
					field:  "StartTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStartTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CalendarEntryValidationError{
				field:  "StartTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if l := utf8.RuneCountInString(m.GetTitle()); l < 3 || l > 512 {
		err := CalendarEntryValidationError{
			field:  "Title",
			reason: "value length must be between 3 and 512 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetContent()) < 20 {
		err := CalendarEntryValidationError{
			field:  "Content",
			reason: "value length must be at least 20 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetContent()) > 1000000 {
		err := CalendarEntryValidationError{
			field:  "Content",
			reason: "value length must be at most 1000000 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetCreatorJob()) > 20 {
		err := CalendarEntryValidationError{
			field:  "CreatorJob",
			reason: "value length must be at most 20 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetAccess()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CalendarEntryValidationError{
					field:  "Access",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CalendarEntryValidationError{
					field:  "Access",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAccess()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CalendarEntryValidationError{
				field:  "Access",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.CreatedAt != nil {

		if all {
			switch v := interface{}(m.GetCreatedAt()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CalendarEntryValidationError{
						field:  "CreatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CalendarEntryValidationError{
						field:  "CreatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CalendarEntryValidationError{
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
					errors = append(errors, CalendarEntryValidationError{
						field:  "UpdatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CalendarEntryValidationError{
						field:  "UpdatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CalendarEntryValidationError{
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
					errors = append(errors, CalendarEntryValidationError{
						field:  "DeletedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CalendarEntryValidationError{
						field:  "DeletedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetDeletedAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CalendarEntryValidationError{
					field:  "DeletedAt",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Calendar != nil {

		if all {
			switch v := interface{}(m.GetCalendar()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CalendarEntryValidationError{
						field:  "Calendar",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CalendarEntryValidationError{
						field:  "Calendar",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCalendar()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CalendarEntryValidationError{
					field:  "Calendar",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Job != nil {

		if utf8.RuneCountInString(m.GetJob()) > 20 {
			err := CalendarEntryValidationError{
				field:  "Job",
				reason: "value length must be at most 20 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.EndTime != nil {

		if all {
			switch v := interface{}(m.GetEndTime()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CalendarEntryValidationError{
						field:  "EndTime",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CalendarEntryValidationError{
						field:  "EndTime",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetEndTime()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CalendarEntryValidationError{
					field:  "EndTime",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.RsvpOpen != nil {
		// no validation rules for RsvpOpen
	}

	if m.CreatorId != nil {
		// no validation rules for CreatorId
	}

	if m.Creator != nil {

		if all {
			switch v := interface{}(m.GetCreator()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CalendarEntryValidationError{
						field:  "Creator",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CalendarEntryValidationError{
						field:  "Creator",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCreator()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CalendarEntryValidationError{
					field:  "Creator",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Recurring != nil {

		if all {
			switch v := interface{}(m.GetRecurring()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CalendarEntryValidationError{
						field:  "Recurring",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CalendarEntryValidationError{
						field:  "Recurring",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetRecurring()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CalendarEntryValidationError{
					field:  "Recurring",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CalendarEntryMultiError(errors)
	}

	return nil
}

// CalendarEntryMultiError is an error wrapping multiple validation errors
// returned by CalendarEntry.ValidateAll() if the designated constraints
// aren't met.
type CalendarEntryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CalendarEntryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CalendarEntryMultiError) AllErrors() []error { return m }

// CalendarEntryValidationError is the validation error returned by
// CalendarEntry.Validate if the designated constraints aren't met.
type CalendarEntryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CalendarEntryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CalendarEntryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CalendarEntryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CalendarEntryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CalendarEntryValidationError) ErrorName() string { return "CalendarEntryValidationError" }

// Error satisfies the builtin error interface
func (e CalendarEntryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCalendarEntry.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CalendarEntryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CalendarEntryValidationError{}

// Validate checks the field values on CalendarEntryRecurring with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CalendarEntryRecurring) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CalendarEntryRecurring with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CalendarEntryRecurringMultiError, or nil if none found.
func (m *CalendarEntryRecurring) ValidateAll() error {
	return m.validate(true)
}

func (m *CalendarEntryRecurring) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetStartedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CalendarEntryRecurringValidationError{
					field:  "StartedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CalendarEntryRecurringValidationError{
					field:  "StartedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStartedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CalendarEntryRecurringValidationError{
				field:  "StartedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Weekly

	if len(errors) > 0 {
		return CalendarEntryRecurringMultiError(errors)
	}

	return nil
}

// CalendarEntryRecurringMultiError is an error wrapping multiple validation
// errors returned by CalendarEntryRecurring.ValidateAll() if the designated
// constraints aren't met.
type CalendarEntryRecurringMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CalendarEntryRecurringMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CalendarEntryRecurringMultiError) AllErrors() []error { return m }

// CalendarEntryRecurringValidationError is the validation error returned by
// CalendarEntryRecurring.Validate if the designated constraints aren't met.
type CalendarEntryRecurringValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CalendarEntryRecurringValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CalendarEntryRecurringValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CalendarEntryRecurringValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CalendarEntryRecurringValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CalendarEntryRecurringValidationError) ErrorName() string {
	return "CalendarEntryRecurringValidationError"
}

// Error satisfies the builtin error interface
func (e CalendarEntryRecurringValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCalendarEntryRecurring.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CalendarEntryRecurringValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CalendarEntryRecurringValidationError{}

// Validate checks the field values on CalendarSub with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CalendarSub) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CalendarSub with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CalendarSubMultiError, or
// nil if none found.
func (m *CalendarSub) ValidateAll() error {
	return m.validate(true)
}

func (m *CalendarSub) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for CalendarId

	// no validation rules for UserId

	// no validation rules for Confirmed

	// no validation rules for Muted

	if m.EntryId != nil {
		// no validation rules for EntryId
	}

	if m.User != nil {

		if all {
			switch v := interface{}(m.GetUser()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CalendarSubValidationError{
						field:  "User",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CalendarSubValidationError{
						field:  "User",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CalendarSubValidationError{
					field:  "User",
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
					errors = append(errors, CalendarSubValidationError{
						field:  "CreatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CalendarSubValidationError{
						field:  "CreatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CalendarSubValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CalendarSubMultiError(errors)
	}

	return nil
}

// CalendarSubMultiError is an error wrapping multiple validation errors
// returned by CalendarSub.ValidateAll() if the designated constraints aren't met.
type CalendarSubMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CalendarSubMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CalendarSubMultiError) AllErrors() []error { return m }

// CalendarSubValidationError is the validation error returned by
// CalendarSub.Validate if the designated constraints aren't met.
type CalendarSubValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CalendarSubValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CalendarSubValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CalendarSubValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CalendarSubValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CalendarSubValidationError) ErrorName() string { return "CalendarSubValidationError" }

// Error satisfies the builtin error interface
func (e CalendarSubValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCalendarSub.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CalendarSubValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CalendarSubValidationError{}

// Validate checks the field values on CalendarEntryRSVP with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CalendarEntryRSVP) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CalendarEntryRSVP with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CalendarEntryRSVPMultiError, or nil if none found.
func (m *CalendarEntryRSVP) ValidateAll() error {
	return m.validate(true)
}

func (m *CalendarEntryRSVP) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for EntryId

	if m.GetUserId() <= 0 {
		err := CalendarEntryRSVPValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := RsvpResponses_name[int32(m.GetResponse())]; !ok {
		err := CalendarEntryRSVPValidationError{
			field:  "Response",
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
					errors = append(errors, CalendarEntryRSVPValidationError{
						field:  "CreatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CalendarEntryRSVPValidationError{
						field:  "CreatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CalendarEntryRSVPValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.User != nil {

		if all {
			switch v := interface{}(m.GetUser()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CalendarEntryRSVPValidationError{
						field:  "User",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CalendarEntryRSVPValidationError{
						field:  "User",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CalendarEntryRSVPValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CalendarEntryRSVPMultiError(errors)
	}

	return nil
}

// CalendarEntryRSVPMultiError is an error wrapping multiple validation errors
// returned by CalendarEntryRSVP.ValidateAll() if the designated constraints
// aren't met.
type CalendarEntryRSVPMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CalendarEntryRSVPMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CalendarEntryRSVPMultiError) AllErrors() []error { return m }

// CalendarEntryRSVPValidationError is the validation error returned by
// CalendarEntryRSVP.Validate if the designated constraints aren't met.
type CalendarEntryRSVPValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CalendarEntryRSVPValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CalendarEntryRSVPValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CalendarEntryRSVPValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CalendarEntryRSVPValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CalendarEntryRSVPValidationError) ErrorName() string {
	return "CalendarEntryRSVPValidationError"
}

// Error satisfies the builtin error interface
func (e CalendarEntryRSVPValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCalendarEntryRSVP.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CalendarEntryRSVPValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CalendarEntryRSVPValidationError{}
