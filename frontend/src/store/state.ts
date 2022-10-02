import * as api from '../api';

export interface RootState {
    items: Map<string, api.Item>;
    collections: Map<string, api.Collection>;
}

export const startingState: RootState = {
    items: new Map<string, api.Item>(),
    collections: new Map<string, api.Collection>(),
};