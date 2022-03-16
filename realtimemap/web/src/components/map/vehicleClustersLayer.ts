import { VehicleStates, mapVehiclesToGeoJson } from './vehicleStates';
import { trySetGeoJsonSource } from './mapUtils';
import { showMarkerLevel } from './vehiclesLayer';

const vehicleClusterSourceId = 'vehicle-clusters';

export const addVehicleClustersLayer = (map: mapboxgl.Map, vehicleStates: VehicleStates) => {

  map.addSource(vehicleClusterSourceId, {
    type: 'geojson',
    data: {
      type: 'FeatureCollection',
      features: [],
    },
    cluster: true,
    clusterMaxZoom: showMarkerLevel, // Max zoom to cluster points on
    clusterRadius: 50, // Radius of each cluster when clustering points (defaults to 50)
  });

  map.addLayer({
    id: 'vehicle-cluster-layer',
    type: 'circle',
    source: vehicleClusterSourceId,
    maxzoom: showMarkerLevel,
    filter: ['has', 'point_count'],
    paint: {
      'circle-color': [
        'step',
        ['get', 'point_count'],
        '#9c313a',
        100,
        '#0c7186',
        750,
        '#73a824',
      ],
      'circle-radius': ['step', ['get', 'point_count'], 20, 100, 30, 750, 40],
    },
  });

  map.addLayer({
    id: 'vehicle-cluster-count-layer',
    type: 'symbol',
    source: vehicleClusterSourceId,
    maxzoom: showMarkerLevel,
    filter: ['has', 'point_count'],
    layout: {
      'text-field': '{point_count_abbreviated}',
      'text-font': ['DIN Offc Pro Medium', 'Arial Unicode MS Bold'],
      'text-size': 12,
    },
    paint: {
      'text-color': '#ffffff',
    },
  });

  setInterval(
    () => updateClusterLayers(map, vehicleStates),
    5000,
  );

};

function updateClusterLayers(map: mapboxgl.Map, vehicleStates: VehicleStates) {
  const data = mapVehiclesToGeoJson(vehicleStates, () => true);
  trySetGeoJsonSource(map, vehicleClusterSourceId, data);
}
