// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: services/completor/completor.proto

package completor

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

// Validate checks the field values on CompleteJobNamesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CompleteJobNamesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CompleteJobNamesRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CompleteJobNamesRequestMultiError, or nil if none found.
func (m *CompleteJobNamesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CompleteJobNamesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetSearch()) > 50 {
		err := CompleteJobNamesRequestValidationError{
			field:  "Search",
			reason: "value length must be at most 50 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CompleteJobNamesRequestMultiError(errors)
	}

	return nil
}

// CompleteJobNamesRequestMultiError is an error wrapping multiple validation
// errors returned by CompleteJobNamesRequest.ValidateAll() if the designated
// constraints aren't met.
type CompleteJobNamesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CompleteJobNamesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CompleteJobNamesRequestMultiError) AllErrors() []error { return m }

// CompleteJobNamesRequestValidationError is the validation error returned by
// CompleteJobNamesRequest.Validate if the designated constraints aren't met.
type CompleteJobNamesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CompleteJobNamesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CompleteJobNamesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CompleteJobNamesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CompleteJobNamesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CompleteJobNamesRequestValidationError) ErrorName() string {
	return "CompleteJobNamesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CompleteJobNamesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCompleteJobNamesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CompleteJobNamesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CompleteJobNamesRequestValidationError{}

// Validate checks the field values on CompleteJobNamesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CompleteJobNamesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CompleteJobNamesResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CompleteJobNamesResponseMultiError, or nil if none found.
func (m *CompleteJobNamesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CompleteJobNamesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetJobs() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CompleteJobNamesResponseValidationError{
						field:  fmt.Sprintf("Jobs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CompleteJobNamesResponseValidationError{
						field:  fmt.Sprintf("Jobs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CompleteJobNamesResponseValidationError{
					field:  fmt.Sprintf("Jobs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CompleteJobNamesResponseMultiError(errors)
	}

	return nil
}

// CompleteJobNamesResponseMultiError is an error wrapping multiple validation
// errors returned by CompleteJobNamesResponse.ValidateAll() if the designated
// constraints aren't met.
type CompleteJobNamesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CompleteJobNamesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CompleteJobNamesResponseMultiError) AllErrors() []error { return m }

// CompleteJobNamesResponseValidationError is the validation error returned by
// CompleteJobNamesResponse.Validate if the designated constraints aren't met.
type CompleteJobNamesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CompleteJobNamesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CompleteJobNamesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CompleteJobNamesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CompleteJobNamesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CompleteJobNamesResponseValidationError) ErrorName() string {
	return "CompleteJobNamesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CompleteJobNamesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCompleteJobNamesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CompleteJobNamesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CompleteJobNamesResponseValidationError{}

// Validate checks the field values on CompleteJobGradesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CompleteJobGradesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CompleteJobGradesRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CompleteJobGradesRequestMultiError, or nil if none found.
func (m *CompleteJobGradesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CompleteJobGradesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetJob()) > 20 {
		err := CompleteJobGradesRequestValidationError{
			field:  "Job",
			reason: "value length must be at most 20 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetSearch()) > 50 {
		err := CompleteJobGradesRequestValidationError{
			field:  "Search",
			reason: "value length must be at most 50 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CompleteJobGradesRequestMultiError(errors)
	}

	return nil
}

// CompleteJobGradesRequestMultiError is an error wrapping multiple validation
// errors returned by CompleteJobGradesRequest.ValidateAll() if the designated
// constraints aren't met.
type CompleteJobGradesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CompleteJobGradesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CompleteJobGradesRequestMultiError) AllErrors() []error { return m }

// CompleteJobGradesRequestValidationError is the validation error returned by
// CompleteJobGradesRequest.Validate if the designated constraints aren't met.
type CompleteJobGradesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CompleteJobGradesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CompleteJobGradesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CompleteJobGradesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CompleteJobGradesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CompleteJobGradesRequestValidationError) ErrorName() string {
	return "CompleteJobGradesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CompleteJobGradesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCompleteJobGradesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CompleteJobGradesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CompleteJobGradesRequestValidationError{}

// Validate checks the field values on CompleteJobGradesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CompleteJobGradesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CompleteJobGradesResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CompleteJobGradesResponseMultiError, or nil if none found.
func (m *CompleteJobGradesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CompleteJobGradesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetGrades() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CompleteJobGradesResponseValidationError{
						field:  fmt.Sprintf("Grades[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CompleteJobGradesResponseValidationError{
						field:  fmt.Sprintf("Grades[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CompleteJobGradesResponseValidationError{
					field:  fmt.Sprintf("Grades[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CompleteJobGradesResponseMultiError(errors)
	}

	return nil
}

// CompleteJobGradesResponseMultiError is an error wrapping multiple validation
// errors returned by CompleteJobGradesResponse.ValidateAll() if the
// designated constraints aren't met.
type CompleteJobGradesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CompleteJobGradesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CompleteJobGradesResponseMultiError) AllErrors() []error { return m }

// CompleteJobGradesResponseValidationError is the validation error returned by
// CompleteJobGradesResponse.Validate if the designated constraints aren't met.
type CompleteJobGradesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CompleteJobGradesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CompleteJobGradesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CompleteJobGradesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CompleteJobGradesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CompleteJobGradesResponseValidationError) ErrorName() string {
	return "CompleteJobGradesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CompleteJobGradesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCompleteJobGradesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CompleteJobGradesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CompleteJobGradesResponseValidationError{}

// Validate checks the field values on CompleteDocumentCategoryRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CompleteDocumentCategoryRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CompleteDocumentCategoryRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// CompleteDocumentCategoryRequestMultiError, or nil if none found.
func (m *CompleteDocumentCategoryRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CompleteDocumentCategoryRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetSearch()) > 128 {
		err := CompleteDocumentCategoryRequestValidationError{
			field:  "Search",
			reason: "value length must be at most 128 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CompleteDocumentCategoryRequestMultiError(errors)
	}

	return nil
}

// CompleteDocumentCategoryRequestMultiError is an error wrapping multiple
// validation errors returned by CompleteDocumentCategoryRequest.ValidateAll()
// if the designated constraints aren't met.
type CompleteDocumentCategoryRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CompleteDocumentCategoryRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CompleteDocumentCategoryRequestMultiError) AllErrors() []error { return m }

// CompleteDocumentCategoryRequestValidationError is the validation error
// returned by CompleteDocumentCategoryRequest.Validate if the designated
// constraints aren't met.
type CompleteDocumentCategoryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CompleteDocumentCategoryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CompleteDocumentCategoryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CompleteDocumentCategoryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CompleteDocumentCategoryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CompleteDocumentCategoryRequestValidationError) ErrorName() string {
	return "CompleteDocumentCategoryRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CompleteDocumentCategoryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCompleteDocumentCategoryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CompleteDocumentCategoryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CompleteDocumentCategoryRequestValidationError{}

// Validate checks the field values on CompleteDocumentCategoryResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *CompleteDocumentCategoryResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CompleteDocumentCategoryResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// CompleteDocumentCategoryResponseMultiError, or nil if none found.
func (m *CompleteDocumentCategoryResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CompleteDocumentCategoryResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetCategories() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CompleteDocumentCategoryResponseValidationError{
						field:  fmt.Sprintf("Categories[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CompleteDocumentCategoryResponseValidationError{
						field:  fmt.Sprintf("Categories[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CompleteDocumentCategoryResponseValidationError{
					field:  fmt.Sprintf("Categories[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CompleteDocumentCategoryResponseMultiError(errors)
	}

	return nil
}

// CompleteDocumentCategoryResponseMultiError is an error wrapping multiple
// validation errors returned by
// CompleteDocumentCategoryResponse.ValidateAll() if the designated
// constraints aren't met.
type CompleteDocumentCategoryResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CompleteDocumentCategoryResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CompleteDocumentCategoryResponseMultiError) AllErrors() []error { return m }

// CompleteDocumentCategoryResponseValidationError is the validation error
// returned by CompleteDocumentCategoryResponse.Validate if the designated
// constraints aren't met.
type CompleteDocumentCategoryResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CompleteDocumentCategoryResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CompleteDocumentCategoryResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CompleteDocumentCategoryResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CompleteDocumentCategoryResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CompleteDocumentCategoryResponseValidationError) ErrorName() string {
	return "CompleteDocumentCategoryResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CompleteDocumentCategoryResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCompleteDocumentCategoryResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CompleteDocumentCategoryResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CompleteDocumentCategoryResponseValidationError{}
