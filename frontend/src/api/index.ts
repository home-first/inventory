import { API_BASE_PATH } from "./const";
import * as Items from "./items";
import * as Collections from "./collections";

export * from "./items";
export * from "./collections";

export class Api {
    async getItems(): Promise<Items.Item[]> {
        const itemResponse = await fetch(`${API_BASE_PATH}/items`, {
            method: "GET",
        });
        return itemResponse.json();
    }

    async getItem(id: string): Promise<Items.Item> {
        const itemResponse = await fetch(`${API_BASE_PATH}/items/${id}`, {
            method: "GET",
        });
        return itemResponse.json();
    }

    async createItem(item: Omit<Items.Item, 'id'>): Promise<Items.Item> {
        const itemResponse = await fetch(`${API_BASE_PATH}/items`, {
            method: "POST",
            body: JSON.stringify(item),
            headers: {
                "Content-Type": "application/json",
            },
        });
        return itemResponse.json();
    }

    async updateItem(item: Partial<Items.Item>): Promise<Items.Item> {
        const itemResponse = await fetch(`${API_BASE_PATH}/items/${item.id}`, {
            method: "PUT",
            body: JSON.stringify(item),
            headers: {
                "Content-Type": "application/json",
            },
        });
        return itemResponse.json();
    }

    async deleteItem(id: string): Promise<void> {
        await fetch(`${API_BASE_PATH}/items/${id}`, {
            method: "DELETE",
        });
    }

    async getCollections(): Promise<Collections.Collection[]> {
        const collectionResponse = await fetch(`${API_BASE_PATH}/collections`, {
            method: "GET",
        });
        return collectionResponse.json();
    }

    async getCollection(id: string): Promise<Collections.Collection> {
        const collectionResponse = await fetch(`${API_BASE_PATH}/collections/${id}`, {
            method: "GET",
        });
        return collectionResponse.json();
    }

    async createCollection(collection: Omit<Collections.Collection, 'id'>): Promise<Collections.Collection> {
        const collectionResponse = await fetch(`${API_BASE_PATH}/collections`, {
            method: "POST",
            body: JSON.stringify(collection),
            headers: {
                "Content-Type": "application/json",
            },
        });
        return collectionResponse.json();
    }

    async updateCollection(collection: Partial<Collections.Collection>): Promise<Collections.Collection> {
        const collectionResponse = await fetch(`${API_BASE_PATH}/collections/${collection.id}`, {
            method: "PUT",
            body: JSON.stringify(collection),
            headers: {
                "Content-Type": "application/json",
            },
        });
        return collectionResponse.json();
    }

    async deleteCollection(id: string): Promise<void> {
        await fetch(`${API_BASE_PATH}/collections/${id}`, {
            method: "DELETE",
        });
    }

    async addItemToCollection(itemId: string, collectionId: string): Promise<void> {
        await fetch(`${API_BASE_PATH}/collections/${collectionId}/items/${itemId}`, {
            method: "PUT",
        });
    }

    async removeItemFromCollection(itemId: string, collectionId: string): Promise<void> {
        await fetch(`${API_BASE_PATH}/collections/${collectionId}/items/${itemId}`, {
            method: "DELETE",
        });
    }
}

export const api = new Api();