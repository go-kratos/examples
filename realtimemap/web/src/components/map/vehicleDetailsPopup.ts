import mapboxgl from 'mapbox-gl';
import { Point } from 'geojson';
import { vehicleLayerId } from './vehiclesLayer';

export const addVehicleDetailsPopup = (map: mapboxgl.Map) => {

  const popup = new mapboxgl.Popup({
    closeButton: false,
    closeOnClick: false,
  });

  map.on('mouseenter', vehicleLayerId, e => {
    if (!e.features || e.features.length < 1) return;

    // Change the cursor style as a UI indicator.
    map.getCanvas().style.cursor = 'pointer';

    const coordinates = (e.features[0].geometry as Point).coordinates as [number, number];
    const description = e.features[0].properties?.vehicleId;

    // Ensure that if the map is zoomed out such that multiple
    // copies of the feature are visible, the popup appears
    // over the copy being pointed to.
    while (Math.abs(e.lngLat.lng - coordinates[0]) > 180) {
      coordinates[0] += e.lngLat.lng > coordinates[0] ? 360 : -360;
    }

    // Populate the popup and set its coordinates
    // based on the feature found.
    popup.setLngLat(coordinates).setHTML(description).addTo(map);
  });

  map.on('mouseleave', vehicleLayerId, function() {
    map.getCanvas().style.cursor = '';
    popup.remove();
  });

};
