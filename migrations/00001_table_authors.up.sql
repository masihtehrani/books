BEGIN;

CREATE TABLE authors
(
    id         VARCHAR(36) PRIMARY KEY             NOT NULL,
    full_name   VARCHAR(255)                        NOT NULL,
    pseudonym  VARCHAR(25)                         NOT NULL,
    username   VARCHAR(25)                         NOT NULL,
    password   VARCHAR(36)                         NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    constraint authors_username_uindex
        unique (username)
);

COMMIT;


