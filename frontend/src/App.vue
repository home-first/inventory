<script setup lang="ts">
import { useRouter } from 'vue-router';
import Menubar from 'primevue/menubar';
import InputText from 'primevue/inputtext';
import type { MenuItem } from 'primevue/menuitem';

const router = useRouter();
const items: MenuItem[] = [
  {
    label: 'Items',
    command: () => {
      router.push({ name: 'items.list'})
    },
  }
]
</script>

<template>
  <Menubar :model="items">
    <template #end>
      <InputText placeholder="Search" type="text" />
    </template>
  </Menubar>
  <router-view />
</template>

<style>
/* .p-menubar {
    position: sticky;
    position: -webkit-sticky;
    z-index: 10000;
    top: 0;
} */
</style>

