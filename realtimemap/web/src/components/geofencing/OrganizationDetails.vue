<template>
  <h3>Geofencing zones:</h3>
  <div v-for='geofence in details.Geofences' v-bind:key='geofence.name'>
    <GeofenceDetails :geofence='geofence'></GeofenceDetails>
  </div>
</template>

<script lang='ts'>
import { defineComponent, PropType } from 'vue';
import { OrganizationDetailsDto } from './api-organization';
import GeofenceDetails from './GeofenceDetails.vue';

export default defineComponent({

  name: 'OrganizationDetails',

  components: {
    GeofenceDetails,
  },

  props: {
    details: {
      type: Object as PropType<OrganizationDetailsDto>,
      required: true,
    },
  },

  computed: {
    anythingToDisplay(): boolean {
      return Object.keys(this.details).length > 0;
    },
  },

});
</script>
