// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: resources/notifications/notifications.proto

package notifications

func (m *CalendarData) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *Data) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Calendar
	if m.Calendar != nil {
		if v, ok := interface{}(m.GetCalendar()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	// Field: CausedBy
	if m.CausedBy != nil {
		if v, ok := interface{}(m.GetCausedBy()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	// Field: Link
	if m.Link != nil {
		if v, ok := interface{}(m.GetLink()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *Link) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *Notification) Sanitize() error {
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

	// Field: ReadAt
	if m.ReadAt != nil {
		if v, ok := interface{}(m.GetReadAt()).(interface{ Sanitize() error }); ok {
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