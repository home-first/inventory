import { ActionTree } from "vuex";
import { RootState } from "./state";
import { api, Collection } from "../api";


const Actions: ActionTree<RootState, RootState> = {
    getItems({ commit }) {
        api.getItems().then(items => {
            commit("addItems", items);
        });
    },
    getItem({ commit }, id) {
        api.getItem(id).then(item => {
            commit("addItems", [item]);
        }).catch(error => {
            console.log(error);
        }
        );
    },
    getCollections({ commit } ) {
        api.getCollections().then(collections => {
            commit("addCollections", collections);
        }).catch(error => {
            console.log(error);
        }
        );
    },
    getCollection({ commit }, id) {
        api.getCollection(id).then(collection => {
            commit("addCollections", [collection]);
        }).catch(error => {
            console.log(error);
        }
        );
    },
    updateCollection({ commit}, collection: Collection) {
        api.updateCollection(collection).then(collection => {
            commit("addCollections", [collection]);
        })
    }
}

export default Actions;