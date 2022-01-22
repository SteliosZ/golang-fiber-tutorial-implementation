-- ./platform/migrations/000001_create_init_tables.up.sql

-- Add UUID extension
CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
-- For more information, please visit:
-- https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
SET
TIMEZONE="Europe/Athens";

-- Create books table
CREATE TABLE vinyls
(
    id           UUID                     DEFAULT uuid_generate_v4() PRIMARY KEY,
    created_at   TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at   TIMESTAMP NULL,
    title        VARCHAR(255) NOT NULL,
    artist       VARCHAR(255) NOT NULL,
    vinyl_status INT          NOT NULL,
    vinyl_attrs  JSONB        NOT NULL
);

-- Add indexes
CREATE INDEX active_vinyls ON vinyls (title) WHERE vinyl_status = 1;