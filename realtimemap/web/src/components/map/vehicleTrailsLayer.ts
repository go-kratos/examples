import { GetTrail } from '@/components/map/api-trail';
import mapboxgl from 'mapbox-gl';
import { trySetGeoJsonSource } from './mapUtils';
import { vehicleLayerId } from './vehiclesLayer';

const vehicleTrailSourceId = 'vehicle-trails';

export const addVehicleTrailLayer = async (map: mapboxgl.Map) => {

  map.addSource(vehicleTrailSourceId, {
    type: 'geojson',
    data: {
      type: 'Feature',
      properties: {},
      geometry: {
        type: 'LineString',
        coordinates: [],
      },
    },
  });

  map.addLayer({
    id: 'vehicle-trail-layer',
    type: 'line',
    source: vehicleTrailSourceId,
    layout: {
      'line-join': 'round',
      'line-cap': 'round',
    },
    paint: {
      'line-color': '#ed4981',
      'line-width': 8,
    },
  });

  let currentlySelectedVehicleId: string | null = null;

  async function drawTrail(vehicleId: string) {
    const trail = await GetTrail(vehicleId);

    trySetGeoJsonSource(map, vehicleTrailSourceId, {
      type: 'Feature',
      properties: {},
      geometry: {
        type: 'LineString',
        coordinates: trail.positions.map((position) => {
          return [position.longitude, position.latitude];
        }),
      },
    });
  }

  async function drawCurrentlySelectedVehicleTrail() {
    if (currentlySelectedVehicleId) {
      drawTrail(currentlySelectedVehicleId);
    }
  }

  map.on('click', vehicleLayerId, async e => {
    const features = map.queryRenderedFeatures(e.point);
    const feature = features[0];
    if (feature != null && feature.properties != null) {
      currentlySelectedVehicleId = feature.properties.vehicleId;
      await drawCurrentlySelectedVehicleTrail();
    }
  });

  setInterval(
    () => drawCurrentlySelectedVehicleTrail(),
    2500,
  );

};
