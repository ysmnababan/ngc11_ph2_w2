CREATE TABLE if NOT EXISTS products (
    product_id serial primary key,
    name varchar(255) UNIQUE,
    stock int check(stock>0),
    price decimal(10,2) check (price>0)
);