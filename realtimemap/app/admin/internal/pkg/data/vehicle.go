package data

import (
	v1 "kratos-realtimemap/api/admin/v1"
)

const MaxPositionHistory = 200

type PositionArray []*v1.Position
type PositionMap map[string]PositionArray

func (m PositionMap) GetPositionsHistory(vehicleId string) PositionArray {
	his, ok := m[vehicleId]
	if !ok {
		return nil
	}
	return his
}

func (m PositionMap) Update(position *v1.Position) {
	_, ok := m[position.VehicleId]
	if !ok {
		his := make(PositionArray, 0, MaxPositionHistory)
		m[position.VehicleId] = his
	}

	if len(m[position.VehicleId]) > MaxPositionHistory {
		m[position.VehicleId] = m[position.VehicleId][1:]
	}
	m[position.VehicleId] = append(m[position.VehicleId], position)
}
