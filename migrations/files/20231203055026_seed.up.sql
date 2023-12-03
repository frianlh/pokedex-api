-- roles table
INSERT INTO public.roles (id, name, created_at, updated_at, deleted_at)
VALUES ('6e3acdce-9b17-498e-aae4-ca8c92cd5b34', 'Admin', '2023-11-30 13:59:02.694320', '2023-11-30 13:59:02.694320',
        null),
       ('1247d7e3-50da-4924-9c22-960a073b5a73', 'User', '2023-11-30 13:59:02.694320', '2023-11-30 13:59:02.694320',
        null);

-- permissions table
INSERT INTO public.permissions (id, name, action, created_at, updated_at, deleted_at)
VALUES ('d2477e17-02db-44f3-bbb7-470706f94888', 'write_monster', 'CREATE', '2023-11-30 13:58:40.742799',
        '2023-11-30 13:58:40.742799', null),
       ('b41b3f95-3a96-4619-bfcc-72887f1fda6d', 'read_monster', 'READ', '2023-11-30 13:58:40.742799',
        '2023-11-30 13:58:40.742799', null),
       ('ccccebbd-d4eb-4229-a06c-632c561868c1', 'update_monster', 'UPDATE', '2023-11-30 13:58:40.742799',
        '2023-11-30 13:58:40.742799', null),
       ('527d1ddc-cdf7-4ff2-9706-8c300530f4ad', 'delete_monster', 'DELETE', '2023-11-30 13:58:40.742799',
        '2023-11-30 13:58:40.742799', null);

-- role permissions table
INSERT INTO public.role_permissions (role_id, permission_id)
VALUES ('6e3acdce-9b17-498e-aae4-ca8c92cd5b34', 'd2477e17-02db-44f3-bbb7-470706f94888'),
       ('6e3acdce-9b17-498e-aae4-ca8c92cd5b34', 'b41b3f95-3a96-4619-bfcc-72887f1fda6d'),
       ('6e3acdce-9b17-498e-aae4-ca8c92cd5b34', 'ccccebbd-d4eb-4229-a06c-632c561868c1'),
       ('6e3acdce-9b17-498e-aae4-ca8c92cd5b34', '527d1ddc-cdf7-4ff2-9706-8c300530f4ad'),
       ('1247d7e3-50da-4924-9c22-960a073b5a73', 'b41b3f95-3a96-4619-bfcc-72887f1fda6d');

-- users table
INSERT INTO public.users (id, name, email, encrypted_password, role_id, created_at, updated_at, deleted_at)
VALUES ('ee5d6b37-843f-43a0-812a-6a8f390104b8', 'Dev Admin', 'dev.admin@gmail.com',
        '$2a$04$uoa3ZYhS6cnTlmD/Uh/hTeRbSd.Jg9asNBKAluOw4hvdsQRYb6LSK', '6e3acdce-9b17-498e-aae4-ca8c92cd5b34',
        '2023-11-30 13:08:13.815321', '2023-11-30 13:08:13.815321', null),
       ('edfc5894-334a-401a-9568-a2f266d9a29c', 'Dev User', 'dev.user@gmail.com',
        '$2a$04$uoa3ZYhS6cnTlmD/Uh/hTeRbSd.Jg9asNBKAluOw4hvdsQRYb6LSK', '1247d7e3-50da-4924-9c22-960a073b5a73',
        '2023-11-30 13:08:13.815321', '2023-11-30 13:08:13.815321', null);

-- monster types table
INSERT INTO public.monster_types (id, name, created_at, updated_at, deleted_at)
VALUES ('3906338a-1393-4ccc-84ea-fa6a75034e09', 'GRASS', '2023-11-30 18:31:17.168980', '2023-11-30 18:31:17.168980',
        null),
       ('8653ceaf-36ea-4c30-b7af-8a5c6147f3a2', 'PSYCHIC', '2023-11-30 18:31:17.168980', '2023-11-30 18:31:17.168980',
        null),
       ('2447cb99-6cb7-42a2-8805-7e0dcd1fe3b6', 'FLYING', '2023-11-30 18:31:17.168980', '2023-11-30 18:31:17.168980',
        null),
       ('2e4081aa-08f8-4b68-a776-cc59d9cbbcbb', 'FIRE', '2023-11-30 18:31:17.168980', '2023-11-30 18:31:17.168980',
        null),
       ('ad80395c-9955-4723-aa55-379121a7167e', 'WATER', '2023-11-30 18:31:17.168980', '2023-11-30 18:31:17.168980',
        null),
       ('9c0ad005-1c06-41ed-ada2-9f155d21919e', 'ELECTRIC', '2023-11-30 18:31:17.168980', '2023-11-30 18:31:17.168980',
        null),
       ('4e2d0029-f754-4dba-8f03-b923a2ce03d3', 'BUG', '2023-11-30 18:31:17.168980', '2023-11-30 18:31:17.168980',
        null);

-- monster categories table
INSERT INTO public.monster_categories (id, name, created_at, updated_at, deleted_at)
VALUES ('6eaac59e-33ef-4f99-8244-512deb0dc151', 'Leaf Monster', '2023-11-30 17:49:08.138266',
        '2023-11-30 17:49:08.138266', null),
       ('9dfef172-34d0-454a-b92e-c8ceba59f487', 'Diving Monster', '2023-11-30 17:49:08.138266',
        '2023-11-30 17:49:08.138266', null),
       ('de75244f-139c-47ac-ba21-3cf42612a0f4', 'Lizard Monster', '2023-11-30 17:49:08.138266',
        '2023-11-30 17:49:08.138266', null);
