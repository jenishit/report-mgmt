-- +goose Up
CREATE TABLE reports (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    report_no VARCHAR(100) UNIQUE,
    visit_id UUID REFERENCES visits(id) NOT NULL,
    generated_by UUID NOT NULL REFERENCES users(id),
    generated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    status VARCHAR(50) NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'final', 'amended')),
    pdf_path VARCHAR(255)
);

CREATE TABLE report_print (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    report_id UUID REFERENCES reports(id),
    printed_by UUID REFERENCES users(id),
    printed_at TIMESTAMP NOT NULL DEFAULT NOW(),
    copy_number INT
);

-- +goose Down
DROP TABLE report_print;
DROP TABLE reports;
