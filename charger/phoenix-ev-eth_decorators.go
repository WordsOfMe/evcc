package charger

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decoratePhoenixEVEth(base *PhoenixEVEth, meter func() (float64, error), meterEnergy func() (float64, error), phaseCurrents func() (float64, float64, float64, error), phaseVoltages func() (float64, float64, float64, error), chargerEx func(float64) error, identifier func() (string, error)) api.Charger {
	switch {
	case chargerEx == nil && identifier == nil && meter == nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages == nil:
		return base

	case chargerEx == nil && identifier == nil && meter != nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*PhoenixEVEth
			api.Meter
		}{
			PhoenixEVEth: base,
			Meter: &decoratePhoenixEVEthMeterImpl{
				meter: meter,
			},
		}

	case chargerEx == nil && identifier == nil && meter != nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*PhoenixEVEth
			api.Meter
			api.MeterEnergy
		}{
			PhoenixEVEth: base,
			Meter: &decoratePhoenixEVEthMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decoratePhoenixEVEthMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargerEx == nil && identifier == nil && meter != nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*PhoenixEVEth
			api.Meter
			api.PhaseCurrents
		}{
			PhoenixEVEth: base,
			Meter: &decoratePhoenixEVEthMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decoratePhoenixEVEthPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargerEx == nil && identifier == nil && meter != nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*PhoenixEVEth
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
		}{
			PhoenixEVEth: base,
			Meter: &decoratePhoenixEVEthMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decoratePhoenixEVEthMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decoratePhoenixEVEthPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargerEx == nil && identifier == nil && meter != nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*PhoenixEVEth
			api.Meter
			api.PhaseVoltages
		}{
			PhoenixEVEth: base,
			Meter: &decoratePhoenixEVEthMeterImpl{
				meter: meter,
			},
			PhaseVoltages: &decoratePhoenixEVEthPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx == nil && identifier == nil && meter != nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*PhoenixEVEth
			api.Meter
			api.MeterEnergy
			api.PhaseVoltages
		}{
			PhoenixEVEth: base,
			Meter: &decoratePhoenixEVEthMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decoratePhoenixEVEthMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decoratePhoenixEVEthPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx == nil && identifier == nil && meter != nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*PhoenixEVEth
			api.Meter
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			PhoenixEVEth: base,
			Meter: &decoratePhoenixEVEthMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decoratePhoenixEVEthPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decoratePhoenixEVEthPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx == nil && identifier == nil && meter != nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*PhoenixEVEth
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			PhoenixEVEth: base,
			Meter: &decoratePhoenixEVEthMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decoratePhoenixEVEthMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decoratePhoenixEVEthPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decoratePhoenixEVEthPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx != nil && identifier == nil && meter == nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*PhoenixEVEth
			api.ChargerEx
		}{
			PhoenixEVEth: base,
			ChargerEx: &decoratePhoenixEVEthChargerExImpl{
				chargerEx: chargerEx,
			},
		}

	case chargerEx != nil && identifier == nil && meter != nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*PhoenixEVEth
			api.ChargerEx
			api.Meter
		}{
			PhoenixEVEth: base,
			ChargerEx: &decoratePhoenixEVEthChargerExImpl{
				chargerEx: chargerEx,
			},
			Meter: &decoratePhoenixEVEthMeterImpl{
				meter: meter,
			},
		}

	case chargerEx != nil && identifier == nil && meter != nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*PhoenixEVEth
			api.ChargerEx
			api.Meter
			api.MeterEnergy
		}{
			PhoenixEVEth: base,
			ChargerEx: &decoratePhoenixEVEthChargerExImpl{
				chargerEx: chargerEx,
			},
			Meter: &decoratePhoenixEVEthMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decoratePhoenixEVEthMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargerEx != nil && identifier == nil && meter != nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*PhoenixEVEth
			api.ChargerEx
			api.Meter
			api.PhaseCurrents
		}{
			PhoenixEVEth: base,
			ChargerEx: &decoratePhoenixEVEthChargerExImpl{
				chargerEx: chargerEx,
			},
			Meter: &decoratePhoenixEVEthMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decoratePhoenixEVEthPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargerEx != nil && identifier == nil && meter != nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*PhoenixEVEth
			api.ChargerEx
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
		}{
			PhoenixEVEth: base,
			ChargerEx: &decoratePhoenixEVEthChargerExImpl{
				chargerEx: chargerEx,
			},
			Meter: &decoratePhoenixEVEthMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decoratePhoenixEVEthMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decoratePhoenixEVEthPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargerEx != nil && identifier == nil && meter != nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*PhoenixEVEth
			api.ChargerEx
			api.Meter
			api.PhaseVoltages
		}{
			PhoenixEVEth: base,
			ChargerEx: &decoratePhoenixEVEthChargerExImpl{
				chargerEx: chargerEx,
			},
			Meter: &decoratePhoenixEVEthMeterImpl{
				meter: meter,
			},
			PhaseVoltages: &decoratePhoenixEVEthPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx != nil && identifier == nil && meter != nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*PhoenixEVEth
			api.ChargerEx
			api.Meter
			api.MeterEnergy
			api.PhaseVoltages
		}{
			PhoenixEVEth: base,
			ChargerEx: &decoratePhoenixEVEthChargerExImpl{
				chargerEx: chargerEx,
			},
			Meter: &decoratePhoenixEVEthMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decoratePhoenixEVEthMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decoratePhoenixEVEthPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx != nil && identifier == nil && meter != nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*PhoenixEVEth
			api.ChargerEx
			api.Meter
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			PhoenixEVEth: base,
			ChargerEx: &decoratePhoenixEVEthChargerExImpl{
				chargerEx: chargerEx,
			},
			Meter: &decoratePhoenixEVEthMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decoratePhoenixEVEthPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decoratePhoenixEVEthPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx != nil && identifier == nil && meter != nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*PhoenixEVEth
			api.ChargerEx
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			PhoenixEVEth: base,
			ChargerEx: &decoratePhoenixEVEthChargerExImpl{
				chargerEx: chargerEx,
			},
			Meter: &decoratePhoenixEVEthMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decoratePhoenixEVEthMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decoratePhoenixEVEthPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decoratePhoenixEVEthPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx == nil && identifier != nil && meter == nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*PhoenixEVEth
			api.Identifier
		}{
			PhoenixEVEth: base,
			Identifier: &decoratePhoenixEVEthIdentifierImpl{
				identifier: identifier,
			},
		}

	case chargerEx == nil && identifier != nil && meter != nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*PhoenixEVEth
			api.Identifier
			api.Meter
		}{
			PhoenixEVEth: base,
			Identifier: &decoratePhoenixEVEthIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decoratePhoenixEVEthMeterImpl{
				meter: meter,
			},
		}

	case chargerEx == nil && identifier != nil && meter != nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*PhoenixEVEth
			api.Identifier
			api.Meter
			api.MeterEnergy
		}{
			PhoenixEVEth: base,
			Identifier: &decoratePhoenixEVEthIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decoratePhoenixEVEthMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decoratePhoenixEVEthMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargerEx == nil && identifier != nil && meter != nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*PhoenixEVEth
			api.Identifier
			api.Meter
			api.PhaseCurrents
		}{
			PhoenixEVEth: base,
			Identifier: &decoratePhoenixEVEthIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decoratePhoenixEVEthMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decoratePhoenixEVEthPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargerEx == nil && identifier != nil && meter != nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*PhoenixEVEth
			api.Identifier
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
		}{
			PhoenixEVEth: base,
			Identifier: &decoratePhoenixEVEthIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decoratePhoenixEVEthMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decoratePhoenixEVEthMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decoratePhoenixEVEthPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargerEx == nil && identifier != nil && meter != nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*PhoenixEVEth
			api.Identifier
			api.Meter
			api.PhaseVoltages
		}{
			PhoenixEVEth: base,
			Identifier: &decoratePhoenixEVEthIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decoratePhoenixEVEthMeterImpl{
				meter: meter,
			},
			PhaseVoltages: &decoratePhoenixEVEthPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx == nil && identifier != nil && meter != nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*PhoenixEVEth
			api.Identifier
			api.Meter
			api.MeterEnergy
			api.PhaseVoltages
		}{
			PhoenixEVEth: base,
			Identifier: &decoratePhoenixEVEthIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decoratePhoenixEVEthMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decoratePhoenixEVEthMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decoratePhoenixEVEthPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx == nil && identifier != nil && meter != nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*PhoenixEVEth
			api.Identifier
			api.Meter
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			PhoenixEVEth: base,
			Identifier: &decoratePhoenixEVEthIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decoratePhoenixEVEthMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decoratePhoenixEVEthPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decoratePhoenixEVEthPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx == nil && identifier != nil && meter != nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*PhoenixEVEth
			api.Identifier
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			PhoenixEVEth: base,
			Identifier: &decoratePhoenixEVEthIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decoratePhoenixEVEthMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decoratePhoenixEVEthMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decoratePhoenixEVEthPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decoratePhoenixEVEthPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx != nil && identifier != nil && meter == nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*PhoenixEVEth
			api.ChargerEx
			api.Identifier
		}{
			PhoenixEVEth: base,
			ChargerEx: &decoratePhoenixEVEthChargerExImpl{
				chargerEx: chargerEx,
			},
			Identifier: &decoratePhoenixEVEthIdentifierImpl{
				identifier: identifier,
			},
		}

	case chargerEx != nil && identifier != nil && meter != nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*PhoenixEVEth
			api.ChargerEx
			api.Identifier
			api.Meter
		}{
			PhoenixEVEth: base,
			ChargerEx: &decoratePhoenixEVEthChargerExImpl{
				chargerEx: chargerEx,
			},
			Identifier: &decoratePhoenixEVEthIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decoratePhoenixEVEthMeterImpl{
				meter: meter,
			},
		}

	case chargerEx != nil && identifier != nil && meter != nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*PhoenixEVEth
			api.ChargerEx
			api.Identifier
			api.Meter
			api.MeterEnergy
		}{
			PhoenixEVEth: base,
			ChargerEx: &decoratePhoenixEVEthChargerExImpl{
				chargerEx: chargerEx,
			},
			Identifier: &decoratePhoenixEVEthIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decoratePhoenixEVEthMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decoratePhoenixEVEthMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargerEx != nil && identifier != nil && meter != nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*PhoenixEVEth
			api.ChargerEx
			api.Identifier
			api.Meter
			api.PhaseCurrents
		}{
			PhoenixEVEth: base,
			ChargerEx: &decoratePhoenixEVEthChargerExImpl{
				chargerEx: chargerEx,
			},
			Identifier: &decoratePhoenixEVEthIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decoratePhoenixEVEthMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decoratePhoenixEVEthPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargerEx != nil && identifier != nil && meter != nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*PhoenixEVEth
			api.ChargerEx
			api.Identifier
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
		}{
			PhoenixEVEth: base,
			ChargerEx: &decoratePhoenixEVEthChargerExImpl{
				chargerEx: chargerEx,
			},
			Identifier: &decoratePhoenixEVEthIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decoratePhoenixEVEthMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decoratePhoenixEVEthMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decoratePhoenixEVEthPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargerEx != nil && identifier != nil && meter != nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*PhoenixEVEth
			api.ChargerEx
			api.Identifier
			api.Meter
			api.PhaseVoltages
		}{
			PhoenixEVEth: base,
			ChargerEx: &decoratePhoenixEVEthChargerExImpl{
				chargerEx: chargerEx,
			},
			Identifier: &decoratePhoenixEVEthIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decoratePhoenixEVEthMeterImpl{
				meter: meter,
			},
			PhaseVoltages: &decoratePhoenixEVEthPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx != nil && identifier != nil && meter != nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*PhoenixEVEth
			api.ChargerEx
			api.Identifier
			api.Meter
			api.MeterEnergy
			api.PhaseVoltages
		}{
			PhoenixEVEth: base,
			ChargerEx: &decoratePhoenixEVEthChargerExImpl{
				chargerEx: chargerEx,
			},
			Identifier: &decoratePhoenixEVEthIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decoratePhoenixEVEthMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decoratePhoenixEVEthMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decoratePhoenixEVEthPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx != nil && identifier != nil && meter != nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*PhoenixEVEth
			api.ChargerEx
			api.Identifier
			api.Meter
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			PhoenixEVEth: base,
			ChargerEx: &decoratePhoenixEVEthChargerExImpl{
				chargerEx: chargerEx,
			},
			Identifier: &decoratePhoenixEVEthIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decoratePhoenixEVEthMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decoratePhoenixEVEthPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decoratePhoenixEVEthPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx != nil && identifier != nil && meter != nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*PhoenixEVEth
			api.ChargerEx
			api.Identifier
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			PhoenixEVEth: base,
			ChargerEx: &decoratePhoenixEVEthChargerExImpl{
				chargerEx: chargerEx,
			},
			Identifier: &decoratePhoenixEVEthIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decoratePhoenixEVEthMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decoratePhoenixEVEthMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decoratePhoenixEVEthPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decoratePhoenixEVEthPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}
	}

	return nil
}

