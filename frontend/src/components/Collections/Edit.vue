<script setup lang="ts">
import { ref, computed } from 'vue'
import { api, EmptyCollection, EmptyItem, Item } from "../../api";
import { useStore } from '../../store';
import { useRouter, useRoute } from 'vue-router';
import InputText from 'primevue/inputtext';
import Button from 'primevue/button';
import ItemSelect from '../Items/Select.vue';


const store = useStore();
const router = useRouter();
const route = useRoute();

let collection = computed(() => {
    const id = route.params.id as string;
    const possibleItem = store.state.collections.get(id);
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
        return EmptyCollection;
    }
});

const allItems = computed(() => {
    return Array.from(store.state.items.values());
});

const name = ref(collection.value.name);
const selectedItems = ref([] as Array<Item>);

collection.value.items.forEach((id) => {
        const item = store.state.items.get(id);
        if (item) {
            return selectedItems.value.push(item);
        }
    });

const onSubmit = async () => {
    const updatedCollection = collection.value;
    updatedCollection.name = name.value;
    updatedCollection.items = selectedItems.value.map(i => i.id);

    await store.dispatch('updateCollection', updatedCollection);
    router.push({name: 'collections.detail', params: {id: collection.value.id}});
};
</script>

<template>
    <div class="p-fluid formgrid grid">
        <h4>Edit Collection</h4>
        <span class="p-float-label">
            <InputText id="name" type="text" v-model="name" />
            <label for="name">Name</label>
        </span>
        <ItemSelect v-model="selectedItems" :items="allItems" :selectable="true"/>
        <Button label="Submit" icon="pi pi-check" @click="onSubmit()"/>

    </div>
</template>
