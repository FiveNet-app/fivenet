// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: resources/qualifications/exam.proto

package qualifications

import (
	"github.com/fivenet-app/fivenet/pkg/html/htmlsanitizer"
)

func (m *ExamQuestion) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Answer
	if m.Answer != nil {
		if v, ok := interface{}(m.GetAnswer()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	// Field: CreatedAt
	if m.CreatedAt != nil {
		if v, ok := interface{}(m.GetCreatedAt()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	// Field: Data
	if m.Data != nil {
		if v, ok := interface{}(m.GetData()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	// Field: Description

	if m.Description != nil {
		*m.Description = htmlsanitizer.StripTags(*m.Description)
	}

	// Field: Title
	m.Title = htmlsanitizer.StripTags(m.Title)

	// Field: UpdatedAt
	if m.UpdatedAt != nil {
		if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *ExamQuestionAnswerData) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *ExamQuestionData) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: FreeText
	if m == nil {
		return nil
	}

	switch v := m.Data.(type) {

	case *ExamQuestionData_FreeText:
		if v, ok := interface{}(v).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

		// Field: Image
	case *ExamQuestionData_Image:
		if v, ok := interface{}(v).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

		// Field: MultipleChoice
	case *ExamQuestionData_MultipleChoice:
		if v, ok := interface{}(v).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

		// Field: Separator
	case *ExamQuestionData_Separator:
		if v, ok := interface{}(v).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

		// Field: SingleChoice
	case *ExamQuestionData_SingleChoice:
		if v, ok := interface{}(v).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

		// Field: Yesno
	case *ExamQuestionData_Yesno:
		if v, ok := interface{}(v).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ExamQuestionImage) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Image
	if m.Image != nil {
		if v, ok := interface{}(m.GetImage()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *ExamQuestionMultipleChoice) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Choices
	for idx, item := range m.Choices {
		_, _ = idx, item

		m.Choices[idx] = htmlsanitizer.StripTags(m.Choices[idx])

	}

	return nil
}

func (m *ExamQuestionSeparator) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *ExamQuestionSingleChoice) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Choices
	for idx, item := range m.Choices {
		_, _ = idx, item

		m.Choices[idx] = htmlsanitizer.StripTags(m.Choices[idx])

	}

	return nil
}

func (m *ExamQuestionText) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *ExamQuestionYesNo) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *ExamQuestions) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Questions
	for idx, item := range m.Questions {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ExamResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Question
	if m.Question != nil {
		if v, ok := interface{}(m.GetQuestion()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	// Field: Response
	if m.Response != nil {
		if v, ok := interface{}(m.GetResponse()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *ExamResponseData) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: FreeText
	if m == nil {
		return nil
	}

	switch v := m.Response.(type) {

	case *ExamResponseData_FreeText:
		if v, ok := interface{}(v).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

		// Field: MultipleChoice
	case *ExamResponseData_MultipleChoice:
		if v, ok := interface{}(v).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

		// Field: Separator
	case *ExamResponseData_Separator:
		if v, ok := interface{}(v).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

		// Field: SingleChoice
	case *ExamResponseData_SingleChoice:
		if v, ok := interface{}(v).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

		// Field: Yesno
	case *ExamResponseData_Yesno:
		if v, ok := interface{}(v).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ExamResponseMultipleChoice) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Choices
	for idx, item := range m.Choices {
		_, _ = idx, item

		m.Choices[idx] = htmlsanitizer.StripTags(m.Choices[idx])

	}

	return nil
}

func (m *ExamResponseSeparator) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *ExamResponseSingleChoice) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Choice
	m.Choice = htmlsanitizer.StripTags(m.Choice)

	return nil
}

func (m *ExamResponseText) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Text
	m.Text = htmlsanitizer.StripTags(m.Text)

	return nil
}

func (m *ExamResponseYesNo) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *ExamResponses) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Responses
	for idx, item := range m.Responses {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ExamUser) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: CreatedAt
	if m.CreatedAt != nil {
		if v, ok := interface{}(m.GetCreatedAt()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	// Field: EndedAt
	if m.EndedAt != nil {
		if v, ok := interface{}(m.GetEndedAt()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	// Field: EndsAt
	if m.EndsAt != nil {
		if v, ok := interface{}(m.GetEndsAt()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	// Field: StartedAt
	if m.StartedAt != nil {
		if v, ok := interface{}(m.GetStartedAt()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}
