<script setup lang="ts">
import { ref } from 'vue'
import { useStore } from '../../store';
import { useRouter } from 'vue-router';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Toolbar from 'primevue/toolbar';
import Button from 'primevue/button';
import { computed } from '@vue/reactivity';
import type { Collection } from '../../api';


const store = useStore();
const router = useRouter();
store.dispatch('getItems');
store.dispatch('getCollections');
const collections = computed(() => Array.from(store.state.collections.values()));
let selectedCollections = ref(null);
const editCollection = (collection: Collection) => {
    router.push({name: 'collections.edit', params: {id: collection.id}});
};
const dt = ref();
const exportCSV = () => {
    dt.value.exportCSV();
};

const newCollection = () => {
    router.push({name: 'collections.new'})
};
</script>

<template>
    <div>
        <Toolbar>
            <template #start>
                <Button label="New Item" icon="pi pi-plus" class="p-button-success mr-2" @click="newCollection()" />
            </template>

            <template #end>
                <Button label="Export" icon="pi pi-upload" class="p-button-help" @click="exportCSV()" />
            </template>
        </Toolbar>
        <DataTable ref="dt" :value="collections" responsiveLayout="scroll" v-model="selectedCollections">
            <Column selectionMode="multiple" style="width: 3rem" :exportable="false"></Column>
            <Column field="id" header="ID">
                <template #body="slotProps">
                    <RouterLink :to="{name: 'collections.detail', params: {id: slotProps.data.id}}">
                        {{slotProps.data.id.substr(0, 8)}}
                    </RouterLink>
                </template>
            </Column>
            <Column field="name" header="Name" />
            <Column field="description" header="Description" />
            <Column :exportable="false">
                <template #body="slotProps">
                    <Button icon="pi pi-pencil" class="p-button-rounded p-button-success mr-2"
                        @click="editCollection(slotProps.data)" />
                </template>
            </Column>
        </DataTable>
    </div>
</template>
