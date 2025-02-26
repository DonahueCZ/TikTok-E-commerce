CREATE TABLE orders (
    `id` BIGINT PRIMARY KEY,
    `user_id` BIGINT NOT NULL,
    `goods_id` BIGINT NOT NULL,
    `status` INT NOT NULL,
    `create_time` BIGINT NOT NULL,
    `goods_count` INT NOT NULL,
    `cost` INT NOT NULL,
    `addressee_info` VARCHAR NOT NULL
);
