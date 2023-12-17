CREATE TABLE IF NOT EXISTS users (
    id INT auto_increment primary key,
    name VARCHAR(100) not null,
    email VARCHAR(70) unique not null,
    password VARCHAR(50) not null,
    addr_cep VARCHAR(9) not null,
    addr_street VARCHAR(50) not null,
    addr_number INT not null,
    addr_district VARCHAR(30) not null,
    addr_city VARCHAR(50) not null,
    addr_state VARCHAR(4) not null,
    confirmation_token VARCHAR(500),
    expiration_confirmation_token TIMESTAMP,
    created_at TIMESTAMP default now(),
    updated_at TIMESTAMP default now() on update now()
);