// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: services/calendar/calendar.proto

package calendar

func (m *CreateOrUpdateCalendarEntryRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Entry
	if m.Entry != nil {
		if v, ok := interface{}(m.GetEntry()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *CreateOrUpdateCalendarEntryResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Entry
	if m.Entry != nil {
		if v, ok := interface{}(m.GetEntry()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *CreateOrUpdateCalendarRequest) Sanitize() error {
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

	return nil
}

func (m *CreateOrUpdateCalendarResponse) Sanitize() error {
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

	return nil
}

func (m *DeleteCalendarEntryRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *DeleteCalendarEntryResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *DeleteCalendarRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *DeleteCalendarResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *GetCalendarEntryRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *GetCalendarEntryResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Entry
	if m.Entry != nil {
		if v, ok := interface{}(m.GetEntry()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *GetCalendarRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *GetCalendarResponse) Sanitize() error {
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

	return nil
}

func (m *GetUpcomingEntriesRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *GetUpcomingEntriesResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Entries
	for idx, item := range m.Entries {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ListCalendarEntriesRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: After
	if m.After != nil {
		if v, ok := interface{}(m.GetAfter()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *ListCalendarEntriesResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Entries
	for idx, item := range m.Entries {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ListCalendarEntryRSVPRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Pagination
	if m.Pagination != nil {
		if v, ok := interface{}(m.GetPagination()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *ListCalendarEntryRSVPResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Entries
	for idx, item := range m.Entries {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

	}

	// Field: Pagination
	if m.Pagination != nil {
		if v, ok := interface{}(m.GetPagination()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *ListCalendarsRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: After
	if m.After != nil {
		if v, ok := interface{}(m.GetAfter()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	// Field: Pagination
	if m.Pagination != nil {
		if v, ok := interface{}(m.GetPagination()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *ListCalendarsResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Calendars
	for idx, item := range m.Calendars {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

	}

	// Field: Pagination
	if m.Pagination != nil {
		if v, ok := interface{}(m.GetPagination()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *ListSubscriptionsRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Pagination
	if m.Pagination != nil {
		if v, ok := interface{}(m.GetPagination()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *ListSubscriptionsResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Pagination
	if m.Pagination != nil {
		if v, ok := interface{}(m.GetPagination()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	// Field: Subs
	for idx, item := range m.Subs {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *RSVPCalendarEntryRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Entry
	if m.Entry != nil {
		if v, ok := interface{}(m.GetEntry()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *RSVPCalendarEntryResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Entry
	if m.Entry != nil {
		if v, ok := interface{}(m.GetEntry()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *ShareCalendarEntryRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *ShareCalendarEntryResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *SubscribeToCalendarRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Sub
	if m.Sub != nil {
		if v, ok := interface{}(m.GetSub()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *SubscribeToCalendarResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Sub
	if m.Sub != nil {
		if v, ok := interface{}(m.GetSub()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}
