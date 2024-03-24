package documents

import (
	"database/sql/driver"

	"github.com/galexrt/fivenet/pkg/utils/protoutils"
	"google.golang.org/protobuf/encoding/protojson"
)

// Scan implements driver.Valuer for protobuf TemplateSchema.
func (x *TemplateSchema) Scan(value any) error {
	switch t := value.(type) {
	case string:
		return protojson.Unmarshal([]byte(t), x)
	case []byte:
		return protojson.Unmarshal(t, x)
	}
	return nil
}

// Value marshals the value into driver.Valuer.
func (x *TemplateSchema) Value() (driver.Value, error) {
	if x == nil {
		return nil, nil
	}

	out, err := protoutils.Marshal(x)
	return string(out), err
}

func (x *Template) GetJob() string {
	return x.CreatorJob
}

func (x *Template) SetJobLabel(label string) {
	x.CreatorJobLabel = &label
}

func (x *TemplateShort) GetJob() string {
	return x.CreatorJob
}

func (x *TemplateShort) SetJobLabel(label string) {
	x.CreatorJobLabel = &label
}
