import { MutationTree } from "vuex"
import { RootState } from "./state"
import * as api from '../api';

const Mutations: MutationTree<RootState> = {
    addItems(state, items: api.Item[]) {
        items.forEach(item => {
            state.items.set(item.id, item);
        });
    },
    addCollections(state, collections: api.Collection[]) {
        collections.forEach(collection => {
            state.collections.set(collection.id, collection);
        });
    }
}

export default Mutations