CREATE TABLE say_hello
(
    id         BIGINT       NOT NULL,
    name       VARCHAR(255) NOT NULL,
    created_at TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP    NOT NULL DEFAULT NOW() ON UPDATE NOW(),
    PRIMARY KEY (id)
) CHARACTER SET utf8mb4
  COLLATE utf8mb4_unicode_ci;
