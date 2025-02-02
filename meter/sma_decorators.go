package meter

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decorateSMA(base *SMA, battery func() (float64, error), batteryCapacity func() float64) api.Meter {
	switch {
	case battery == nil && batteryCapacity == nil:
		return base

	case battery != nil && batteryCapacity == nil:
		return &struct {
			*SMA
			api.Battery
		}{
			SMA: base,
			Battery: &decorateSMABatteryImpl{
				battery: battery,
			},
		}

	case battery != nil && batteryCapacity != nil:
		return &struct {
			*SMA
			api.Battery
			api.BatteryCapacity
		}{
			SMA: base,
			Battery: &decorateSMABatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateSMABatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
		}
	}

	return nil
}

type decorateSMABatteryImpl struct {
	battery func() (float64, error)
}

func (impl *decorateSMABatteryImpl) Soc() (float64, error) {
	return impl.battery()
}

type decorateSMABatteryCapacityImpl struct {
	batteryCapacity func() float64
}

func (impl *decorateSMABatteryCapacityImpl) Capacity() float64 {
	return impl.batteryCapacity()
}
