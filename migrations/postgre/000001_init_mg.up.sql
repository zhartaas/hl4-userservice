CREATE TABLE users (
                       id UUID PRIMARY KEY,
                       name VARCHAR(100) NOT NULL,
                       email VARCHAR(100) UNIQUE NOT NULL,
                       address VARCHAR(100) NOT NULL,
                       registration_date DATE NOT NULL DEFAULT CURRENT_DATE,
                       role VARCHAR(20) NOT NULL CHECK (role IN ('administrator', 'client'))
)

