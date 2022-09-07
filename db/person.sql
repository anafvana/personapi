SELECT 'CREATE DATABASE personapi'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'personapi')\gexec

GRANT ALL PRIVILEGES ON DATABASE personapi TO user123;
\c personapi;

CREATE TABLE IF NOT EXISTS Person (
    bruker_id SERIAL PRIMARY KEY, 
    fornavn TEXT,
    etternavn TEXT
);