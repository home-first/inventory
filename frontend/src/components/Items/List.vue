<script setup lang="ts">
import { ref } from 'vue'
import { useStore } from '../../store';
import { useRouter } from 'vue-router';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Toolbar from 'primevue/toolbar';
import Button from 'primevue/button';
import { computed } from '@vue/reactivity';
import type { Item } from '../../api';


const store = useStore();
const router = useRouter();
store.dispatch('getItems');
store.dispatch('getCollections');
const items = computed(() => Array.from(store.state.items.values()));
let selectedItems = ref(null);
const editItem = (item: Item) => {
    router.push({name: 'items.edit', params: {id: item.id}});
};
const dt = ref();
const exportCSV = () => {
    dt.value.exportCSV();
};

const newItem = () => {
    router.push({name: 'items.new'})
};
</script>

<template>
    <Toolbar>
        <template #start>
            <Button label="New" icon="pi pi-plus" class="p-button-success mr-2" @click="newItem()" />
        </template>

        <template #end>
            <Button label="Export" icon="pi pi-upload" class="p-button-help" @click="exportCSV()" />
        </template>
    </Toolbar>
    <DataTable ref="dt" :value="items" responsiveLayout="scroll" v-model:selection="selectedItems">
        <Column selectionMode="multiple" style="width: 3rem" :exportable="false"></Column>
        <Column field="id" header="ID">
            <template #body="slotProps">
                <RouterLink :to="{name: 'items.detail', params: {id: slotProps.data.id}}">
                    {{slotProps.data.id.substr(0, 8)}}
                </RouterLink>
            </template>
        </Column>
        <Column field="name" header="Name" />
        <Column field="description" header="Description" />
        <Column :exportable="false">
            <template #body="slotProps">
                <Button icon="pi pi-pencil" class="p-button-rounded p-button-success mr-2"
                    @click="editItem(slotProps.data)" />
            </template>
        </Column>
    </DataTable>
</template>
