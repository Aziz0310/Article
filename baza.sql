CREATE TABLE author (
    id VARCHAR(40) PRIMARY KEY NOT NULL,
    first_name VARCHAR(45) NOT NULL,
    last_name VARCHAR(45) NOT NULL,
    created_at TIMESTAMP DEFAULT now(), 
    updated_at TIMESTAMP, 
    deleted_at TIMESTAMP
);

create table article(
    id VARCHAR(40) PRIMARY KEY NOT NULL, 
    title VARCHAR(55) NOT NULL, 
    body text NOT NULL, 
    author_id VARCHAR(40) NOT NULL REFERENCES author(id), 
    created_at TIMESTAMP DEFAULT now(), 
    updated_at TIMESTAMP, 
    deleted_at TIMESTAMP
);