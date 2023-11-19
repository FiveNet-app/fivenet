package state

import (
	"context"
	"strings"

	"github.com/galexrt/fivenet/gen/go/proto/resources/centrum"
	"github.com/galexrt/fivenet/pkg/utils"
	"golang.org/x/exp/slices"
)

func (s *State) GetUnit(job string, id uint64) *centrum.Unit {
	u, _ := s.units.Get(JobIdKey(job, id))
	return u
}

func (s *State) ListUnits(job string) ([]*centrum.Unit, bool) {
	us := []*centrum.Unit{}

	ids, err := s.units.Keys(job)
	if err != nil {
		return us, false
	}

	for _, id := range ids {
		unit, ok := s.units.Get(id)
		if !ok || unit == nil {
			continue
		}

		us = append(us, unit)
	}

	slices.SortFunc(us, func(a, b *centrum.Unit) int {
		return strings.Compare(a.Name, b.Name)
	})

	return us, true
}

func (s *State) FilterUnits(job string, statuses []centrum.StatusUnit, notStatuses []centrum.StatusUnit) []*centrum.Unit {
	units, ok := s.ListUnits(job)
	if !ok {
		return nil
	}

	us := []*centrum.Unit{}
	for i := 0; i < len(units); i++ {
		include := true

		// Include statuses that should be listed
		if len(statuses) > 0 && !utils.InSlice(statuses, units[i].Status.Status) {
			include = false
		} else if len(notStatuses) > 0 {
			// Which statuses to ignore
			for _, status := range notStatuses {
				if units[i].Status != nil && units[i].Status.Status == status {
					include = false
					break
				}
			}
		}

		if include {
			us = append(us, units[i])
		}
	}

	return us
}

func (s *State) DeleteUnit(job string, id uint64) error {
	return s.units.Delete(JobIdKey(job, id))
}

func (s *State) UpdateUnit(ctx context.Context, job string, id uint64, unit *centrum.Unit) error {
	s.units.ComputeUpdate(JobIdKey(job, id), true, func(key string, existing *centrum.Unit) (*centrum.Unit, error) {
		if existing == nil {
			return unit, nil
		}

		existing.Merge(unit)

		return existing, nil
	})

	return nil
}

func (s *State) UpdateUnitStatus(ctx context.Context, job string, id uint64, status *centrum.UnitStatus) error {
	s.units.ComputeUpdate(JobIdKey(job, id), true, func(key string, existing *centrum.Unit) (*centrum.Unit, error) {
		if existing == nil {
			return nil, nil
		}

		existing.Status = status

		return existing, nil
	})

	return nil
}
