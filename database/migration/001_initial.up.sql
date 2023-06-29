CREATE TABLE IF NOT EXISTS users (
  id BIGSERIAL PRIMARY KEY,
  email TEXT UNIQUE NOT NULL,
  first_name TEXT NOT NULL,
  last_name TEXT NOT NULL,
  avatar TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS properties (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  town VARCHAR(50) NOT NULL,
  postal_code VARCHAR(20) NOT NULL,
  type VARCHAR(20) NOT NULL,
  min_price INTEGER NOT NULL DEFAULT 0,
  max_price INTEGER NOT NULL DEFAULT 0,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  created_by BIGINT NOT NULL REFERENCES users ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS amenities (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  provider VARCHAR(100) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  property_id BIGINT NOT NULL REFERENCES properties ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS property_units (
  id BIGSERIAL PRIMARY KEY,
  bathrooms INTEGER NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  property_id BIGINT NOT NULL REFERENCES properties ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS tenants (
  id BIGSERIAL PRIMARY KEY,
  start_date TIMESTAMP NOT NULL,
  end_date TIMESTAMP,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  property_unit_id BIGINT NOT NULL REFERENCES property_units ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS bedrooms (
  id BIGSERIAL PRIMARY KEY,
  bedroom_number INTEGER NOT NULL,
  en_suite BOOLEAN NOT NULL DEFAULT false,
  master BOOLEAN NOT NULL DEFAULT false,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  property_unit_id BIGINT NOT NULL REFERENCES property_units ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS caretakers (
  id BIGSERIAL PRIMARY KEY,
  first_name VARCHAR(100) NOT NULL,
  last_name VARCHAR(100) NOT NULL,
  idVerification TEXT NOT NULL,
  country_code VARCHAR(5) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  phone VARCHAR(12) UNIQUE,
  verified BOOLEAN NOT NULL DEFAULT false
);

CREATE TABLE IF NOT EXISTS shoots (
  id BIGSERIAL PRIMARY KEY,
  shoot_date TIMESTAMP NOT NULL,
  property_id BIGINT NOT NULL REFERENCES properties ON DELETE CASCADE,
  property_unit_id BIGINT NOT NULL REFERENCES property_units ON DELETE CASCADE,
  status VARCHAR(20) NOT NULL DEFAULT 'pending',
  caretaker_id BIGINT NOT NULL REFERENCES caretakers ON DELETE CASCADE
);
