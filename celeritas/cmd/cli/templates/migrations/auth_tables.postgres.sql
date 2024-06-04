DROP TABLE IF EXISTS users CASCADE;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name varchar(255) NOT NULL,
    last_name varchar(255) NOT NULL,
    user_active integer NOT NULL,
    email varchar(255) NOT NULL,
    password char(60) NOT NULL,
    created_at timestamp without time zone DEFAULT NULL,
    updated_at timestamp without time zone DEFAULT NULL,
    UNIQUE (email)
);

-- DROP TABLE IF EXISTS remember_tokens CASCADE;

-- CREATE TABLE remember_tokens (
--     id SERIAL PRIMARY KEY,
--     user_id integer NOT NULL,
--     remember_token varchar(100) NOT NULL DEFAULT '',
--     created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
--     updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
--     FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
-- );

DROP TABLE IF EXISTS remember_tokens CASCADE;

CREATE TABLE remember_tokens (
    id SERIAL PRIMARY KEY,
    user_id integer NOT NULL,
    remember_token varchar(100) NOT NULL DEFAULT '',
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- Создание функции триггера для обновления временной метки
CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Создание триггера, который вызывает функцию при обновлении записей
CREATE TRIGGER update_remember_tokens_modtime
BEFORE UPDATE ON remember_tokens
FOR EACH ROW
EXECUTE FUNCTION update_modified_column();

DROP TABLE IF EXISTS tokens CASCADE;

CREATE TABLE tokens (
    id SERIAL PRIMARY KEY,
    user_id integer NOT NULL,
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    token varchar(255) NOT NULL,
    token_hash bytea DEFAULT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expiry timestamp without time zone NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
);