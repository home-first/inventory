<script setup lang="ts">
import { ref } from 'vue'
import { useStore } from '../../store';
import { useRouter, useRoute } from 'vue-router';
import { computed } from '@vue/reactivity';
import { EmptyCollection, EmptyItem, Item } from '../../api';
import ItemSelect from '../Items/Select.vue';


const store = useStore();
const router = useRouter();
const route = useRoute();
const collection = computed(() => {
    const id = route.params.id as string;
    const possibleCollection = store.state.collections.get(id);
    if (possibleCollection) {
        return possibleCollection;
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
        return EmptyCollection;
    }
});

const items = computed(() => {
    return collection.value.items.map((id) => {
        const item = store.state.items.get(id);
        if (item) {
            return item;
        } else {
            console.log("Item not found in store");
            return EmptyItem;
        }
    });
});
</script>

<template>
    <div>
        <h4>{{ collection.name }}</h4>

        <ItemSelect :items="items" />
    </div>
</template>
