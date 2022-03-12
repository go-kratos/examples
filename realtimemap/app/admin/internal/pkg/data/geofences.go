package data

import (
	geo "github.com/kellydunn/golang-geo"
	v1 "kratos-realtimemap/api/admin/v1"
)

// CircularGeofence 圆形地理围栏
type CircularGeofence struct {
	Name           string    // 名字
	CentralPoint   geo.Point // 中心点
	RadiusInMeters float64   // 米半径
	VehiclesInZone map[string]struct{}
}

type CircularGeofenceArray []*CircularGeofence

// IncludesPosition 是否包含了该点
func (geofence *CircularGeofence) IncludesPosition(latitude float64, longitude float64) bool {
	point := geo.NewPoint(latitude, longitude)
	return geofence.CentralPoint.GreatCircleDistance(point)*1000 < geofence.RadiusInMeters
}

func (geofence *CircularGeofence) getMapKeys() []string {
	keys := make([]string, len(geofence.VehiclesInZone))

	i := 0
	for k := range geofence.VehiclesInZone {
		keys[i] = k
		i++
	}

	return keys
}

type Turnover struct {
	Status           bool
	VehicleId        string
	OrganizationName string
	GeofenceName     string
}
type TurnoverArray []*Turnover

func (geofence *CircularGeofence) Update(position *v1.Position) *Turnover {
	_, vehicleIsInZone := geofence.VehiclesInZone[position.VehicleId]
	if geofence.IncludesPosition(position.Latitude, position.Longitude) {
		if !vehicleIsInZone {
			// 进入区域
			//fmt.Println(position.VehicleId, "IN")
			geofence.VehiclesInZone[position.VehicleId] = struct{}{}
			return &Turnover{Status: true, VehicleId: position.VehicleId, GeofenceName: geofence.Name}
		}
	} else {
		if vehicleIsInZone {
			// 离开区域
			//fmt.Println(position.VehicleId, "OUT")
			delete(geofence.VehiclesInZone, position.VehicleId)
			return &Turnover{Status: false, VehicleId: position.VehicleId, GeofenceName: geofence.Name}
		}
	}
	return nil
}

var (
	Airport = &CircularGeofence{
		Name:           "Airport",
		CentralPoint:   *geo.NewPoint(60.31146, 24.96907),
		RadiusInMeters: 2000,
		VehiclesInZone: map[string]struct{}{},
	}

	Downtown = &CircularGeofence{
		Name:           "Downtown",
		CentralPoint:   *geo.NewPoint(60.16422983026082, 24.941068845053014),
		RadiusInMeters: 1700,
		VehiclesInZone: map[string]struct{}{},
	}

	RailwaySquare = &CircularGeofence{
		Name:           "Railway Square",
		CentralPoint:   *geo.NewPoint(60.171285, 24.943936),
		RadiusInMeters: 150,
		VehiclesInZone: map[string]struct{}{},
	}

	LauttasaariIsland = &CircularGeofence{
		Name:           "Lauttasaari island",
		CentralPoint:   *geo.NewPoint(60.158536, 24.873788),
		RadiusInMeters: 1400,
		VehiclesInZone: map[string]struct{}{},
	}

	LaajasaloIsland = &CircularGeofence{
		Name:           "Laajasalo island",
		CentralPoint:   *geo.NewPoint(60.16956184470527, 25.052851825093114),
		RadiusInMeters: 2200,
		VehiclesInZone: map[string]struct{}{},
	}

	KallioDistrict = &CircularGeofence{
		Name:           "Kallio district",
		CentralPoint:   *geo.NewPoint(60.18260263288996, 24.953588638997264),
		RadiusInMeters: 600,
		VehiclesInZone: map[string]struct{}{},
	}
)
