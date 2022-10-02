import { InjectionKey } from 'vue'
import { createStore, Store, useStore as baseUseStore } from 'vuex';
import { RootState, startingState } from './state';
import * as api from '../api';
import Mutations from './mutations';
import Actions from './actions';
import Getters from './getters';

export const key: InjectionKey<Store<RootState>> = Symbol()

export const store = createStore<RootState>({
    state: startingState,
    mutations: Mutations,
    actions: Actions,
    getters: Getters,
});

export function useStore() {
    return baseUseStore(key)
}
