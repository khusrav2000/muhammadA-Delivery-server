CREATE TYPE user_role as ENUM (
    'admin',
    'pharmacy_worker',
    'sales_rep'
);

CREATE TYPE permission as ENUM(
    'mobile_main',
    'catalog',
    'orders',
    'favorite',
    'profile',
    'pharmacies',
    'create_users',
    'edit_users',
    'delete_users',
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
    'make_order_delivered'
);

CREATE TABLE role_permissions (
    id bigserial not null primary key,
    role user_role not null,
    permission permission not null,
    has boolean default FALSE
);





CREATE TABLE user_roles (
    user_id bigserial not null,
    role user_role not null
);