CREATE TABLE invoice (
    id serial,
    reference_year int4 NULL,
    reference_month int4 NULL,
    document varchar(14) NULL,
    "description" varchar(256) NULL,
    amount numeric(16,2) NULL,
    is_active bool NULL,
    created_at timestamp NULL,
    deactivated_at timestamp NULL,
    CONSTRAINT invoice_PK PRIMARY KEY (id)
);