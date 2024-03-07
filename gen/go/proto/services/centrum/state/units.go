package state

import (
	"context"
	"slices"
	"strings"

	"github.com/galexrt/fivenet/gen/go/proto/resources/centrum"
)

func (s *State) GetUnit(ctx context.Context, job string, id uint64) (*centrum.Unit, error) {
	return s.units.GetOrLoad(ctx, JobIdKey(job, id))
}

func (s *State) ListUnits(ctx context.Context, job string) ([]*centrum.Unit, bool) {
	us := []*centrum.Unit{}

	ids, err := s.units.Keys(ctx, job)
	if err != nil {
		return us, false
	}

	for _, id := range ids {
		unit, err := s.units.GetOrLoad(ctx, id)
		if err != nil {
			continue
		}

		us = append(us, unit)
	}

	slices.SortFunc(us, func(a, b *centrum.Unit) int {
		return strings.Compare(a.Name, b.Name)
	})

	return us, true
}

func (s *State) FilterUnits(ctx context.Context, job string, statuses []centrum.StatusUnit, notStatuses []centrum.StatusUnit, filterFn func(unit *centrum.Unit) bool) []*centrum.Unit {
	units, ok := s.ListUnits(ctx, job)
	if !ok {
		return nil
	}

	us := []*centrum.Unit{}
	for i := 0; i < len(units); i++ {
		include := true

		// Include statuses that should be listed
		if len(statuses) > 0 && !slices.Contains(statuses, units[i].Status.Status) {
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

		if include && (filterFn == nil || filterFn(units[i])) {
			us = append(us, units[i])
		}
	}

	return us
}

func (s *State) DeleteUnit(ctx context.Context, job string, id uint64) error {
	return s.units.Delete(ctx, JobIdKey(job, id))
}

func (s *State) UpdateUnit(ctx context.Context, job string, id uint64, unit *centrum.Unit) error {
	if err := s.units.ComputeUpdate(ctx, JobIdKey(job, id), true, func(key string, existing *centrum.Unit) (*centrum.Unit, error) {
		if existing == nil {
			return unit, nil
		}

		existing.Merge(unit)

		return existing, nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *State) UpdateUnitStatus(ctx context.Context, job string, id uint64, status *centrum.UnitStatus) error {
	if err := s.units.ComputeUpdate(ctx, JobIdKey(job, id), true, func(key string, existing *centrum.Unit) (*centrum.Unit, error) {
		if existing == nil {
			return nil, nil
		}

		existing.Status = status

		return existing, nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *State) UpdateUnitUsers(ctx context.Context, job string, id uint64, users []*centrum.UnitAssignment) error {
	if err := s.units.ComputeUpdate(ctx, JobIdKey(job, id), true, func(key string, existing *centrum.Unit) (*centrum.Unit, error) {
		if existing == nil {
			return nil, nil
		}

		if users == nil {
			existing.Users = []*centrum.UnitAssignment{}
		} else {
			existing.Users = users
		}

		return existing, nil
	}); err != nil {
		return err
	}

	return nil
}
