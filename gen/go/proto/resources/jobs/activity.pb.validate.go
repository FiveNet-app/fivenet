// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: resources/jobs/activity.proto

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

// Validate checks the field values on JobsUserActivity with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *JobsUserActivity) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on JobsUserActivity with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// JobsUserActivityMultiError, or nil if none found.
func (m *JobsUserActivity) ValidateAll() error {
	return m.validate(true)
}

func (m *JobsUserActivity) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if utf8.RuneCountInString(m.GetJob()) > 20 {
		err := JobsUserActivityValidationError{
			field:  "Job",
			reason: "value length must be at most 20 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetTargetUserId() <= 0 {
		err := JobsUserActivityValidationError{
			field:  "TargetUserId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetTargetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, JobsUserActivityValidationError{
					field:  "TargetUser",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, JobsUserActivityValidationError{
					field:  "TargetUser",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTargetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return JobsUserActivityValidationError{
				field:  "TargetUser",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ActivityType

	if utf8.RuneCountInString(m.GetReason()) > 255 {
		err := JobsUserActivityValidationError{
			field:  "Reason",
			reason: "value length must be at most 255 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, JobsUserActivityValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, JobsUserActivityValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return JobsUserActivityValidationError{
				field:  "Data",
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
					errors = append(errors, JobsUserActivityValidationError{
						field:  "CreatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, JobsUserActivityValidationError{
						field:  "CreatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return JobsUserActivityValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.SourceUserId != nil {

		if m.GetSourceUserId() <= 0 {
			err := JobsUserActivityValidationError{
				field:  "SourceUserId",
				reason: "value must be greater than 0",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.SourceUser != nil {

		if all {
			switch v := interface{}(m.GetSourceUser()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, JobsUserActivityValidationError{
						field:  "SourceUser",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, JobsUserActivityValidationError{
						field:  "SourceUser",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetSourceUser()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return JobsUserActivityValidationError{
					field:  "SourceUser",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return JobsUserActivityMultiError(errors)
	}

	return nil
}

// JobsUserActivityMultiError is an error wrapping multiple validation errors
// returned by JobsUserActivity.ValidateAll() if the designated constraints
// aren't met.
type JobsUserActivityMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m JobsUserActivityMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m JobsUserActivityMultiError) AllErrors() []error { return m }

// JobsUserActivityValidationError is the validation error returned by
// JobsUserActivity.Validate if the designated constraints aren't met.
type JobsUserActivityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JobsUserActivityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JobsUserActivityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JobsUserActivityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JobsUserActivityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JobsUserActivityValidationError) ErrorName() string { return "JobsUserActivityValidationError" }

// Error satisfies the builtin error interface
func (e JobsUserActivityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJobsUserActivity.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JobsUserActivityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JobsUserActivityValidationError{}

// Validate checks the field values on JobsUserActivityData with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *JobsUserActivityData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on JobsUserActivityData with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// JobsUserActivityDataMultiError, or nil if none found.
func (m *JobsUserActivityData) ValidateAll() error {
	return m.validate(true)
}

func (m *JobsUserActivityData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	oneofDataPresent := false
	switch v := m.Data.(type) {
	case *JobsUserActivityData_AbsenceDate:
		if v == nil {
			err := JobsUserActivityDataValidationError{
				field:  "Data",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofDataPresent = true

		if all {
			switch v := interface{}(m.GetAbsenceDate()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, JobsUserActivityDataValidationError{
						field:  "AbsenceDate",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, JobsUserActivityDataValidationError{
						field:  "AbsenceDate",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetAbsenceDate()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return JobsUserActivityDataValidationError{
					field:  "AbsenceDate",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *JobsUserActivityData_GradeChange:
		if v == nil {
			err := JobsUserActivityDataValidationError{
				field:  "Data",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofDataPresent = true

		if all {
			switch v := interface{}(m.GetGradeChange()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, JobsUserActivityDataValidationError{
						field:  "GradeChange",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, JobsUserActivityDataValidationError{
						field:  "GradeChange",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetGradeChange()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return JobsUserActivityDataValidationError{
					field:  "GradeChange",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *JobsUserActivityData_LabelsChange:
		if v == nil {
			err := JobsUserActivityDataValidationError{
				field:  "Data",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofDataPresent = true

		if all {
			switch v := interface{}(m.GetLabelsChange()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, JobsUserActivityDataValidationError{
						field:  "LabelsChange",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, JobsUserActivityDataValidationError{
						field:  "LabelsChange",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetLabelsChange()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return JobsUserActivityDataValidationError{
					field:  "LabelsChange",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *JobsUserActivityData_NameChange:
		if v == nil {
			err := JobsUserActivityDataValidationError{
				field:  "Data",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofDataPresent = true

		if all {
			switch v := interface{}(m.GetNameChange()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, JobsUserActivityDataValidationError{
						field:  "NameChange",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, JobsUserActivityDataValidationError{
						field:  "NameChange",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetNameChange()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return JobsUserActivityDataValidationError{
					field:  "NameChange",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofDataPresent {
		err := JobsUserActivityDataValidationError{
			field:  "Data",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return JobsUserActivityDataMultiError(errors)
	}

	return nil
}

// JobsUserActivityDataMultiError is an error wrapping multiple validation
// errors returned by JobsUserActivityData.ValidateAll() if the designated
// constraints aren't met.
type JobsUserActivityDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m JobsUserActivityDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m JobsUserActivityDataMultiError) AllErrors() []error { return m }

// JobsUserActivityDataValidationError is the validation error returned by
// JobsUserActivityData.Validate if the designated constraints aren't met.
type JobsUserActivityDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JobsUserActivityDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JobsUserActivityDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JobsUserActivityDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JobsUserActivityDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JobsUserActivityDataValidationError) ErrorName() string {
	return "JobsUserActivityDataValidationError"
}

// Error satisfies the builtin error interface
func (e JobsUserActivityDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJobsUserActivityData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JobsUserActivityDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JobsUserActivityDataValidationError{}

// Validate checks the field values on ColleagueAbsenceDate with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ColleagueAbsenceDate) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ColleagueAbsenceDate with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ColleagueAbsenceDateMultiError, or nil if none found.
func (m *ColleagueAbsenceDate) ValidateAll() error {
	return m.validate(true)
}

func (m *ColleagueAbsenceDate) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetAbsenceBegin()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ColleagueAbsenceDateValidationError{
					field:  "AbsenceBegin",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ColleagueAbsenceDateValidationError{
					field:  "AbsenceBegin",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAbsenceBegin()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ColleagueAbsenceDateValidationError{
				field:  "AbsenceBegin",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetAbsenceEnd()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ColleagueAbsenceDateValidationError{
					field:  "AbsenceEnd",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ColleagueAbsenceDateValidationError{
					field:  "AbsenceEnd",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAbsenceEnd()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ColleagueAbsenceDateValidationError{
				field:  "AbsenceEnd",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ColleagueAbsenceDateMultiError(errors)
	}

	return nil
}

// ColleagueAbsenceDateMultiError is an error wrapping multiple validation
// errors returned by ColleagueAbsenceDate.ValidateAll() if the designated
// constraints aren't met.
type ColleagueAbsenceDateMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ColleagueAbsenceDateMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ColleagueAbsenceDateMultiError) AllErrors() []error { return m }

// ColleagueAbsenceDateValidationError is the validation error returned by
// ColleagueAbsenceDate.Validate if the designated constraints aren't met.
type ColleagueAbsenceDateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ColleagueAbsenceDateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ColleagueAbsenceDateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ColleagueAbsenceDateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ColleagueAbsenceDateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ColleagueAbsenceDateValidationError) ErrorName() string {
	return "ColleagueAbsenceDateValidationError"
}

// Error satisfies the builtin error interface
func (e ColleagueAbsenceDateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sColleagueAbsenceDate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ColleagueAbsenceDateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ColleagueAbsenceDateValidationError{}

// Validate checks the field values on ColleagueGradeChange with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ColleagueGradeChange) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ColleagueGradeChange with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ColleagueGradeChangeMultiError, or nil if none found.
func (m *ColleagueGradeChange) ValidateAll() error {
	return m.validate(true)
}

func (m *ColleagueGradeChange) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Grade

	// no validation rules for GradeLabel

	if len(errors) > 0 {
		return ColleagueGradeChangeMultiError(errors)
	}

	return nil
}

// ColleagueGradeChangeMultiError is an error wrapping multiple validation
// errors returned by ColleagueGradeChange.ValidateAll() if the designated
// constraints aren't met.
type ColleagueGradeChangeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ColleagueGradeChangeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ColleagueGradeChangeMultiError) AllErrors() []error { return m }

// ColleagueGradeChangeValidationError is the validation error returned by
// ColleagueGradeChange.Validate if the designated constraints aren't met.
type ColleagueGradeChangeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ColleagueGradeChangeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ColleagueGradeChangeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ColleagueGradeChangeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ColleagueGradeChangeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ColleagueGradeChangeValidationError) ErrorName() string {
	return "ColleagueGradeChangeValidationError"
}

// Error satisfies the builtin error interface
func (e ColleagueGradeChangeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sColleagueGradeChange.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ColleagueGradeChangeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ColleagueGradeChangeValidationError{}

// Validate checks the field values on ColleagueLabelsChange with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ColleagueLabelsChange) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ColleagueLabelsChange with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ColleagueLabelsChangeMultiError, or nil if none found.
func (m *ColleagueLabelsChange) ValidateAll() error {
	return m.validate(true)
}

func (m *ColleagueLabelsChange) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetAdded() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ColleagueLabelsChangeValidationError{
						field:  fmt.Sprintf("Added[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ColleagueLabelsChangeValidationError{
						field:  fmt.Sprintf("Added[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ColleagueLabelsChangeValidationError{
					field:  fmt.Sprintf("Added[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetRemoved() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ColleagueLabelsChangeValidationError{
						field:  fmt.Sprintf("Removed[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ColleagueLabelsChangeValidationError{
						field:  fmt.Sprintf("Removed[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ColleagueLabelsChangeValidationError{
					field:  fmt.Sprintf("Removed[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ColleagueLabelsChangeMultiError(errors)
	}

	return nil
}

// ColleagueLabelsChangeMultiError is an error wrapping multiple validation
// errors returned by ColleagueLabelsChange.ValidateAll() if the designated
// constraints aren't met.
type ColleagueLabelsChangeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ColleagueLabelsChangeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ColleagueLabelsChangeMultiError) AllErrors() []error { return m }

// ColleagueLabelsChangeValidationError is the validation error returned by
// ColleagueLabelsChange.Validate if the designated constraints aren't met.
type ColleagueLabelsChangeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ColleagueLabelsChangeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ColleagueLabelsChangeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ColleagueLabelsChangeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ColleagueLabelsChangeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ColleagueLabelsChangeValidationError) ErrorName() string {
	return "ColleagueLabelsChangeValidationError"
}

// Error satisfies the builtin error interface
func (e ColleagueLabelsChangeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sColleagueLabelsChange.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ColleagueLabelsChangeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ColleagueLabelsChangeValidationError{}

// Validate checks the field values on ColleagueNameChange with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ColleagueNameChange) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ColleagueNameChange with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ColleagueNameChangeMultiError, or nil if none found.
func (m *ColleagueNameChange) ValidateAll() error {
	return m.validate(true)
}

func (m *ColleagueNameChange) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.Prefix != nil {
		// no validation rules for Prefix
	}

	if m.Suffix != nil {
		// no validation rules for Suffix
	}

	if len(errors) > 0 {
		return ColleagueNameChangeMultiError(errors)
	}

	return nil
}

// ColleagueNameChangeMultiError is an error wrapping multiple validation
// errors returned by ColleagueNameChange.ValidateAll() if the designated
// constraints aren't met.
type ColleagueNameChangeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ColleagueNameChangeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ColleagueNameChangeMultiError) AllErrors() []error { return m }

// ColleagueNameChangeValidationError is the validation error returned by
// ColleagueNameChange.Validate if the designated constraints aren't met.
type ColleagueNameChangeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ColleagueNameChangeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ColleagueNameChangeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ColleagueNameChangeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ColleagueNameChangeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ColleagueNameChangeValidationError) ErrorName() string {
	return "ColleagueNameChangeValidationError"
}

// Error satisfies the builtin error interface
func (e ColleagueNameChangeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sColleagueNameChange.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ColleagueNameChangeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ColleagueNameChangeValidationError{}
