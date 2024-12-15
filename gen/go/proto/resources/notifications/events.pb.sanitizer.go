// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: resources/notifications/events.proto

package notifications

func (m *JobEvent) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: JobProps
	switch v := m.Data.(type) {

	case *JobEvent_JobProps:
		if v, ok := interface{}(v).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *JobGradeEvent) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *SystemEvent) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *UserEvent) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Notification
	switch v := m.Data.(type) {

	case *UserEvent_Notification:
		if v, ok := interface{}(v).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

	}

	return nil
}
