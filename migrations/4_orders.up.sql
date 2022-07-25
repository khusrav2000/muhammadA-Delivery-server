CREATE TABLE IF NOT EXISTS orders (
    id bigserial not null primary key,
    pharmacy_id bigserial not null,
    deliveryman_id bigserial not null,
    pharmacist_id bigserial not null,
    order_date timestamp without time zone,
    delivery_date timestamp without time zone
);