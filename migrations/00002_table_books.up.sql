BEGIN;

CREATE TABLE books
(
    id           VARCHAR(36) PRIMARY KEY             NOT NULL,
    author_id    VARCHAR(36) REFERENCES authors (id),
    title        VARCHAR(36)                         NOT NULL,
    description  TEXT                                NOT NULL,
    is_published TINYINT(1) DEFAULT 0 NOT NULL,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL
);

COMMIT;
