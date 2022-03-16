<template>
  <div class='m-3'>
    <h2 class='mt-0 mb-0'>
      Geofencing
    </h2>
    <div class='mt-6'>
      <div class='p-field'>
        <span class='p-float-label'>
          <Dropdown id='organization-select' v-model='selectedItem' :options='organizations'
                    optionLabel='name' />
          <label for='organization-select'>Select an organization</label>
        </span>
      </div>
    </div>
    <div class='mt-5' v-if='details'>
      <OrganizationDetails :details='details' />
    </div>
  </div>
</template>

<script lang='ts'>
import { defineComponent } from 'vue';
import {
  browseOrganizations,
  OrganizationDto,
  OrganizationDetailsDto,
  getDetails,
} from './api-organization';
import OrganizationDetails from './OrganizationDetails.vue';

export default defineComponent({
  data() {
    return {
      organizations: [] as OrganizationDto[],
      details: null as OrganizationDetailsDto | null,
      selectedItem: null as OrganizationDto | null,
    };
  },

  watch: {
    selectedItem(selectedOrganization) {
      if (selectedOrganization) {
        this.getDetails(selectedOrganization.id);
      }
    },
  },

  components: {
    OrganizationDetails,
  },

  created() {
    this.fetchData();
  },

  methods: {
    async fetchData() {
      this.organizations = await browseOrganizations();
      // console.log(this.organizations)
    },
    async getDetails(id: string) {
      this.details = await getDetails(id);

      const geofences = this.details
        .Geofences
        .map(geofence => ({
          long: geofence.longitude,
          lat: geofence.latitude,
          radiusInMeters: geofence.radiusInMeters,
        }));

      this.$emit('geofences-updated', geofences);
    },
  },

});
</script>

<style scoped>
.p-dropdown {
  width: 100%;
}
</style>
