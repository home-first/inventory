<script setup lang="ts">
import { ref } from 'vue'
import { api } from "../../api";
import { useStore } from '../../store';
import { useRouter } from 'vue-router';
import InputText from 'primevue/inputtext';
import Textarea from 'primevue/textarea';
import Button from 'primevue/button';


const store = useStore();
const router = useRouter();

let name = ref('');
let description = ref('');

const onSubmit = async () => {
    const item = await api.createItem({
        name: name.value,
        description: description.value,
        collections: [] //TODO: add collections later
    })

    store.dispatch('addItem', item);
    router.push({name: 'items.detail', params: {id: item.id}});
};
</script>

<template>
    <div class="p-fluid formgrid grid">
        <span class="p-float-label">
            <InputText id="name" type="text" v-model="name" />
            <label for="name">Name</label>
        </span>
        <span class="p-float-label">
            <Textarea v-model="description" :autoResize="true" rows="5" cols="30" />
            <label for="description">Description</label>
        </span>
        <Button label="Submit" icon="pi pi-check" @click="onSubmit()"/>
    </div>
</template>
