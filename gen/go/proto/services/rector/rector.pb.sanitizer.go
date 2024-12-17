// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: services/rector/rector.proto

package rector

import (
	"github.com/fivenet-app/fivenet/pkg/html/htmlsanitizer"
)

func (m *AttrsUpdate) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: ToRemove
	for idx, item := range m.ToRemove {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

	}

	// Field: ToUpdate
	for idx, item := range m.ToUpdate {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *CreateRoleRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *CreateRoleResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Role
	if m.Role != nil {
		if v, ok := interface{}(m.GetRole()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *DeleteFactionRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *DeleteFactionResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *DeleteRoleRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *DeleteRoleResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *GetJobPropsRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *GetJobPropsResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: JobProps
	if m.JobProps != nil {
		if v, ok := interface{}(m.GetJobProps()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *GetPermissionsRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *GetPermissionsResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Attributes
	for idx, item := range m.Attributes {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

	}

	// Field: Permissions
	for idx, item := range m.Permissions {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *GetRoleRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *GetRoleResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Role
	if m.Role != nil {
		if v, ok := interface{}(m.GetRole()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *GetRolesRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *GetRolesResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Roles
	for idx, item := range m.Roles {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *PermItem) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *PermsUpdate) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: ToUpdate
	for idx, item := range m.ToUpdate {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *SetJobPropsRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: JobProps
	if m.JobProps != nil {
		if v, ok := interface{}(m.GetJobProps()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *SetJobPropsResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: JobProps
	if m.JobProps != nil {
		if v, ok := interface{}(m.GetJobProps()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *UpdateRoleLimitsRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Attrs
	if m.Attrs != nil {
		if v, ok := interface{}(m.GetAttrs()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	// Field: Perms
	if m.Perms != nil {
		if v, ok := interface{}(m.GetPerms()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *UpdateRoleLimitsResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *UpdateRolePermsRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Attrs
	if m.Attrs != nil {
		if v, ok := interface{}(m.GetAttrs()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	// Field: Perms
	if m.Perms != nil {
		if v, ok := interface{}(m.GetPerms()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *UpdateRolePermsResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *ViewAuditLogRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: From
	if m.From != nil {
		if v, ok := interface{}(m.GetFrom()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	// Field: Methods
	for idx, item := range m.Methods {
		_, _ = idx, item

		m.Methods[idx] = htmlsanitizer.StripTags(m.Methods[idx])

	}

	// Field: Pagination
	if m.Pagination != nil {
		if v, ok := interface{}(m.GetPagination()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	// Field: Services
	for idx, item := range m.Services {
		_, _ = idx, item

		m.Services[idx] = htmlsanitizer.StripTags(m.Services[idx])

	}

	// Field: Sort
	if m.Sort != nil {
		if v, ok := interface{}(m.GetSort()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	// Field: To
	if m.To != nil {
		if v, ok := interface{}(m.GetTo()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *ViewAuditLogResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Logs
	for idx, item := range m.Logs {
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