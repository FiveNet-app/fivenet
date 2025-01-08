package sync

import (
	"context"

	"github.com/fivenet-app/fivenet/gen/go/proto/resources/common/database"
	"github.com/fivenet-app/fivenet/gen/go/proto/resources/sync"
	"github.com/fivenet-app/fivenet/pkg/utils/dbutils/tables"
	jet "github.com/go-jet/jet/v2/mysql"
)

var (
	tJobs         = tables.Jobs
	tUsers        = tables.Users
	tVehicles     = tables.OwnedVehicles
	tLicenses     = tables.Licenses
	tUserLicenses = tables.UserLicenses
)

func (s *Server) GetStatus(ctx context.Context, req *GetStatusRequest) (*GetStatusResponse, error) {
	resp := &GetStatusResponse{}

	// Jobs
	jobsStmt := tJobs.
		SELECT(
			jet.COUNT(tJobs.Name),
		).
		FROM(tJobs)

	var jobsCount database.DataCount
	if err := jobsStmt.QueryContext(ctx, s.db, &jobsCount); err != nil {
		return nil, err
	}
	resp.Jobs = &sync.DataStatus{
		Count: jobsCount.TotalCount,
	}

	// Users
	usersStmt := tUsers.
		SELECT(
			jet.COUNT(tUsers.ID),
		).
		FROM(tUsers)

	var usersCount database.DataCount
	if err := usersStmt.QueryContext(ctx, s.db, &usersCount); err != nil {
		return nil, err
	}
	resp.Users = &sync.DataStatus{
		Count: usersCount.TotalCount,
	}

	// Vehicles
	vehiclesStmt := tVehicles.
		SELECT(
			jet.COUNT(tVehicles.Plate),
		).
		FROM(tVehicles)

	var vehiclesCount database.DataCount
	if err := vehiclesStmt.QueryContext(ctx, s.db, &vehiclesCount); err != nil {
		return nil, err
	}
	resp.Vehicles = &sync.DataStatus{
		Count: vehiclesCount.TotalCount,
	}

	// Licenses
	licensesStmt := tLicenses.
		SELECT(
			jet.COUNT(tLicenses.Type),
		).
		FROM(tLicenses)

	var licensesCount database.DataCount
	if err := licensesStmt.QueryContext(ctx, s.db, &licensesCount); err != nil {
		return nil, err
	}
	resp.Licenses = &sync.DataStatus{
		Count: licensesCount.TotalCount,
	}

	return resp, nil
}
