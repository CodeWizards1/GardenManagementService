BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE garden_type AS ENUM ('balcony', 'rooftop', 'indoor', 'community', 'backyard');
CREATE TYPE plant_status AS ENUM ('planned', 'planted', 'growing', 'harvesting', 'dormant');

CREATE TABLE IF NOT EXISTS gardens (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    name VARCHAR(100) NOT NULL,
    type garden_type,
    area_sqm DECIMAL(6, 2),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS plants (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    garden_id UUID REFERENCES gardens(id),
    species VARCHAR(100) NOT NULL,
    quantity INTEGER,
    planting_date DATE,
    status plant_status DEFAULT 'planned',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS care_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    plant_id UUID REFERENCES plants(id),
    action VARCHAR(50),
    notes TEXT,
    logged_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

COMMIT;