// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: resources/common/grpcws/grpcws.proto

package grpcws

func (m *Body) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Data

	return nil
}

func (m *Cancel) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *Complete) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *Failure) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Headers
	for idx, item := range m.Headers {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *GrpcFrame) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Body
	if m == nil {
		return nil
	}

	switch v := m.Payload.(type) {

	case *GrpcFrame_Body:
		if v, ok := interface{}(v).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

		// Field: Cancel
	case *GrpcFrame_Cancel:
		if v, ok := interface{}(v).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

		// Field: Complete
	case *GrpcFrame_Complete:
		if v, ok := interface{}(v).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

		// Field: Failure
	case *GrpcFrame_Failure:
		if v, ok := interface{}(v).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

		// Field: Header
	case *GrpcFrame_Header:
		if v, ok := interface{}(v).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

		// Field: Ping
	case *GrpcFrame_Ping:
		if v, ok := interface{}(v).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Header) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Headers
	for idx, item := range m.Headers {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *HeaderValue) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *Ping) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}
