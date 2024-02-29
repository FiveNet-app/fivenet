package common

import (
	"database/sql/driver"

	"google.golang.org/protobuf/encoding/protojson"
)

func (x *TranslateItem) Scan(value any) error {
	switch t := value.(type) {
	case string:
		return protojson.Unmarshal([]byte(t), x)
	case []byte:
		return protojson.Unmarshal(t, x)
	}
	return nil
}

// Scan implements driver.Valuer for protobuf TranslateItem.
func (x *TranslateItem) Value() (driver.Value, error) {
	if x == nil {
		return nil, nil
	}

	out, err := protojson.Marshal(x)
	return string(out), err
}
