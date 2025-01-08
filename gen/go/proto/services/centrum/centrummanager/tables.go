package centrummanager

import (
	"github.com/fivenet-app/fivenet/pkg/utils/dbutils/tables"
	"github.com/fivenet-app/fivenet/query/fivenet/table"
)

var (
	tUnits           = table.FivenetCentrumUnits.AS("unit")
	tUnitStatus      = table.FivenetCentrumUnitsStatus.AS("unitstatus")
	tUnitUser        = table.FivenetCentrumUnitsUsers.AS("unitassignment")
	tUsers           = tables.Users.AS("colleague")
	tUserProps       = table.FivenetUserProps
	tJobsUserProps   = table.FivenetJobsUserProps.AS("jobs_user_props")
	tCentrumSettings = table.FivenetCentrumSettings
	tCentrumUsers    = table.FivenetCentrumUsers
	tDispatch        = table.FivenetCentrumDispatches.AS("dispatch")
	tDispatchStatus  = table.FivenetCentrumDispatchesStatus.AS("dispatchstatus")
	tDispatchUnit    = table.FivenetCentrumDispatchesAsgmts.AS("dispatchassignment")
)
