-- +goose Up
CREATE TABLE doctors (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(100) UNIQUE,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    phone VARCHAR(50),
    qualification VARCHAR(100),
    registration_no VARCHAR(100),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_by UUID DEFAULT NULL
);

CREATE TABLE patients (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(100) UNIQUE,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    phone VARCHAR(50),
    mrn VARCHAR(50),
    dob DATE NOT NULL,
    gender VARCHAR(10) NOT NULL DEFAULT 'O'
                     CHECK (gender IN ('M','F','O')),
    p_address VARCHAR(50),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_by UUID DEFAULT NULL
);

-- +goose Down
DROP TABLE patients;
DROP TABLE doctors;