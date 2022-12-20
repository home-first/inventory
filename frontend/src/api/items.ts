


export interface Item {
    id: string;
    name: string;
    description: string;
    collections: Array<string>;
}

export const EmptyItem: Item = {
    id: "",
    name: "",
    description: "",
    collections: [],
};