-- +goose Up

-- Seed the two privileged roles. ROLE_ADMIN must exist up front because role
-- creation is now admin-only, and ROLE_SUPER_ADMIN is the one account
-- allowed to create other admins (see auth middleware / handler checks).
INSERT INTO role (id, role_name)
VALUES (gen_random_uuid(), 'ROLE_SUPER_ADMIN')
ON CONFLICT (role_name) DO NOTHING;

INSERT INTO role (id, role_name)
VALUES (gen_random_uuid(), 'ROLE_ADMIN')
ON CONFLICT (role_name) DO NOTHING;

-- Bootstrap super admin account: superadmin@gmail.com / Jenish@123
-- (bcrypt hash below, cost 14, matches valueobjects.PasswordHashCost)
INSERT INTO users (id, role_id, email, password)
SELECT gen_random_uuid(), r.id, 'superadmin@gmail.com',
       '$2a$14$9dmTS3BE2vIpWI13U/awie5f32yz0DqCTWFJtwsM7c1uFBDZm76.S'
FROM role r
WHERE r.role_name = 'ROLE_SUPER_ADMIN'
ON CONFLICT (email) DO NOTHING;

INSERT INTO profile (id, user_id, email, first_name, last_name, is_active)
SELECT gen_random_uuid(), u.id, u.email, 'Super', 'Admin', TRUE
FROM users u
WHERE u.email = 'superadmin@gmail.com'
ON CONFLICT (email) DO NOTHING;

-- +goose Down
DELETE FROM profile WHERE email = 'superadmin@gmail.com';
DELETE FROM users WHERE email = 'superadmin@gmail.com';
DELETE FROM role WHERE role_name = 'ROLE_SUPER_ADMIN';