type decoratePhoenixEVEthChargerExImpl struct {
	chargerEx func(float64) error
}

func (impl *decoratePhoenixEVEthChargerExImpl) MaxCurrentMillis(p0 float64) error {
	return impl.chargerEx(p0)
}

type decoratePhoenixEVEthIdentifierImpl struct {
	identifier func() (string, error)
}

func (impl *decoratePhoenixEVEthIdentifierImpl) Identify() (string, error) {
	return impl.identifier()
}

type decoratePhoenixEVEthMeterImpl struct {
	meter func() (float64, error)
}

func (impl *decoratePhoenixEVEthMeterImpl) CurrentPower() (float64, error) {
	return impl.meter()
}

type decoratePhoenixEVEthMeterEnergyImpl struct {
	meterEnergy func() (float64, error)
}

func (impl *decoratePhoenixEVEthMeterEnergyImpl) TotalEnergy() (float64, error) {
	return impl.meterEnergy()
}

type decoratePhoenixEVEthPhaseCurrentsImpl struct {
	phaseCurrents func() (float64, float64, float64, error)
}

func (impl *decoratePhoenixEVEthPhaseCurrentsImpl) Currents() (float64, float64, float64, error) {
	return impl.phaseCurrents()
}

type decoratePhoenixEVEthPhaseVoltagesImpl struct {
	phaseVoltages func() (float64, float64, float64, error)
}

func (impl *decoratePhoenixEVEthPhaseVoltagesImpl) Voltages() (float64, float64, float64, error) {
	return impl.phaseVoltages()
}
