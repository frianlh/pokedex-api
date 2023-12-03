CREATE TABLE IF NOT EXISTS public.monsters
(
    id                  uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    monster_code        integer          NOT NULL,
    name                varchar(255)     NOT NULL,
    monster_category_id uuid             NOT NULL,
    description         text             NOT NULL,
    length              float8           NOT NULL,
    weight              int              NOT NULL,
    hp                  int              NOT NULL,
    attack              int              NOT NULL,
    defends             int              NOT NULL,
    speed               int              NOT NULL,
    is_caught           bool             NOT NULL DEFAULT FALSE,
    image_name          text             NOT NULL,
    created_at          timestamp        NOT NULL DEFAULT now(),
    updated_at          timestamp        NOT NULL DEFAULT now(),
    deleted_at          timestamp
);