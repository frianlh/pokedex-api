CREATE TABLE IF NOT EXISTS public.users
(
    id                 uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    name               varchar(255)     NOT NULL,
    email              varchar(255)     NOT NULL,
    encrypted_password text             NOT NULL,
    role_id            uuid             NOT NULL,
    created_at         timestamp        NOT NULL DEFAULT now(),
    updated_at         timestamp        NOT NULL DEFAULT now(),
    deleted_at         timestamp
);