// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: resources/jobs/colleagues.proto

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

// Validate checks the field values on Colleague with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Colleague) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Colleague with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ColleagueMultiError, or nil
// if none found.
func (m *Colleague) ValidateAll() error {
	return m.validate(true)
}

func (m *Colleague) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUserId() <= 0 {
		err := ColleagueValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetJob()) > 20 {
		err := ColleagueValidationError{
			field:  "Job",
			reason: "value length must be at most 20 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetJobGrade() <= -1 {
		err := ColleagueValidationError{
			field:  "JobGrade",
			reason: "value must be greater than -1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetFirstname()); l < 1 || l > 50 {
		err := ColleagueValidationError{
			field:  "Firstname",
			reason: "value length must be between 1 and 50 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetLastname()); l < 1 || l > 50 {
		err := ColleagueValidationError{
			field:  "Lastname",
			reason: "value length must be between 1 and 50 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetDateofbirth()) != 10 {
		err := ColleagueValidationError{
			field:  "Dateofbirth",
			reason: "value length must be 10 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if all {
		switch v := interface{}(m.GetProps()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ColleagueValidationError{
					field:  "Props",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ColleagueValidationError{
					field:  "Props",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProps()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ColleagueValidationError{
				field:  "Props",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.Identifier != nil {

		if utf8.RuneCountInString(m.GetIdentifier()) > 64 {
			err := ColleagueValidationError{
				field:  "Identifier",
				reason: "value length must be at most 64 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.JobLabel != nil {

		if utf8.RuneCountInString(m.GetJobLabel()) > 50 {
			err := ColleagueValidationError{
				field:  "JobLabel",
				reason: "value length must be at most 50 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.JobGradeLabel != nil {

		if utf8.RuneCountInString(m.GetJobGradeLabel()) > 50 {
			err := ColleagueValidationError{
				field:  "JobGradeLabel",
				reason: "value length must be at most 50 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.PhoneNumber != nil {

		if utf8.RuneCountInString(m.GetPhoneNumber()) > 20 {
			err := ColleagueValidationError{
				field:  "PhoneNumber",
				reason: "value length must be at most 20 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.Avatar != nil {

		if all {
			switch v := interface{}(m.GetAvatar()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ColleagueValidationError{
						field:  "Avatar",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ColleagueValidationError{
						field:  "Avatar",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetAvatar()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ColleagueValidationError{
					field:  "Avatar",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Email != nil {

		if l := utf8.RuneCountInString(m.GetEmail()); l < 6 || l > 80 {
			err := ColleagueValidationError{
				field:  "Email",
				reason: "value length must be between 6 and 80 runes, inclusive",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return ColleagueMultiError(errors)
	}

	return nil
}

// ColleagueMultiError is an error wrapping multiple validation errors returned
// by Colleague.ValidateAll() if the designated constraints aren't met.
type ColleagueMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ColleagueMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ColleagueMultiError) AllErrors() []error { return m }

// ColleagueValidationError is the validation error returned by
// Colleague.Validate if the designated constraints aren't met.
type ColleagueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ColleagueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ColleagueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ColleagueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ColleagueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ColleagueValidationError) ErrorName() string { return "ColleagueValidationError" }

// Error satisfies the builtin error interface
func (e ColleagueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sColleague.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ColleagueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ColleagueValidationError{}

// Validate checks the field values on JobsUserProps with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *JobsUserProps) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on JobsUserProps with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in JobsUserPropsMultiError, or
// nil if none found.
func (m *JobsUserProps) ValidateAll() error {
	return m.validate(true)
}

func (m *JobsUserProps) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUserId() <= 0 {
		err := JobsUserPropsValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetJob()) > 20 {
		err := JobsUserPropsValidationError{
			field:  "Job",
			reason: "value length must be at most 20 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.AbsenceBegin != nil {

		if all {
			switch v := interface{}(m.GetAbsenceBegin()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, JobsUserPropsValidationError{
						field:  "AbsenceBegin",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, JobsUserPropsValidationError{
						field:  "AbsenceBegin",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetAbsenceBegin()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return JobsUserPropsValidationError{
					field:  "AbsenceBegin",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.AbsenceEnd != nil {

		if all {
			switch v := interface{}(m.GetAbsenceEnd()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, JobsUserPropsValidationError{
						field:  "AbsenceEnd",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, JobsUserPropsValidationError{
						field:  "AbsenceEnd",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetAbsenceEnd()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return JobsUserPropsValidationError{
					field:  "AbsenceEnd",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Note != nil {
		// no validation rules for Note
	}

	if m.Labels != nil {

		if all {
			switch v := interface{}(m.GetLabels()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, JobsUserPropsValidationError{
						field:  "Labels",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, JobsUserPropsValidationError{
						field:  "Labels",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetLabels()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return JobsUserPropsValidationError{
					field:  "Labels",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.NamePrefix != nil {

		if utf8.RuneCountInString(m.GetNamePrefix()) > 12 {
			err := JobsUserPropsValidationError{
				field:  "NamePrefix",
				reason: "value length must be at most 12 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.NameSuffix != nil {

		if utf8.RuneCountInString(m.GetNameSuffix()) > 12 {
			err := JobsUserPropsValidationError{
				field:  "NameSuffix",
				reason: "value length must be at most 12 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return JobsUserPropsMultiError(errors)
	}

	return nil
}

// JobsUserPropsMultiError is an error wrapping multiple validation errors
// returned by JobsUserProps.ValidateAll() if the designated constraints
// aren't met.
type JobsUserPropsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m JobsUserPropsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m JobsUserPropsMultiError) AllErrors() []error { return m }

// JobsUserPropsValidationError is the validation error returned by
// JobsUserProps.Validate if the designated constraints aren't met.
type JobsUserPropsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JobsUserPropsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JobsUserPropsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JobsUserPropsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JobsUserPropsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JobsUserPropsValidationError) ErrorName() string { return "JobsUserPropsValidationError" }

// Error satisfies the builtin error interface
func (e JobsUserPropsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJobsUserProps.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JobsUserPropsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JobsUserPropsValidationError{}
