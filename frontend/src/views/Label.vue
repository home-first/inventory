<script setup lang="ts">
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://vuejs.org/api/sfc-script-setup.html#script-setup
import HelloWorld from '../components/HelloWorld.vue'
import { useRoute } from 'vue-router';
const route = useRoute();
const { label } = route.params;
</script>

<template>
  <img alt="Vue logo" src="../assets/logo.png" />
  <HelloWorld :msg=label></HelloWorld>
</template>
