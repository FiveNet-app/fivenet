// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: resources/documents/requests.proto

package documents

func (m *DocRequest) Sanitize() error {
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

	// Field: Creator
	if m.Creator != nil {
		if v, ok := interface{}(m.GetCreator()).(interface{ Sanitize() error }); ok {
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