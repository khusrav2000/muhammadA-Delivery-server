DELETE FROM role_permissions;
--Admin
INSERT INTO role_permissions(role,permission,has) VALUES ('admin', 'mobile_main', TRUE);
INSERT INTO role_permissions(role,permission,has) VALUES ('admin', 'catalog', TRUE);
INSERT INTO role_permissions(role,permission,has) VALUES ('admin', 'orders', TRUE);
INSERT INTO role_permissions(role,permission,has) VALUES ('admin', 'favorite', TRUE);
INSERT INTO role_permissions(role,permission,has) VALUES ('admin', 'profile', TRUE);
INSERT INTO role_permissions(role,permission,has) VALUES ('admin', 'create_users', TRUE);
INSERT INTO role_permissions(role,permission,has) VALUES ('admin', 'edit_users', TRUE);
INSERT INTO role_permissions(role,permission,has) VALUES ('admin', 'delete_users', TRUE);
INSERT INTO role_permissions(role,permission,has) VALUES ('admin', 'add_product', TRUE);
INSERT INTO role_permissions(role,permission,has) VALUES ('admin', 'edit_product', TRUE);
INSERT INTO role_permissions(role,permission,has) VALUES ('admin', 'delete_product', TRUE);
INSERT INTO role_permissions(role,permission,has) VALUES ('admin', 'add_pharmacy', TRUE);
INSERT INTO role_permissions(role,permission,has) VALUES ('admin', 'edit_pharmacy', TRUE);
INSERT INTO role_permissions(role,permission,has) VALUES ('admin', 'delete_pharmacy', TRUE);

--sales repse...
INSERT INTO role_permissions(role,permission,has) VALUES ('sales_rep', 'orders', TRUE);
INSERT INTO role_permissions(role,permission,has) VALUES ('sales_rep', 'profile', TRUE);
INSERT INTO role_permissions(role,permission,has) VALUES ('sales_rep', 'add_pharmacy', TRUE);
INSERT INTO role_permissions(role,permission,has) VALUES ('sales_rep', 'edit_pharmacy', TRUE);
INSERT INTO role_permissions(role,permission,has) VALUES ('sales_rep', 'delete_pharmacy', TRUE);
INSERT INTO role_permissions(role,permission,has) VALUES ('sales_rep', 'order_accept', TRUE);
INSERT INTO role_permissions(role,permission,has) VALUES ('sales_rep', 'make_order_delivered', TRUE);

--pharmacy_worker
INSERT INTO role_permissions(role,permission,has) VALUES ('pharmacy_worker', 'mobile_main', TRUE);
INSERT INTO role_permissions(role,permission,has) VALUES ('pharmacy_worker', 'catalog', TRUE);
INSERT INTO role_permissions(role,permission,has) VALUES ('pharmacy_worker', 'orders', TRUE);
INSERT INTO role_permissions(role,permission,has) VALUES ('pharmacy_worker', 'favorite', TRUE);
INSERT INTO role_permissions(role,permission,has) VALUES ('pharmacy_worker', 'profile', TRUE);
INSERT INTO role_permissions(role,permission,has) VALUES ('pharmacy_worker', 'order_create', TRUE);
INSERT INTO role_permissions(role,permission,has) VALUES ('pharmacy_worker', 'order_edit', TRUE);
INSERT INTO role_permissions(role,permission,has) VALUES ('pharmacy_worker', 'order_delete', TRUE);
INSERT INTO role_permissions(role,permission,has) VALUES ('pharmacy_worker', 'make_order_delivered', TRUE);
