export enum UsersRoleOptions {
    Admin = "admin",
    Editor = "editor"
}

export interface Users {
    expand?: { [key: string]: unknown; };
    id: string;
    email: string;
    emailVisibility: boolean;
    verified: boolean;
    name: string;
    avatar: string;
    created: string;
    updated: string;
    display_name: string;
    role: UsersRoleOptions;
}

export interface Categories {
    expand?: { [key: string]: unknown; };
    id: string;
    name: string;
    slug: string;
    description: string;
    image_url: string;
    is_active: boolean;
    created: string;
    updated: string;
}

export interface ProductsExpanded {
    category_id?: Categories;
    [key: string]: unknown;
}

export interface Products {
    expand?: ProductsExpanded;
    id: string;
    category_id: string;
    name: string;
    slug: string;
    description: string;
    price: number;
    images: object | null | "";
    is_active: boolean;
    created: string;
    updated: string;
}

export enum InventoryStatusOptions {
    InStock = "in_stock",
    Low = "low",
    Out = "out"
}

export interface InventoryExpanded {
    product_id?: Products;
    [key: string]: unknown;
}

export interface Inventory {
    expand?: InventoryExpanded;
    id: string;
    product_id: string;
    quantity: number;
    reserved: number;
    status: InventoryStatusOptions;
    created: string;
    updated: string;
}

export enum PurchaseIntentsStatusOptions {
    Pending = "pending",
    Contacted = "contacted",
    Closed = "closed"
}

export interface PurchaseIntentsExpanded {
    product_id?: Products;
    [key: string]: unknown;
}

export interface PurchaseIntents {
    expand?: PurchaseIntentsExpanded;
    id: string;
    product_id: string;
    customer_name: string;
    customer_email: string;
    customer_phone: string;
    quantity: number;
    message: string;
    status: PurchaseIntentsStatusOptions;
    price_at_intent: number;
    created: string;
    updated: string;
}