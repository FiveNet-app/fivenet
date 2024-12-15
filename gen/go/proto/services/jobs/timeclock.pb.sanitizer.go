// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: services/jobs/timeclock.proto

package jobs

func (m *GetTimeclockStatsRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *GetTimeclockStatsResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Stats
	if m.Stats != nil {
		if v, ok := interface{}(m.GetStats()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	// Field: Weekly
	for idx, item := range m.Weekly {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ListInactiveEmployeesRequest) Sanitize() error {
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

	// Field: Sort
	if m.Sort != nil {
		if v, ok := interface{}(m.GetSort()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *ListInactiveEmployeesResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Colleagues
	for idx, item := range m.Colleagues {
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

func (m *ListTimeclockRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Date
	if m.Date != nil {
		if v, ok := interface{}(m.GetDate()).(interface{ Sanitize() error }); ok {
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

	// Field: Sort
	if m.Sort != nil {
		if v, ok := interface{}(m.GetSort()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *ListTimeclockResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Daily
	if m == nil {
		return nil
	}

	switch v := m.Entries.(type) {

	case *ListTimeclockResponse_Daily:
		if v, ok := interface{}(v).(interface{ Sanitize() error }); ok {
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

	// Field: Range
	if m == nil {
		return nil
	}

	switch v := m.Entries.(type) {

	case *ListTimeclockResponse_Range:
		if v, ok := interface{}(v).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

	}

	// Field: Stats
	if m.Stats != nil {
		if v, ok := interface{}(m.GetStats()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	// Field: StatsWeekly
	for idx, item := range m.StatsWeekly {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

	}

	// Field: Weekly
	if m == nil {
		return nil
	}

	switch v := m.Entries.(type) {

	case *ListTimeclockResponse_Weekly:
		if v, ok := interface{}(v).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *TimeclockDay) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Date
	if m.Date != nil {
		if v, ok := interface{}(m.GetDate()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
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

func (m *TimeclockRange) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Date
	if m.Date != nil {
		if v, ok := interface{}(m.GetDate()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
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

func (m *TimeclockWeekly) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Date
	if m.Date != nil {
		if v, ok := interface{}(m.GetDate()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
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
