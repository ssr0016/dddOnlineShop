
CREATE TABLE products (
    id serial PRIMARY KEY,
    sku varchar(100) NOT NULL,
    name varchar(100) NOT NULL,
    price int NOT NULL DEFAULT 0,
    stock int NOT NULL DEFAULT 0,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp NOT NULL DEFAULT NOW()
)