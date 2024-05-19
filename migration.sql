-- Миграция для создания таблицы Order
CREATE TABLE IF NOT EXISTS "order" (
    uid                 VARCHAR(255) PRIMARY KEY,
    track_number        VARCHAR(255),
    entry               VARCHAR(255),
    locale              VARCHAR(255),
    internal_signature  VARCHAR(255),
    customer_id         VARCHAR(255),
    delivery_service    VARCHAR(255),
    shardkey            VARCHAR(255),
    sm_id               INT,
    date_created        VARCHAR(255),
    oof_shard           VARCHAR(255)
);

-- Миграция для создания таблицы Delivery
CREATE TABLE IF NOT EXISTS delivery (
    id       serial PRIMARY KEY,
    order_uid VARCHAR(255) REFERENCES "order"(uid),
    name     VARCHAR(255),
    phone    VARCHAR(255),
    zip      VARCHAR(255),
    city     VARCHAR(255),
    address  VARCHAR(255),
    region   VARCHAR(255),
    email    VARCHAR(255)
);

-- Миграция для создания таблицы Payment
CREATE TABLE IF NOT EXISTS payment (
    id            serial PRIMARY KEY,
    order_uid     VARCHAR(255) REFERENCES "order"(uid),
    transaction   VARCHAR(255),
    currency      VARCHAR(255),
    provider      VARCHAR(255),
    request_id    VARCHAR(255),
    amount        INT,
    payment_dt    INT,
    bank          VARCHAR(255),
    delivery_cost INT,
    goods_total   INT,
    custom_fee    INT
);

-- Миграция для создания таблицы Item
CREATE TABLE IF NOT EXISTS item (
    id           serial PRIMARY KEY,
    order_uid    VARCHAR(255) REFERENCES "order"(uid),
    chrt_id      INT,
    track_number VARCHAR(255),
    price        INT,
    rid          VARCHAR(255),
    name         VARCHAR(255),
    sale         INT,
    size         VARCHAR(255),
    total_price  INT,
    nm_id        INT,
    brand        VARCHAR(255),
    status       INT
);