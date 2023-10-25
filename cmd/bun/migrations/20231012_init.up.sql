SET statement_timeout = 0;

--bun:split

CREATE TABLE "public"."users"
(
    id              SERIAL PRIMARY KEY,
    email           VARCHAR(50) UNIQUE NOT NULL,
    password        VARCHAR(200)        not null,
    role            varchar(50),
    is_deleted      bool,
    created_at      timestamptz,
    last_updated_at timestamptz
);
