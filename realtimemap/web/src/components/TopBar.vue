<template>

  <div class='w-full bg-primary text-0 p-3 flex align-items-center'>

    <i class='pi pi-map text-3xl'></i>

    <div class='ml-3 flex flex-column'>
      <div class='font-semibold'>
        Real-time Map
      </div>
      <div class='text-sm'>
        Kratos showcase
      </div>
    </div>

    <div class='flex-1'>
    </div>

    <a href='https://github.com/tx7do/kratos-realtimemap' target='_blank' rel='noopener noreferrer'
       class='h-full px-3 block'>
      <i class='pi pi-github text-3xl'></i>
    </a>

  </div>

</template>

<script lang='ts'>
import { defineComponent } from 'vue';

export default defineComponent({});
</script>

<style scoped>

.protoactor-logo {
  height: 2rem;
}

.protoactor-logo:hover {
  opacity: .75;
}

a:link {
  color: var(--surface-0);
}

a:visited {
  color: var(--surface-0);
}

a:hover {
  color: var(--cyan-100);
}

a:active {
  color: var(--cyan-100);
}

</style>
