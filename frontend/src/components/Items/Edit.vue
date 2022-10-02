<script setup lang="ts">
import { ref, computed } from 'vue'
import { api, EmptyItem } from "../../api";
import { useStore } from '../../store';
import { useRouter, useRoute } from 'vue-router';
import InputText from 'primevue/inputtext';
import Textarea from 'primevue/textarea';
import Button from 'primevue/button';


const store = useStore();
const router = useRouter();
const route = useRoute();

let item = computed(() => {
    const id = route.params.id as string;
    const possibleItem = store.state.items.get(id);
    if (possibleItem) {
        return possibleItem;
    } else {
        router.push({name: "404"})
        return EmptyItem;
    }
});

const onSubmit = async () => {
    const updateItem = await api.updateItem(item.value)

    store.dispatch('addItem', item);
    router.push({name: 'items.detail', params: {id: item.value.id}});
};
</script>

<template>
    <div class="p-fluid formgrid grid">
        <h4>Edit Item</h4>
        <span class="p-float-label">
            <InputText id="name" type="text" v-model="item.name" />
            <label for="name">Name</label>
        </span>
        <h4>Description</h4>
        <span class="p-float-label">
            <Textarea v-model="item.description" :autoResize="true" rows="5" cols="30" />
            <label for="description">Description</label>
        </span>
        <Button label="Submit" icon="pi pi-check" @click="onSubmit()"/>
    </div>
</template>
