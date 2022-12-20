

export interface Collection {
    id: string;
    name: string;
    items: string[];
}

export const EmptyCollection: Collection = {
    id: "",
    name: "",
    items: [],
};