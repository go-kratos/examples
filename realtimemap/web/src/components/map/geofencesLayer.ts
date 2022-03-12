import mapboxgl from 'mapbox-gl';
import turfCircle from '@turf/circle';
import { trySetGeoJsonSource } from './mapUtils';

export interface Geofence {
  long: number;
  lat: number;
  radiusInMeters: number;
}

const geofencesSourceId = 'geofences';

export const addGeofencesLayer = (map: mapboxgl.Map) => {

  map.addSource(geofencesSourceId, {
    type: 'geojson',
    data: {
      type: 'FeatureCollection',
      features: [],
    },
  });

  map.addLayer({
    id: 'geofences',
    type: 'line',
    source: geofencesSourceId,
    layout: {},
    paint: {
      'line-color': '#9c27b0',
      'line-width': 5,
    },
  });

};

export const setGeofences = (map: mapboxgl.Map, geofences: Geofence[] | undefined) => {
  trySetGeoJsonSource(map, geofencesSourceId, {
    type: 'FeatureCollection',
    features: mapGeofencesToPolygons(geofences),
  });
};

const mapGeofencesToPolygons = (geofences: Geofence[] | undefined): GeoJSON.Feature<GeoJSON.Polygon, GeoJSON.GeoJsonProperties>[] => {
  return geofences
    ? geofences.map(mapGeofenceToPolygon)
    : [];
};

const mapGeofenceToPolygon = (geofence: Geofence): GeoJSON.Feature<GeoJSON.Polygon, GeoJSON.GeoJsonProperties> => {
  const radiusInKilometers = geofence.radiusInMeters / 1000;
  return turfCircle([geofence.long, geofence.lat], radiusInKilometers, {
    steps: 25,
    units: 'kilometers',
  });
};
