<script setup lang="ts">
import { ref } from 'vue'
import { useStore } from '../../store';
import { useRouter } from 'vue-router';
import DataTable, { DataTableEmits } from 'primevue/datatable';
import Column from 'primevue/column';
import Toolbar from 'primevue/toolbar';
import Button from 'primevue/button';
import { computed } from '@vue/reactivity';
import type { Item } from '../../api';

interface Props {
    items: Array<Item>;
    selectable?: boolean;
    modelValue?: Array<Item>;
}

const props = withDefaults(defineProps<Props>(), {
    selectable: false,
    modelValue: () => [] as Array<Item>,
});

let emit = defineEmits<{
    (event: 'update:modelValue', payload: Array<Item>): void;
}>();

const onSelectionUpdate = (...params: Parameters<DataTableEmits['update:selection']>) => {
    emit('update:modelValue', params[0]);
};


</script>

<template>
    <div>
        <DataTable :value="props.items" responsiveLayout="scroll" :selection="props.modelValue"
            @update:selection="onSelectionUpdate">
            <Column selectionMode="multiple" headerStyle="width: 3em" v-if="props.selectable"></Column>
            <Column field="id" header="ID">
                <template #body="slotProps">
                    <RouterLink :to="{name: 'items.detail', params: {id: slotProps.data.id}}">
                        {{slotProps.data.id.substr(0, 8)}}
                    </RouterLink>
                </template>
            </Column>
            <Column field="name" header="Name" />
            <Column field="description" header="Description" />
        </DataTable>
    </div>
</template>
