// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: resources/common/error.proto

package common

func (m *Error) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Content
	if m.Content != nil {
		if v, ok := interface{}(m.GetContent()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	// Field: Title
	if m.Title != nil {
		if v, ok := interface{}(m.GetTitle()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}
