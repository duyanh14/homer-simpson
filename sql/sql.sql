ALTER TABLE user_roles ADD CONSTRAINT user_roles_user_id_role_id UNIQUE (user_id, role_id);
ALTER TABLE role_permissions ADD CONSTRAINT user_roles_user_id_permission_id UNIQUE (permission_id, role_id);

ALTER TABLE access ADD CONSTRAINT accepts_type UNIQUE (type);