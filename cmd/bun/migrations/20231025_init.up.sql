SET statement_timeout = 0;

--bun:split

CREATE TABLE "public"."members"
(
    id         SERIAL PRIMARY KEY,
    user_id    int         not null,
    cafe_id    int         not null,
    nickname   VARCHAR(50) NOT NULL,
    is_banned  bool        not null default false,
    created_at timestamptz
);

create unique index user_cafe_id_unique on members (user_id, cafe_id);

create unique index cafe_id_nickname_unique on members (cafe_id, nickname);
