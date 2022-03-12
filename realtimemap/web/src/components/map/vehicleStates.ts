import { PositionDto } from '@/hub';
import { getBoundsWithMargin } from './boundsWithMargin';
import { GeoJSONSourceData } from './mapUtils';

// const stepsInAnimation = 10;

// this must have this shape to be compatible with mapbox
export interface VehiclePosition {
  lat: number;
  lng: number;
  heading: number;
}

export interface VehicleState {
  vehicleId: string;
  steps: number;
  delta: VehiclePosition;
  currentPosition: VehiclePosition;
  speed: number;
  nextPosition: VehiclePosition;
  shouldAnimate: boolean;
  icon: string;
}

export type VehicleStates = { [vehicleId: string]: VehicleState };

export function handlePositionEvent(vehicleStates: VehicleStates, positionDto: PositionDto) {
  if (!vehicleStates[positionDto.vehicle_id]) {
    console.log("handlePositionEvent", positionDto)
    vehicleStates[positionDto.vehicle_id] = createVehicleFromState(positionDto);
  }

  updateVehicleFromEvent(vehicleStates, positionDto);
}

function createVehicleFromState(e: PositionDto): VehicleState {
  return {
    vehicleId: e.vehicle_id,
    speed: 0,
    steps: 0,
    nextPosition: { lng: e.longitude, lat: e.latitude, heading: e.heading },
    currentPosition: { lng: e.longitude, lat: e.latitude, heading: e.heading },
    delta: { lat: 0, lng: 0, heading: 0 },
    shouldAnimate: false,
    icon: '',
  };
}

function updateVehicleFromEvent(
  vehicleStates: VehicleStates,
  positionDto: PositionDto,
) {
  const vehicleState = vehicleStates[positionDto.vehicle_id];

  vehicleState.currentPosition = {
    lat: positionDto.latitude,
    lng: positionDto.longitude,
    heading: positionDto.heading,
  };

  // console.log(positionDto.vehicleId)
  // console.log(vehicleState.nextPosition);
  // console.log(positionDto);

  // const lng = positionDto.longitude - vehicleState.nextPosition.lng;
  // const lat = positionDto.latitude - vehicleState.nextPosition.lat;
  // let heading = positionDto.heading - vehicleState.nextPosition.heading;

  // if (lng != 0) {
  //   console.log(lng, lat);
  // }

  // //prevent full rotations when next and current course cross between 0 and 360
  // if (heading > 180) {
  //   heading -= 360;
  // }

  // if (heading < -180) {
  //   heading += 360;
  // }

  // vehicleState.steps = stepsInAnimation;
  // vehicleState.delta = {
  //   lng: lng / stepsInAnimation,
  //   lat: lat / stepsInAnimation,
  //   heading: heading / stepsInAnimation,
  // };
  // //console.log(vehicleState.delta );
  // vehicleState.currentPosition = vehicleState.nextPosition;
  // vehicleState.nextPosition = {
  //   lng: positionDto.longitude,
  //   lat: positionDto.latitude,
  //   heading: positionDto.heading,
  // };
  // vehicleState.shouldAnimate =
  //   vehicleState.delta.lat != 0 ||
  //   vehicleState.delta.lng != 0 ||
  //   vehicleState.delta.heading != 0;

  if (positionDto.doors_open) {
    //console.log("doors open...")
    vehicleState.icon = 'doors-open';
  } else if (
    (positionDto.speed != undefined && positionDto.speed > 0) ||
    vehicleState.shouldAnimate
  ) {
    vehicleState.icon = 'moving';
  } else {
    // todo: use better icon
    vehicleState.icon = 'moving';
  }
}

export function mapVehiclesToGeoJson(
  vehicleStates: VehicleStates,
  predicate: (vehicleState: VehicleState) => boolean,
): GeoJSONSourceData {
  return {
    type: 'FeatureCollection',
    features: Object.values(vehicleStates)
      .filter(predicate)
      .map((vehicleState) => ({
        type: 'Feature',
        geometry: {
          type: 'Point',
          coordinates: [
            vehicleState.currentPosition.lng,
            vehicleState.currentPosition.lat,
          ],
        },
        properties: {
          course: vehicleState.currentPosition.heading,
          vehicleId: vehicleState.vehicleId,
          speed: vehicleState.speed,
          icon: vehicleState.icon,
        },
      })),
  };
}

export function clearVehiclesOutsideOfViewbox(map: mapboxgl.Map, vehicleStates: VehicleStates) {

  const biggerBounds = getBoundsWithMargin(map);

  Object.values(vehicleStates)
    .filter(vehicleState => !biggerBounds.contains(vehicleState.currentPosition))
    .map(vehicleState => vehicleState.vehicleId)
    .forEach(id => delete vehicleStates[id]);

}
