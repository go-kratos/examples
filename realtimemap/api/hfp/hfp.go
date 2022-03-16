package hfp

import (
	v1 "kratos-realtimemap/api/admin/v1"
	"strings"
	"time"
)

// 具体文档: https://digitransit.fi/en/developers/apis/4-realtime-api/vehicle-positions/

type Topic struct {
	Prefix       string // /hfp/ is the root of the topic tree.
	Version      string // v2 is the current version of the HFP topic and the payload format.
	JourneyType  string // The type of the journey. Either journey, deadrun or signoff.
	TemporalType string // The status of the journey, ongoing or upcoming.

	EventType     string // One of vp, due, arr, dep, ars, pde, pas, wait, doo, doc, tlr, tla, da, dout, ba, bout, vja, vjout.
	TransportMode string // The type of the vehicle. One of bus, tram, train, ferry, metro, ubus (used by U-line buses and other vehicles with limited realtime information) or robot (used by robot buses).

	// operator_id/vehicle_number uniquely identifies the vehicle.
	OperatorId    string // The unique ID of the operator that owns the vehicle.
	VehicleNumber string // The vehicle number that can be seen painted on the side of the vehicle, often next to the front door. Different operators may use overlapping vehicle numbers.

	RouteId      string // The ID of the route the vehicle is running on.
	DirectionId  string // The line direction of the trip, either 1 or 2.
	Headsign     string // The destination name, e.g. Aviapolis.
	StartTime    string // The scheduled start time of the trip
	NextStop     string // The ID of next stop or station.
	GeohashLevel string // The geohash level represents the magnitude of change in the GPS coordinates since the previous message from the same vehicle.
	Geohash      string // The latitude and the longitude of the vehicle.
	Sid          string // Junction ID, corresponds to sid in the payload.
}

func (t *Topic) GetVehicleUID() string {
	return t.OperatorId + "." + t.VehicleNumber
}

func (t *Topic) Parse(s string) {
	//0/1       /2        /3             /4              /5           /6               /7            /8               /9         /10            /11        /12          /13         /14             /15       /16
	// /<prefix>/<version>/<journey_type>/<temporal_type>/<event_type>/<transport_mode>/<operator_id>/<vehicle_number>/<route_id>/<direction_id>/<headsign>/<start_time>/<next_stop>/<geohash_level>/<geohash>/<sid>/#

	topicParts := strings.Split(s, "/")

	//t.Prefix = topicParts[0]
	//t.Version = topicParts[1]
	//t.JourneyType = topicParts[2]
	//t.TemporalType = topicParts[3]

	t.EventType = topicParts[5]     // vp, due, arr, dep, ars, pde, pas, wait, doo, doc, tlr, tla, da, dout, ba, bout, vja, vjout
	t.TransportMode = topicParts[6] // bus, tram, train, ferry, metro, ubus
	t.OperatorId = topicParts[7]
	t.VehicleNumber = topicParts[8]

	//t.RouteId = topicParts[9]
	//t.DirectionId = topicParts[10]
	//t.Headsign = topicParts[11]
	//t.StartTime = topicParts[12]
	//t.NextStop = topicParts[13]
	//t.GeohashLevel = topicParts[14]
	//t.Geohash = topicParts[15]
	//t.Sid = topicParts[16]
}

type Payload struct {
	Longitude *float64   `json:"long"` // 经度(WGS84)
	Latitude  *float64   `json:"lat"`  // 纬度(WGS84)
	Heading   *int32     `json:"hdg"`  // 朝向角度[0, 360]
	DoorState *int32     `json:"drst"` // 门状态 0:所有门都已关闭 1:有门打开
	Timestamp *time.Time `json:"tst"`  // 时间戳
	Speed     *float64   `json:"spd"`  // 车速(m/s)
	Odometer  *int32     `json:"odo"`  // 里程(m)
}

func (p *Payload) IsValid() bool {
	return p != nil && p.Latitude != nil && p.Longitude != nil && p.Heading != nil && p.Timestamp != nil && p.Speed != nil && p.DoorState != nil
}

type Event struct {
	VehicleId  string // 车辆ID
	OperatorId string // 司机ID

	VehiclePosition *Payload `json:"VP"`  // 坐标
	DoorOpen        *Payload `json:"DOO"` // 开门
	DoorClosed      *Payload `json:"DOC"` // 关门
}

func (e *Event) GetPayload() *Payload {
	if e.VehiclePosition != nil {
		return e.VehiclePosition
	} else if e.DoorOpen != nil {
		return e.DoorOpen
	} else if e.DoorClosed != nil {
		return e.DoorClosed
	} else {
		return nil
	}
}

func (e *Event) MapToPosition() *v1.Position {
	var payload = e.GetPayload()

	if !payload.IsValid() {
		return nil
	}

	return &v1.Position{
		VehicleId: e.VehicleId,
		OrgId:     e.OperatorId,
		Latitude:  *payload.Latitude,
		Longitude: *payload.Longitude,
		Heading:   *payload.Heading,
		Timestamp: (*payload.Timestamp).UnixMilli(),
		Speed:     *payload.Speed,
	}
}
