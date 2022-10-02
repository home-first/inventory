<script setup lang="ts">
import { ref } from 'vue'
import { useStore } from '../../store';
import { useRouter, useRoute } from 'vue-router';
import { computed } from '@vue/reactivity';


const store = useStore();
const router = useRouter();
const route = useRoute();
store.dispatch('getItem', route.params.id).then(()=>{
    
});
const item = computed(() => {
    const id = route.params.id as string;
    return store.state.items.get(id);
});
</script>

<template>
    <div v-if="item !== undefined ">
        <h4>{{ item.name }}</h4>

        <p>Associated Collections {{ item.collections.length }}</p>
    </div>
    <div v-else>
        404
    </div>
</template>
