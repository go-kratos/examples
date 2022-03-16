import { VehicleStates, mapVehiclesToGeoJson } from './vehicleStates';
import { getBoundsWithMargin } from './boundsWithMargin';
import { trySetGeoJsonSource } from './mapUtils';

export const showMarkerLevel = 12;

const vehicleSourceId = 'vehicles';
export const vehicleLayerId = 'vehicle-layer';

export const addVehiclesLayer = (map: mapboxgl.Map, vehicleStates: VehicleStates) => {

  map.addSource(vehicleSourceId, {
    type: 'geojson',
    data: {
      type: 'FeatureCollection',
      features: [],
    },
  });

  map.addLayer({
    id: vehicleLayerId,
    type: 'symbol',
    source: vehicleSourceId,
    minzoom: showMarkerLevel,
    layout: {
      'icon-image': ['get', 'icon'],
      'icon-size': ['interpolate', ['linear'], ['zoom'], 9, 0.05, 15, 0.4],
      'icon-allow-overlap': true,
      'icon-rotate': ['get', 'course'],
    },
  });

  map.loadImage('/vehicle-moving.png', (error, image) => {
    if (error) throw error;
    map.addImage('moving', image);
  });

  map.loadImage('/vehicle-doors-open.png', (error, image) => {
    if (error) throw error;
    map.addImage('doors-open', image);
  });

  setInterval(
    () => updateVehicleLayers(map, vehicleStates),
    1000,
  );

};

function updateVehicleLayers(map: mapboxgl.Map, vehicleStates: VehicleStates) {
  if (map.getZoom() < showMarkerLevel) {
    return;
  }

  // expand viewport so we ingest things just outside the bounds also.
  const biggerBounds = getBoundsWithMargin(map);

  const data = mapVehiclesToGeoJson(vehicleStates, vehicle =>
    biggerBounds.contains(vehicle.currentPosition),
  );

  trySetGeoJsonSource(map, vehicleSourceId, data);
}
