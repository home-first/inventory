import { createStore } from 'vuex';
import { RootState } from './state';
import * as api from '../api';


export const store = createStore<RootState>({
    state: {
        items: [],
    },
    mutations: {
        addItems(state, items: api.Item[]) {
            state.items.concat(items);
            state.items.sort((a, b) => {
                return a.id - b.id;
            });
        }
    }
});