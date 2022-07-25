CREATE TABLE IF NOT EXISTS order_products (
    id bigserial not null primary key,
    order_id bigserial not null,
    product_name varchar,
    price float,
    count int
);