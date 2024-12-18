// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: resources/internet/ads.proto

package internet

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

// Validate checks the field values on Ad with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Ad) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Ad with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in AdMultiError, or nil if none found.
func (m *Ad) ValidateAll() error {
	return m.validate(true)
}

func (m *Ad) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AdValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AdValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AdValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Disabled

	if _, ok := AdType_name[int32(m.GetAdType())]; !ok {
		err := AdValidationError{
			field:  "AdType",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetTitle()); l < 3 || l > 255 {
		err := AdValidationError{
			field:  "Title",
			reason: "value length must be between 3 and 255 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetDescription()); l < 3 || l > 1024 {
		err := AdValidationError{
			field:  "Description",
			reason: "value length must be between 3 and 1024 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.UpdatedAt != nil {

		if all {
			switch v := interface{}(m.GetUpdatedAt()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AdValidationError{
						field:  "UpdatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AdValidationError{
						field:  "UpdatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AdValidationError{
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
					errors = append(errors, AdValidationError{
						field:  "DeletedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AdValidationError{
						field:  "DeletedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetDeletedAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AdValidationError{
					field:  "DeletedAt",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.StartsAt != nil {

		if all {
			switch v := interface{}(m.GetStartsAt()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AdValidationError{
						field:  "StartsAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AdValidationError{
						field:  "StartsAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetStartsAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AdValidationError{
					field:  "StartsAt",
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
					errors = append(errors, AdValidationError{
						field:  "EndsAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AdValidationError{
						field:  "EndsAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetEndsAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AdValidationError{
					field:  "EndsAt",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Image != nil {

		if all {
			switch v := interface{}(m.GetImage()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AdValidationError{
						field:  "Image",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AdValidationError{
						field:  "Image",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetImage()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AdValidationError{
					field:  "Image",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.ApproverId != nil {
		// no validation rules for ApproverId
	}

	if m.ApproverJob != nil {
		// no validation rules for ApproverJob
	}

	if m.CreatorId != nil {
		// no validation rules for CreatorId
	}

	if m.CreatorJob != nil {
		// no validation rules for CreatorJob
	}

	if len(errors) > 0 {
		return AdMultiError(errors)
	}

	return nil
}

// AdMultiError is an error wrapping multiple validation errors returned by
// Ad.ValidateAll() if the designated constraints aren't met.
type AdMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AdMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AdMultiError) AllErrors() []error { return m }

// AdValidationError is the validation error returned by Ad.Validate if the
// designated constraints aren't met.
type AdValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AdValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AdValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AdValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AdValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AdValidationError) ErrorName() string { return "AdValidationError" }

// Error satisfies the builtin error interface
func (e AdValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAd.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AdValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AdValidationError{}
