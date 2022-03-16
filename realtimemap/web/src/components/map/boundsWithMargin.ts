import mapboxgl from 'mapbox-gl';

export function getBoundsWithMargin(map: mapboxgl.Map) {

  const bounds = map.getBounds();

  // get bounds with 10% margin on each side (proportional to a viewbox)
  const marginPercent = 0.1;

  const sw = bounds.getSouthWest();
  const ne = bounds.getNorthEast();

  const lngMargin = Math.abs(sw.lng - ne.lng) * marginPercent;
  const latMargin = Math.abs(sw.lat - ne.lat) * marginPercent;

  return new mapboxgl.LngLatBounds(
    { lat: sw.lat - latMargin, lng: sw.lng - lngMargin },
    { lat: ne.lat + latMargin, lng: ne.lng + lngMargin },
  );
}
