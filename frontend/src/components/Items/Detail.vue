<script setup lang="ts">
import { ref } from 'vue'
import { useStore } from '../../store';
import { useRouter, useRoute } from 'vue-router';
import { computed } from '@vue/reactivity';
import { EmptyItem } from '../../api';


const store = useStore();
const router = useRouter();
const route = useRoute();
store.dispatch('getItem', route.params.id).then(()=>{
    
});
const item = computed(() => {
    const id = route.params.id as string;
    const possibleItem = store.state.items.get(id);
    if (possibleItem) {
        return possibleItem;
    } else {
        router.push({
            name: "404",
            params: {
                pathMatch: route.path.substring(1).split('/')
            },
            query: route.query,
            hash: route.hash,
            }
        );
        return EmptyItem;
    }
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
