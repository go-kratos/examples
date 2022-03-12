package data

import (
	v1 "kratos-realtimemap/api/admin/v1"
	"sort"
)

type Organization struct {
	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Geofences CircularGeofenceArray
}

type OrganizationMap map[string]*Organization

func (o OrganizationMap) Update(position *v1.Position) TurnoverArray {
	var turnovers TurnoverArray
	for _, org := range o {
		for _, geofence := range org.Geofences {
			turnover := geofence.Update(position)
			if turnover != nil {
				turnover.OrganizationName = org.Name
				turnovers = append(turnovers, turnover)
			}
		}
	}
	return turnovers
}

type OrganizationBaseInfoArray []*v1.Organization

func (o OrganizationMap) MapToBaseInfoArray() OrganizationBaseInfoArray {
	var result = make(OrganizationBaseInfoArray, 0, len(o))

	for _, org := range o {
		if len(org.Geofences) > 0 {
			result = append(result, &v1.Organization{
				Id:   org.Id,
				Name: org.Name,
			})
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})

	return result
}

type GeofenceArray []*v1.Geofence

func (o Organization) MapToGeofenceArray() GeofenceArray {
	var geofences = make(GeofenceArray, 0, len(o.Geofences))

	for _, geofence := range o.Geofences {
		geofences = append(geofences,
			&v1.Geofence{
				Name:           geofence.Name,
				Longitude:      geofence.CentralPoint.Lng(),
				Latitude:       geofence.CentralPoint.Lat(),
				RadiusInMeters: geofence.RadiusInMeters,
				VehiclesInZone: geofence.getMapKeys(),
			})
	}

	sort.Slice(geofences, func(i, j int) bool {
		return geofences[i].Name < geofences[j].Name
	})

	return geofences
}

func NewOrganizationMapWithDefaultData() OrganizationMap {
	return OrganizationMap{
		"0006": {
			Id:   "006",
			Name: "Oy Pohjolan Liikenne Ab",
		},
		"0012": {
			Id:        "0012",
			Name:      "Helsingin Bussiliikenne Oy",
			Geofences: CircularGeofenceArray{Airport, KallioDistrict, RailwaySquare},
		},
		"0017": {
			Id:        "0017",
			Name:      "Tammelundin Liikenne Oy",
			Geofences: CircularGeofenceArray{LaajasaloIsland},
		},
		"0018": {
			Id:        "0018",
			Name:      "Pohjolan Kaupunkiliikenne Oy",
			Geofences: CircularGeofenceArray{KallioDistrict, LauttasaariIsland, RailwaySquare},
		},
		"0020": {
			Id:   "0020",
			Name: "Bus Travel Åbergin Linja Oy",
		},
		"0021": {
			Id:   "0021",
			Name: "Bus Travel Oy Reissu Ruoti",
		},
		"0022": {
			Id:        "0022",
			Name:      "Nobina Finland Oy",
			Geofences: CircularGeofenceArray{Airport, KallioDistrict, LaajasaloIsland},
		},
		"0030": {
			Id:        "0030",
			Name:      "Savonlinja Oy",
			Geofences: CircularGeofenceArray{Airport, Downtown},
		},
		"0036": {
			Id:   "0036",
			Name: "Nurmijärven Linja Oy",
		},
		"0040": {
			Id:   "0040",
			Name: "HKL-Raitioliikenne",
		},
		"0045": {
			Id:   "0045",
			Name: "Transdev Vantaa Oy",
		},
		"0047": {
			Id:   "0047",
			Name: "Taksikuljetus Oy",
		},
		"0050": {
			Id:   "0050",
			Name: "HKL-Metroliikenne",
		},
		"0051": {
			Id:   "0051",
			Name: "Korsisaari Oy",
		},
		"0054": {
			Id:   "0054",
			Name: "V-S Bussipalvelut Oy",
		},
		"0055": {
			Id:   "0055",
			Name: "Transdev Helsinki Oy",
		},
		"0058": {
			Id:   "0058",
			Name: "Koillisen Liikennepalvelut Oy",
		},
		"0060": {
			Id:   "0060",
			Name: "Suomenlinnan Liikenne Oy",
		},
		"0059": {
			Id:   "0059",
			Name: "Tilausliikenne Nikkanen Oy",
		},
		"0089": {
			Id:   "0089",
			Name: "Metropolia",
		},
		"0090": {
			Id:   "0090",
			Name: "VR Oy",
		},
		"0195": {
			Id:   "0195",
			Name: "Siuntio1",
		},
	}
}

var AllOrganizations = NewOrganizationMapWithDefaultData()
