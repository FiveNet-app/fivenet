// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: services/calendar/calendar.proto

package permscalendar

import (
	"github.com/fivenet-app/fivenet/pkg/perms"
)

const (
	CalendarServicePerm perms.Category = "CalendarService"

	CalendarServiceCreateOrUpdateCalendarPerm            perms.Name = "CreateOrUpdateCalendar"
	CalendarServiceCreateOrUpdateCalendarFieldsPermField perms.Key  = "Fields"
	CalendarServiceCreateOrUpdateCalendarEntryPerm       perms.Name = "CreateOrUpdateCalendarEntry"
	CalendarServiceDeleteCalendarPerm                    perms.Name = "DeleteCalendar"
	CalendarServiceDeleteCalendarEntriesPerm             perms.Name = "DeleteCalendarEntries"
)
