-- +goose Up
CREATE TABLE visits (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    visit_no VARCHAR(100) UNIQUE,
    patient_id UUID NOT NULL REFERENCES patients(id),
    doctor_id UUID REFERENCES doctors(id),
    registered_by UUID NOT NULL REFERENCES users(id),
    visit_date TIMESTAMP NOT NULL DEFAULT NOW(),
    v_status VARCHAR(100) NOT NULL DEFAULT 'in_progress'
                     CHECK (v_status IN ('registered','in_progress','completed', 'cancelled')),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE order_item (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    visit_id UUID NOT NULL REFERENCES visits(id),
    test_id UUID NOT NULL REFERENCES test_catalog(id),
    panel_id UUID REFERENCES panels(id),
    status VARCHAR(100) NOT NULL DEFAULT 'collected' CHECK (status IN ('completed','collected','result_entered')),
    price NUMERIC(10,2),
    collected_by UUID NOT NULL REFERENCES users(id),
    collected_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE result (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id UUID NOT NULL REFERENCES order_item(id),
    parameter_id UUID NOT NULL REFERENCES test_parameters(id),
    result_value VARCHAR(50) NOT NULL,
    flag VARCHAR(50) NOT NULL DEFAULT 'normal' CHECK (flag IN ('low','normal','high', 'critical', 'na')),
    performed_by UUID NOT NULL REFERENCES users(id),
    performed_at TIMESTAMP NOT NULL DEFAULT NOW(),
    verified_by UUID NOT NULL REFERENCES users(id),
    remarks VARCHAR(100),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE result;
DROP TABLE order_item;
DROP TABLE visits;