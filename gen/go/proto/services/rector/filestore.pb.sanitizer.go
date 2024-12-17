// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: services/rector/filestore.proto

package rector

func (m *DeleteFileRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *DeleteFileResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *ListFilesRequest) Sanitize() error {
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

func (m *ListFilesResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Files
	for idx, item := range m.Files {
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

func (m *UploadFileRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: File
	if m.File != nil {
		if v, ok := interface{}(m.GetFile()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *UploadFileResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: File
	if m.File != nil {
		if v, ok := interface{}(m.GetFile()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}