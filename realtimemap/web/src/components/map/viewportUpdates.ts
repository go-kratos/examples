import { HubConnection } from '@/hub';
import { throttle } from 'lodash';

export const handleViewportUpdates = (map: mapboxgl.Map, connection: HubConnection) => {

  const throttledUpdateViewport = throttle(setViewport, 1000);

  map.on('zoomend', () => {
    throttledUpdateViewport(map, connection);
  });

  map.on('move', () => {
    throttledUpdateViewport(map, connection);
  });

  setTimeout(
    () => setViewport(map, connection),
    500,
  );

};

function setViewport(map: mapboxgl.Map, connection: HubConnection) {
  const bounds = map.getBounds();
  const sw = bounds.getSouthWest();
  const ne = bounds.getNorthEast();
  connection.setViewport(sw.lng, sw.lat, ne.lng, ne.lat);
}
