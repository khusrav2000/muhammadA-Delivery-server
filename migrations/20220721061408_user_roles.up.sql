CREATE TYPE user_role as ENUM (
    'admin',
    'pharmacy_worker',
    'sales_rep'
);

CREATE TYPE role_permission as ENUM(
    'mobile_main',
    'catalog',
    'orders',
    'favorite',
    'profile',
    'add_product',
    'edit_product',
    'delete_product',
    'add_pharmacy',
    'edit_pharmacy',
    'delete_pharmacy',
    'order_create',
    'order_edit',
    'order_delete',
    'order_accept',
    'make_order_delivered',

);

CREATE TABLE user_roles (
    user_id not null,
    role_id bigserial not null,
);

CREATE TABLE role (
    id bigserial not null primary key,

);