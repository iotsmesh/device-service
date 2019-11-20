CREATE TABLE IF NOT EXISTS devices (
    id varchar(31) NOT NULL CONSTRAINT device_pkey PRIMARY KEY,
    name varchar(255),
    comment varchar,
    user_id varchar(31),
    type varchar(255),
    tenant_id varchar(31)
);