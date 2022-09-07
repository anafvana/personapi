DROP DATABASE personapi;
CREATE DATABASE personapi ENCODING 'SQL_ASCII' TEMPLATE template0 LC_COLLATE 'en_US.UTF-8' LC_CTYPE 'en_US.UTF-8';
\c personapi;

CREATE TABLE Person (
    bruker_id SERIAL PRIMARY KEY, 
    fornavn TEXT,
    etternavn TEXT
);