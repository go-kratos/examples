<template>

  <div class='flex flex-row'>
    <GeofencingOrganizations class='flex-1'
                             @geofences-updated="$emit('geofences-updated', $event)" />
    <div class='notifications-placeholder' />
  </div>

</template>

<script lang='ts'>

import { defineComponent } from 'vue';
import GeofencingOrganizations from './GeofencingOrganizations.vue';

export default defineComponent({
  name: 'GeofencingPanel',

  components: {
    GeofencingOrganizations,
  },
});

</script>

<style scoped>

.notifications-placeholder {
  width: 27rem;
}

</style>
