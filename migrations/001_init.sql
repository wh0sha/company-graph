CREATE TABLE IF NOT EXISTS companies (
    id SERIAL PRIMARY KEY, 
    name VARCHAR(255) NOT NULL,
    inn VARCHAR(12),
    ogrn VARCHAR(13),
    address TEXT,
    source VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW()
)