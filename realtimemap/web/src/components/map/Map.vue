<template>
  <div>
    <div id='map'></div>
  </div>
</template>

<script lang='ts'>
import { defineComponent, PropType } from 'vue';
import mapboxgl from 'mapbox-gl';
import 'mapbox-gl/dist/mapbox-gl.css';
import mapboxConfig from '@/mapboxConfig';
import { addVehicleTrailLayer } from './vehicleTrailsLayer';
import { addGeofencesLayer, Geofence, setGeofences } from './geofencesLayer';
import { VehicleStates, handlePositionEvent, clearVehiclesOutsideOfViewbox } from './vehicleStates';
import { addVehicleDetailsPopup } from './vehicleDetailsPopup';
import { addVehiclesLayer } from './vehiclesLayer';
import { addVehicleClustersLayer } from './vehicleClustersLayer';
import { handleViewportUpdates } from './viewportUpdates';
import { HubConnection } from '@/hub';

export default defineComponent({
  name: 'Map',

  props: {
    geofences: {
      type: Array as PropType<Geofence[]>,
      require: true,
    },
    hubConnection: {
      type: Object as PropType<HubConnection>,
      require: true,
    },
  },

  data() {
    return {
      // it will be set in mounted and available later on
      map: undefined as unknown as mapboxgl.Map,
    };
  },

  mounted() {
    mapboxgl.accessToken = mapboxConfig.getAccessToken();

    this.map = new mapboxgl.Map({
      container: 'map',
      style: 'mapbox://styles/mapbox/streets-v11',
      center: [24.938, 60.169],
      zoom: 8,
    });

    const vehicleStates: VehicleStates = {};

    if (this.hubConnection !== undefined) {
      this.hubConnection.onPositions(positions => {
        for (const position of positions) {
          handlePositionEvent(vehicleStates, position);
        }
      });

    }

    this.map.on('load', () => {

      addVehicleClustersLayer(this.map, vehicleStates);
      addVehiclesLayer(this.map, vehicleStates);
      addVehicleTrailLayer(this.map);
      addGeofencesLayer(this.map);

      addVehicleDetailsPopup(this.map);

      if (this.hubConnection !== undefined)
        handleViewportUpdates(this.map, this.hubConnection);

      setInterval(
        () => clearVehiclesOutsideOfViewbox(this.map, vehicleStates),
        5000,
      );

    });

  },

  watch: {
    geofences(newGeofences: Geofence[] | undefined) {
      setGeofences(this.map, newGeofences);
    },
  },

});
</script>

<style>
body {
  margin: 0;
  padding: 0;
}

#map {
  position: relative;
  /* top: 20px; */
  height: 100%;
  width: 100%;
}
</style>
