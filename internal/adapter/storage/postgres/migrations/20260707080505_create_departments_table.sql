-- +goose Up
CREATE TABLE departments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL UNIQUE,
    description VARCHAR(255),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE panels (
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    department_id UUID NOT NULL REFERENCES departments(id),
    name          VARCHAR(100) NOT NULL,
    code          VARCHAR(30) UNIQUE,
    panel_price   NUMERIC(10,2),
    is_active     BOOLEAN NOT NULL DEFAULT TRUE,
    updated_at     TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE test_catalog (
    id               UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    department_id    UUID NOT NULL REFERENCES departments(id),
    name             VARCHAR(100) NOT NULL,
    code             VARCHAR(30) UNIQUE,
    sample_type      VARCHAR(50) CHECK (sample_type IN ('WHOLE_BLOOD','SERUM','URINE','PLASMA','SWAB','STOOL','OTHER')),
    price            NUMERIC(10,2) NOT NULL DEFAULT 0,
    turnaround_hours SMALLINT,
    is_active        BOOLEAN NOT NULL DEFAULT TRUE,
    updated_at        TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE panel_components (
    panel_id    UUID NOT NULL REFERENCES panels(id) ON DELETE CASCADE,
    test_id     UUID NOT NULL REFERENCES test_catalog(id),
    sequence_no SMALLINT NOT NULL DEFAULT 1,
    PRIMARY KEY (panel_id, test_id)
);

CREATE TABLE test_parameters (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    test_id     UUID NOT NULL REFERENCES test_catalog(id) ON DELETE CASCADE,
    name        VARCHAR(100) NOT NULL,
    unit        VARCHAR(30),
    result_type VARCHAR(20) NOT NULL DEFAULT 'numeric'
                    CHECK (result_type IN ('numeric','text','option')),
    sequence_no SMALLINT NOT NULL DEFAULT 1,
    is_active   BOOLEAN NOT NULL DEFAULT TRUE,
    updated_at   TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE reference_ranges (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    parameter_id UUID NOT NULL REFERENCES test_parameters(id) ON DELETE CASCADE,
    gender       VARCHAR(10) NOT NULL DEFAULT 'ALL'
                     CHECK (gender IN ('M','F','ALL')),
    age_min_years NUMERIC(5,2) NOT NULL DEFAULT 0,
    age_max_years NUMERIC(5,2) NOT NULL DEFAULT 150,
    low_value    NUMERIC(12,4),
    high_value   NUMERIC(12,4),
    text_range   VARCHAR(255),
    notes        VARCHAR(255),
    updated_at    TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE departments;
DROP TABLE panels;
DROP TABLE test_catalog;
DROP TABLE panel_components;
DROP TABLE test_parameters;
DROP TABLE reference_ranges;