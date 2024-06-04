CREATE TABLE IF NOT EXISTS transactions (
    transaction_id serial primary key,
    user_id int not null,
    product_id int not null,
    quantity int check(quantity>0) not null,
    total_amount decimal(10,2) not null,
    constraint fk_user_id foreign key (user_id) references users(user_id),
    constraint fk_product_id foreign key (product_id) references products(product_id)
);