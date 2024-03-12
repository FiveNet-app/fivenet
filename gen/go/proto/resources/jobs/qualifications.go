package jobs

import (
	"database/sql/driver"

	"github.com/galexrt/fivenet/pkg/utils/protoutils"
	"google.golang.org/protobuf/encoding/protojson"
)

func (x *QualificationJobAccess) SetJob(job string) {
	x.Job = job
}

func (x *QualificationJobAccess) SetJobLabel(label string) {
	x.JobLabel = &label
}

func (x *QualificationJobAccess) GetJobGrade() int32 {
	return x.MinimumGrade
}

func (x *QualificationJobAccess) SetJobGrade(grade int32) {
	x.MinimumGrade = grade
}

func (x *QualificationJobAccess) SetJobGradeLabel(label string) {
	x.JobGradeLabel = &label
}

func (x *QualificationDiscordSettings) Scan(value any) error {
	switch t := value.(type) {
	case string:
		return protojson.Unmarshal([]byte(t), x)
	case []byte:
		return protojson.Unmarshal(t, x)
	}
	return nil
}

// Scan implements driver.Valuer for protobuf QualificationDiscordSettings.
func (x *QualificationDiscordSettings) Value() (driver.Value, error) {
	if x == nil {
		return nil, nil
	}

	out, err := protoutils.Marshal(x)
	return string(out), err
}
