// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: services/rector/config.proto

package rector

func (m *GetAppConfigRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *GetAppConfigResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Config
	if m.Config != nil {
		if v, ok := interface{}(m.GetConfig()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *UpdateAppConfigRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Config
	if m.Config != nil {
		if v, ok := interface{}(m.GetConfig()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *UpdateAppConfigResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Config
	if m.Config != nil {
		if v, ok := interface{}(m.GetConfig()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}
