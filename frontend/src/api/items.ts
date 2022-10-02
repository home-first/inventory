


export interface Item {
    id: string;
    name: string;
    collections: Array<string>;
}

export const EmptyItem: Item = {
    id: "",
    name: "",
    collections: [],
};