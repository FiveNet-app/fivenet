// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: resources/internet/domain.proto

package internet

func (m *Domain) Sanitize() error {
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

	// Field: DeletedAt
	if m.DeletedAt != nil {
		if v, ok := interface{}(m.GetDeletedAt()).(interface{ Sanitize() error }); ok {
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