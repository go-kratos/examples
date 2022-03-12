<template>
  <Toast />
</template>

<script lang='ts'>
import { HubConnection } from '@/hub';
import { defineComponent, PropType } from 'vue';

export default defineComponent({
  name: 'GeofencingNotifications',

  props: {
    hubConnection: {
      type: Object as PropType<HubConnection>,
      require: true,
    },
  },

  mounted() {
    if (this.hubConnection !== undefined) {
      this.hubConnection.onNotification(notification => {
        if (notification.includes('entered')) {
          this.$toast.add({
            severity: 'success',
            detail: notification,
            life: 3000,
          });
        } else {
          this.$toast.add({
            severity: 'info',
            detail: notification,
            life: 3000,
          });
        }
      });
    }
  },
});
</script>

<style scoped>
/* replace info icon with exit icon */
.p-toast-message-icon.pi.pi-info-circle:before {
  content: "\e971" !important;
}

/* replace check icon with entry icon */
.p-toast-message-icon.pi.pi-check:before {
  content: "\e970" !important;
}
</style>

<style>
/* for some reason, toast container needs to be styled globally */
.p-toast-top-right {
  top: 5.5rem !important;
}
</style>
